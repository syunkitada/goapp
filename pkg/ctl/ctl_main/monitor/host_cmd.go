package monitor

import (
	"fmt"
	"os"
	// "time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var getHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Show hosts",
	Long: `Show hosts
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetHost(); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetHost() error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Monitor.CtlGetHost(token.Token, getCmdIndexFlag)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetHost.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})
	for _, host := range resp.HostMap {
		table.Append([]string{
			host.Name,
		})
	}
	table.Render()

	return nil
}
