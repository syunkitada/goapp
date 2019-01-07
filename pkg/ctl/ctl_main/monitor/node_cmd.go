package monitor

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var getNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Show nodes",
	Long: `Show nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetNode(); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetNode() error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Monitor.CtlGetNode(token.Token, "%")
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetNode.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Kind", "Status", "Status Reason", "State", "State Reason", "Updated At", "Created At"})
	for _, node := range resp.Nodes {
		table.Append([]string{
			node.Cluster,
			node.Name,
			node.Kind,
			node.Status,
			node.StatusReason,
			node.State,
			node.StateReason,
			fmt.Sprint(time.Unix(node.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(node.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}
