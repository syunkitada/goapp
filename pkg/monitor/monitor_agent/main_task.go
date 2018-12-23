package monitor_agent

import (
	"fmt"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	"github.com/hpcloud/tail"
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

	t, err := tail.TailFile("/home/owner/.goapp/logs/goapp-resource-api.log", tail.Config{Follow: true})
	if err != nil {
		fmt.Print(err)
		return nil
	}
	for line := range t.Lines {
		fmt.Println(line.Text)
		break
	}
	t.Cleanup()

	return nil
}
