package resource

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func init() {
	GetCmd.AddCommand(GetClusterCmd)
	GetCmd.AddCommand(GetNodeCmd)
	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long: `Show resource api status
	`,
}

var GetClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "get cluster",
	Long: `get cluster
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := NewResourceCtl(&config.Conf)
		if err := ctl.GetCluster(); err != nil {
			glog.Fatal(err)
		}
	},
}

var GetNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "get node",
	Long: `get node
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := NewResourceCtl(&config.Conf)
		if err := ctl.GetNode(); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *ResourceCtl) GetCluster() error {
	reply, err := ctl.resourceApiClient.Status()
	if err != nil {
		return err
	}

	data := [][]string{
		[]string{"clusterA", reply.Msg},
		[]string{"clusterB", reply.Msg},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Status"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}

func (ctl *ResourceCtl) GetNode() error {
	req := &resource_api_grpc_pb.GetNodeRequest{
		Target: "%",
	}

	reply, err := ctl.resourceApiClient.GetNode(req)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Kind", "Role", "Status", "Status Reason", "State", "State Reason", "Updated At"})
	for _, node := range reply.Nodes {
		table.Append([]string{
			node.Name,
			node.Kind,
			node.Role,
			node.Status,
			node.StatusReason,
			node.State,
			node.StateReason,
			fmt.Sprint(time.Unix(node.UpdatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}
