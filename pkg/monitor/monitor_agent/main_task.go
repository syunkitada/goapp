package monitor_agent

import (
	"fmt"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	"github.com/hpcloud/tail"
)

func (srv *MonitorAgentServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.Report(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *MonitorAgentServer) Report(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err) }()

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
