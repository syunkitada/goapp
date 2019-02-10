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

type ResponseGetCluster struct {
	Clusters []resource_api_grpc_pb.Cluster
	TraceId  string
	Err      string
}

var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Show clusters",
	Long: `Show clusters
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var target string
		if len(args) > 1 {
			target = args[0]
		} else {
			target = ""
		}

		ctl := New(&config.Conf, nil)
		if _, err := ctl.GetCluster(target); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *ResourceCtl) GetCluster(target string) (*ResponseGetCluster, error) {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var token *authproxy_client.ResponseIssueToken
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return nil, err
	}

	req := resource_api_grpc_pb.ActionRequest{
		Target: target,
	}

	var resp ResponseGetCluster
	if err = ctl.client.Request(tctx, token, "GetCluster", &req, &resp); err != nil {
		return nil, err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Updated At", "Created At"})
	for _, cluster := range resp.Clusters {
		table.Append([]string{
			cluster.Name,
			fmt.Sprint(time.Unix(cluster.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(cluster.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return &resp, nil
}
