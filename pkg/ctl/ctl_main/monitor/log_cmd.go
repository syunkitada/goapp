package monitor

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var getLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Show logs",
	Long: `Show logs
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetLog(); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetLog() error {
	var err error
	traceId := logger.NewTraceId()
	startTime := logger.StartCtlTrace(traceId, appName)
	defer func() {
		logger.EndCtlTrace(traceId, appName, startTime, err)
	}()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	authproxy.Monitor.CtlGetLog(token.Token)

	return nil
}
