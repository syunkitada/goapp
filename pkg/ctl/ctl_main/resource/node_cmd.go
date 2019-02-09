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
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken(tctx)
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
	table.SetHeader([]string{"Cluster", "Kind", "Name", "Status", "Status Reason", "Updated At"})
	for clusterName, nodes := range resp.ClusterNodesMap {
		for _, node := range nodes.Nodes {
			table.Append([]string{
				clusterName,
				node.Kind,
				node.Name,
				node.Status,
				node.StatusReason,
				fmt.Sprint(time.Unix(node.UpdatedAt.Seconds, 0)),
			})
		}
	}
	table.Render()

	return nil
}
