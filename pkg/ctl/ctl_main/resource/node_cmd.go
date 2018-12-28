package resource

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
	traceId := logger.NewTraceContext()
	startTime := logger.StartCtlTrace(traceId, appName)
	defer func() {
		logger.EndCtlTrace(traceId, appName, startTime, err)
	}()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Resource.CtlGetNode(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetNode.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, node := range resp.Nodes {
		table.Append([]string{
			node.Cluster,
			node.Name,
			node.Status,
			node.StatusReason,
			fmt.Sprint(time.Unix(node.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(node.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}
