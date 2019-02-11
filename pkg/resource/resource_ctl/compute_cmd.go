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
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseCompute struct {
	Computes []resource_api_grpc_pb.Compute
	Tctx     resource_api_grpc_pb.TraceContext
}

var getComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Show computes",
	Long: `Show computes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var target string
		if len(args) > 1 {
			target = args[0]
		} else {
			target = ""
		}

		ctl := New(&config.Conf, nil)
		if err := ctl.GetCompute(getCmdClusterFlag, target); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteComputeCmd = &cobra.Command{
	Use:   "compute [compute-name]",
	Short: "Delete compute",
	Long: `Delete compute
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctl := New(&config.Conf, nil)
		if err := ctl.DeleteCompute(deleteCmdClusterFlag, args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *ResourceCtl) GetCompute(cluster string, target string) error {
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
	var resp ResponseCompute
	if err = ctl.client.Request(tctx, token, "GetCompute", &req, &resp); err != nil {
		return err
	}

	ctl.outputCompute(&resp)

	return nil
}

func (ctl *ResourceCtl) CreateCompute(tctx *logger.TraceContext, token *authproxy_client.ResponseIssueToken, spec string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := resource_api_grpc_pb.ActionRequest{
		Spec: spec,
	}
	var resp ResponseCompute
	if err = ctl.client.Request(tctx, token, "CreateCompute", &req, &resp); err != nil {
		return err
	}

	ctl.outputCompute(&resp)

	return nil
}

func (ctl *ResourceCtl) UpdateCompute(tctx *logger.TraceContext, token *authproxy_client.ResponseIssueToken, spec string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := resource_api_grpc_pb.ActionRequest{
		Spec: spec,
	}
	var resp ResponseCompute
	if err = ctl.client.Request(tctx, token, "UpdateCompute", &req, &resp); err != nil {
		return err
	}

	ctl.outputCompute(&resp)
	return nil
}

func (ctl *ResourceCtl) DeleteCompute(cluster string, target string) error {
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
	var resp ResponseCompute
	if err = ctl.client.Request(tctx, token, "DeleteCompute", &req, &resp); err != nil {
		return err
	}

	ctl.outputCompute(&resp)
	return nil
}

func (ctl *ResourceCtl) outputCompute(resp *ResponseCompute) {
	if resp.Tctx.StatusCode != codes.Ok {
		fmt.Printf("Failed %s: %s\n", resp.Tctx.ActionName, resp.Tctx.Err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, compute := range resp.Computes {
		table.Append([]string{
			compute.Cluster,
			compute.Name,
			compute.Status,
			compute.StatusReason,
			fmt.Sprint(time.Unix(compute.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(compute.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()
}
