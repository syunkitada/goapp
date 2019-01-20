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

var getIndexCmd = &cobra.Command{
	Use:   "index",
	Short: "Show indexes",
	Long: `Show indexes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetIndex(); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetIndex() error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Monitor.CtlGetIndex(token.Token)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetIndex.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Count", "States", "Warnings", "Errors"})
	for _, index := range resp.IndexMap {
		table.Append([]string{
			index.Name,
			fmt.Sprint(index.Count),
			fmt.Sprint(index.States),
			fmt.Sprint(index.Warnings),
			fmt.Sprint(index.Errors),
		})
	}
	table.Render()

	return nil
}
