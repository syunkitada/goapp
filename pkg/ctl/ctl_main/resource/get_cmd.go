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
)

var (
	getCmdClusterFlag string
)

func init() {
	GetCmd.PersistentFlags().StringVarP(&getCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")

	GetCmd.AddCommand(getComputeCmd)
	GetCmd.AddCommand(getClusterCmd)
	GetCmd.AddCommand(getNetworkV4Cmd)
	GetCmd.AddCommand(getNodeCmd)
	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show resource",
	Long: `Show resource
	`,
}

var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Show clusters",
	Long: `Show clusters
	`,
	Run: func(cmd *cobra.Command, args []string) {
		GetCluster()
	},
}

var getNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Show nodes",
	Long: `Show nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		GetNode()
	},
}

var getComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Show computes",
	Long: `Show computes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		GetCompute()
	},
}

func GetCluster() {
	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		glog.Fatal(err)
	}

	resp, err := authproxy.Resource.CtlGetCluster(token.Token)
	if err != nil {
		glog.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})
	for _, cluster := range resp.Clusters {
		table.Append([]string{
			cluster.Name,
		})
	}
	table.Render()
}

func GetNode() {
	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		glog.Fatal(err)
	}

	resp, err := authproxy.Resource.CtlGetNode(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		glog.Info(resp)
		glog.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Kind", "Role", "Status", "Status Reason", "State", "State Reason", "Updated At"})
	for _, node := range resp.Nodes {
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
}

func GetCompute() {
	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		glog.Fatal(err)
	}

	resp, err := authproxy.Resource.CtlGetCompute(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		glog.Info(resp)
		glog.Fatal(err)
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
