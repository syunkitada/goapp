package server

import (
	"fmt"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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
	// TODO Report metrics, logs
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

	fmt.Println("DEBUG queries", queries)
	return
}
