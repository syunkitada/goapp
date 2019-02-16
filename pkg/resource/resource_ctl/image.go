package resource_ctl

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseImage struct {
	Images []resource_api_grpc_pb.Image
	Tctx   authproxy_grpc_pb.TraceContext
}

var getImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Show images",
	Long: `Show images
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var target string
		if len(args) > 1 {
			target = args[0]
		} else {
			target = ""
		}

		ctl := New(&config.Conf, nil)
		if err := ctl.GetImage(getCmdClusterFlag, target); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteImageCmd = &cobra.Command{
	Use:   "image [image-name]",
	Short: "Delete image",
	Long: `Delete image
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctl := New(&config.Conf, nil)
		if err := ctl.DeleteImage(deleteCmdClusterFlag, args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *ResourceCtl) GetImage(cluster string, target string) error {
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
	var resp ResponseImage
	if err = ctl.client.Request(tctx, token, "GetImage", &req, &resp); err != nil {
		return err
	}

	ctl.outputImage(&resp)

	return nil
}

func (ctl *ResourceCtl) CreateImage(tctx *logger.TraceContext, token *authproxy_client.ResponseIssueToken, spec string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := resource_api_grpc_pb.ActionRequest{
		Spec: spec,
	}
	var resp ResponseImage
	if err = ctl.client.Request(tctx, token, "CreateImage", &req, &resp); err != nil {
		return err
	}

	ctl.outputImage(&resp)

	return nil
}

func (ctl *ResourceCtl) UpdateImage(tctx *logger.TraceContext, token *authproxy_client.ResponseIssueToken, spec string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := resource_api_grpc_pb.ActionRequest{
		Spec: spec,
	}
	var resp ResponseImage
	if err = ctl.client.Request(tctx, token, "UpdateImage", &req, &resp); err != nil {
		return err
	}

	ctl.outputImage(&resp)
	return nil
}

func (ctl *ResourceCtl) DeleteImage(cluster string, target string) error {
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
	var resp ResponseImage
	if err = ctl.client.Request(tctx, token, "DeleteImage", &req, &resp); err != nil {
		return err
	}

	ctl.outputImage(&resp)
	return nil
}

func (ctl *ResourceCtl) outputImage(resp *ResponseImage) {
	if resp.Tctx.StatusCode != codes.Ok {
		fmt.Printf("Failed %s: %s\n", resp.Tctx.ActionName, resp.Tctx.Err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, image := range resp.Images {
		table.Append([]string{
			image.Cluster,
			image.Name,
			image.Status,
			image.StatusReason,
			fmt.Sprint(time.Unix(image.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(image.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()
}
