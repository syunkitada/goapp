package resource_ctl

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseGetNode struct {
	Nodes   []resource_api_grpc_pb.Node
	TraceId string
	Err     string
}

var getNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Show nodes",
	Long: `Show nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var target string
		if len(args) > 1 {
			target = args[0]
		} else {
			target = ""
		}

		ctl := New(&config.Conf, nil)
		if _, err := ctl.GetNode(getCmdClusterFlag, target); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *ResourceCtl) GetNode(cluster string, target string) (*ResponseGetNode, error) {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var token *authproxy_client.ResponseIssueToken
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return nil, err
	}

	req := resource_api_grpc_pb.ActionRequest{
		Cluster: cluster,
		Target:  target,
	}
	var resp ResponseGetNode
	if err = ctl.client.Request(tctx, token, "GetNode", &req, &resp); err != nil {
		return nil, err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Kind", "Name", "Status", "Status Reason", "Updated At"})
	for _, node := range resp.Nodes {
		table.Append([]string{
			node.Cluster,
			node.Kind,
			node.Name,
			node.Status,
			node.StatusReason,
			fmt.Sprint(time.Unix(node.UpdatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return &resp, nil
}
