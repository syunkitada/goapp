package server

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/readers/log_reader"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) ReaderLoop() {
	var tctx *logger.TraceContext
	var startTime time.Time
	var err error
	loopInterval := 10 * time.Second
	logger.StdoutInfof("Start mainLoop")
	for {
		tctx = srv.NewTraceContext()
		startTime = logger.StartTrace(tctx)
		err = srv.ReaderTask(tctx)
		logger.EndTrace(tctx, startTime, err, 0)
		time.Sleep(loopInterval)
	}
}

func (srv *Server) ReaderTask(tctx *logger.TraceContext) (err error) {
	for metricsName, metricsReader := range srv.metricsReaderMap {
		err = metricsReader.Read(tctx)
		if err != nil {
			logger.Warningf(tctx, "Failed metricsReader.Read(): %s, err=%v", metricsName, err)
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
	events := make([]resource_api_spec.ResourceEvent, 0, 10)
	metrics := make([]resource_api_spec.ResourceMetric, 0, 100)
	logs := make([]resource_api_spec.ResourceLog, 0, 100)

	for _, metricsReader := range srv.metricsReaderMap {
		tmpMetrics, tmpEvents := metricsReader.Report()
		metrics = append(metrics, tmpMetrics...)
		events = append(events, tmpEvents...)
	}

	for _, logReader := range srv.logReaderMap {
		tmpLogs, tmpEvents := logReader.Report()
		logs = append(logs, tmpLogs...)
		events = append(events, tmpEvents...)
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "ReportNode",
			Data: resource_api_spec.ReportNode{
				Name:     srv.baseConf.Host,
				Project:  srv.clusterConf.Agent.ReportProject,
				Warning:  "",
				Warnings: 0,
				Error:    "",
				Errors:   0,
				Logs:     logs,
				Metrics:  metrics,
				Events:   events,
			},
		},
	}

	if _, err = srv.apiClient.ResourceVirtualAdminReportNode(tctx, queries); err != nil {
		return
	}

	for _, metricsReader := range srv.metricsReaderMap {
		metricsReader.Reported()
	}
	for _, logReader := range srv.logReaderMap {
		logReader.Reported()
	}

	return
}
