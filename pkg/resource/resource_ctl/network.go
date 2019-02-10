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

type ResponseGetNetwork struct {
	Networks []resource_api_grpc_pb.Network
	TraceId  string
	Err      string
}

type ResponseCreateNetwork struct {
	Network resource_api_grpc_pb.Network
	TraceId string
	Err     string
}

type ResponseUpdateNetwork struct {
	Network resource_api_grpc_pb.Network
	TraceId string
	Err     string
}

type ResponseDeleteNetwork struct {
	TraceId string
	Err     string
}

var getNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Show networks",
	Long: `Show networks
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var target string
		if len(args) > 1 {
			target = args[0]
		} else {
			target = ""
		}

		ctl := New(&config.Conf, nil)
		if err := ctl.GetNetwork(getCmdClusterFlag, target); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteNetworkCmd = &cobra.Command{
	Use:   "network [network-name]",
	Short: "Show networks",
	Long: `Show networks
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctl := New(&config.Conf, nil)
		if err := ctl.DeleteNetwork(deleteCmdClusterFlag, args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *ResourceCtl) GetNetwork(cluster string, target string) error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var token *authproxy_client.ResponseIssueToken
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return err
	}

	req := resource_api_grpc_pb.ActionRequest{
		Cluster: cluster,
		Target:  target,
	}
	var resp ResponseGetNetwork
	if err = ctl.client.Request(tctx, token, "GetNetwork", &req, &resp); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, network := range resp.Networks {
		table.Append([]string{
			network.Cluster,
			network.Name,
			network.Status,
			network.StatusReason,
			fmt.Sprint(time.Unix(network.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(network.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}

func (ctl *ResourceCtl) CreateNetwork(tctx *logger.TraceContext, token *authproxy_client.ResponseIssueToken, spec string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := resource_api_grpc_pb.ActionRequest{
		Spec: spec,
	}
	var resp ResponseCreateNetwork
	if err = ctl.client.Request(tctx, token, "CreateNetwork", &req, &resp); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Network.Cluster,
		resp.Network.Name,
		resp.Network.Status,
		resp.Network.StatusReason,
		fmt.Sprint(time.Unix(resp.Network.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Network.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func (ctl *ResourceCtl) UpdateNetwork(tctx *logger.TraceContext, token *authproxy_client.ResponseIssueToken, spec string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := resource_api_grpc_pb.ActionRequest{
		Spec: spec,
	}
	var resp ResponseUpdateNetwork
	if err = ctl.client.Request(tctx, token, "UpdateNetwork", &req, &resp); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Network.Cluster,
		resp.Network.Name,
		resp.Network.Status,
		resp.Network.StatusReason,
		fmt.Sprint(time.Unix(resp.Network.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Network.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func (ctl *ResourceCtl) DeleteNetwork(cluster string, target string) error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var token *authproxy_client.ResponseIssueToken
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return err
	}

	req := resource_api_grpc_pb.ActionRequest{
		Cluster: cluster,
		Target:  target,
	}
	var resp ResponseDeleteNetwork
	if err = ctl.client.Request(tctx, token, "DeleteNetwork", &req, &resp); err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
