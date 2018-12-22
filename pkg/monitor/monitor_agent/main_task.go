package monitor_agent

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *MonitorAgentServer) MainTask(traceId string) error {
	if err := srv.Report(traceId); err != nil {
		return err
	}

	return nil
}

func (srv *MonitorAgentServer) Report(traceId string) error {
	var err error
	startTime := logger.StartTaskTrace(traceId, srv.Host, srv.Name)
	defer func() { logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, err) }()

	return nil
}
