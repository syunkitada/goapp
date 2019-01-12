package monitor_agent

import (
	"fmt"

	// "github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

var _ = fmt.Printf // For debugging: TODO Remove

func (srv *MonitorAgentServer) MainTask(tctx *logger.TraceContext) error {
	var err error

	for _, metricReader := range srv.metricReaders {
		err = metricReader.Read(tctx)
		if err != nil {
			logger.Warningf(tctx, err, "Failed metricReader.Read(): %v", metricReader.GetName())
		}
	}

	if srv.logReaderRefreshCount == 0 {
		for logName, logConf := range srv.conf.Monitor.AgentApp.LogMap {
			if _, ok := srv.logReaderMap[logName]; ok {
				continue
			}

			reader, err := NewLogReader(srv.conf, logName, &logConf)
			if err != nil {
				continue
			}
			srv.logReaderMap[logName] = reader
		}
	}

	if srv.logReaderRefreshCount >= srv.logReaderRefreshSpan {
		srv.logReaderRefreshCount = 0
	} else {
		srv.logReaderRefreshCount += 1
	}

	for logName, logReader := range srv.logReaderMap {
		err = logReader.ReadUntilEOF(tctx)
		if err != nil {
			logger.Warningf(tctx, err, "Failed logReader.ReadUntilEOF(): %v", logName)
		}
	}

	if srv.reportCount == 0 {
		srv.Report(tctx)
	}

	if srv.reportCount >= srv.reportSpan {
		srv.reportCount = 0
	} else {
		srv.reportCount += 1
	}

	return nil
}

func (srv *MonitorAgentServer) Report(tctx *logger.TraceContext) error {
	var err error

	pbMetrics := make([]*monitor_api_grpc_pb.Metric, 0, 100)
	for _, metricReader := range srv.metricReaders {
		metrics := metricReader.Report()
		pbMetrics = append(pbMetrics, metrics...)
	}

	pbLogs := make([]*monitor_api_grpc_pb.Log, 0, 100)
	for _, logReader := range srv.logReaderMap {
		logs := logReader.Report()
		pbLogs = append(pbLogs, logs...)
	}

	req := &monitor_api_grpc_pb.ReportRequest{
		Index:   srv.reportIndex,
		Project: srv.reportProject,
		Host:    srv.Host,
		Metrics: pbMetrics,
		Logs:    pbLogs,
	}

	_, err = srv.monitorApiClient.Report(req)
	if err != nil {
		return err
	}

	for _, metricReader := range srv.metricReaders {
		metricReader.Reported()
	}

	for _, logReader := range srv.logReaderMap {
		logReader.Reported()
	}

	return err
}
