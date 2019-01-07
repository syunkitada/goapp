package monitor_agent

import (
	// "fmt"

	// "github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

func (srv *MonitorAgentServer) MainTask(tctx *logger.TraceContext) error {
	var err error

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
	pbLogs := []*monitor_api_grpc_pb.Log{}
	for _, logReader := range srv.logReaderMap {
		logs := logReader.GetLogs()
		pbLogs = append(pbLogs, logs...)
	}

	req := &monitor_api_grpc_pb.ReportRequest{
		Index:   srv.reportIndex,
		Project: srv.reportProject,
		Logs:    pbLogs,
	}

	_, err := srv.monitorApiClient.Report(req)
	return err
}
