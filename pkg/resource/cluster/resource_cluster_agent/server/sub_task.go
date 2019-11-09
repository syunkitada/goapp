package server

import (
	"fmt"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/readers/log_reader"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) StartSubLoop() {
	go srv.SubLoop()
}

func (srv *Server) SubLoop() {
	var tctx *logger.TraceContext
	var startTime time.Time
	var err error
	loopInterval := 10 * time.Second
	logger.StdoutInfof("Start mainLoop")
	for {
		tctx = srv.NewTraceContext()
		startTime = logger.StartTrace(tctx)
		err = srv.SubTask(tctx)
		logger.EndTrace(tctx, startTime, err, 0)
		time.Sleep(loopInterval)
	}
}

func (srv *Server) SubTask(tctx *logger.TraceContext) (err error) {
	fmt.Println("subTask")

	for metricName, metricReader := range srv.metricReaderMap {
		err = metricReader.Read(tctx)
		if err != nil {
			logger.Warningf(tctx, "Failed metricReader.Read(): %s, err=%v", metricName, err)
		}
	}

	// reinitialize logMap
	if srv.logReaderRefreshCount == 0 {
		for logName, logConf := range srv.logMap {
			if _, ok := srv.logReaderMap[logName]; ok {
				continue
			}

			reader, err := log_reader.New(srv.baseConf, logName, &logConf)
			if err != nil {
				continue
			}
			srv.logReaderMap[logName] = reader
		}
	}

	if srv.logReaderRefreshCount >= srv.logReaderRefreshSpan {
		srv.logReaderRefreshCount = 0
	} else {
		srv.logReaderRefreshCount++
	}

	for logName, logReader := range srv.logReaderMap {
		err = logReader.ReadUntilEOF(tctx)
		if err != nil {
			logger.Warningf(tctx, "Failed logReader.ReadUntilEOF(): %s, err=%v", logName, err)
		}
	}

	if srv.reportCount == 0 {
		if err = srv.Report(tctx); err != nil {
			return err
		}
	}

	if srv.reportCount >= srv.reportSpan {
		srv.reportCount = 0
	} else {
		srv.reportCount++
	}

	return
}

func (srv *Server) Report(tctx *logger.TraceContext) (err error) {
	alerts := make([]resource_api_spec.ResourceAlert, 0, 10)
	metrics := make([]resource_api_spec.ResourceMetric, 0, 100)
	logs := make([]resource_api_spec.ResourceLog, 0, 100)

	for _, metricReader := range srv.metricReaderMap {
		tmpMetrics, tmpAlerts := metricReader.Report()
		metrics = append(metrics, tmpMetrics...)
		alerts = append(alerts, tmpAlerts...)
	}

	for _, logReader := range srv.logReaderMap {
		tmpLogs, tmpAlerts := logReader.Report()
		logs = append(logs, tmpLogs...)
		alerts = append(alerts, tmpAlerts...)
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "ReportResource",
			Data: resource_api_spec.ReportResource{
				Host:     srv.baseConf.Host,
				Project:  srv.clusterConf.Agent.ReportProject,
				Warning:  "",
				Warnings: 0,
				Error:    "",
				Errors:   0,
				Logs:     logs,
				Metrics:  metrics,
				Alerts:   alerts,
			},
		},
	}

	if _, err = srv.apiClient.ResourceVirtualAdminReportResource(tctx, queries); err != nil {
		return
	}

	for _, metricReader := range srv.metricReaderMap {
		metricReader.Reported()
	}
	for _, logReader := range srv.logReaderMap {
		logReader.Reported()
	}
	return
}
