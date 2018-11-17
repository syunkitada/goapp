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
	GetCmdClusterFlag string
)

func init() {
	GetCmd.PersistentFlags().StringVarP(&GetCmdClusterFlag, "cluster", "c", "all", "Filtering by cluster")

	GetCmd.AddCommand(GetClusterCmd)
	GetCmd.AddCommand(GetNodeCmd)
	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show resource",
	Long: `Show resource
	`,
}

var GetClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Show clusters",
	Long: `Show clusters
	`,
	Run: func(cmd *cobra.Command, args []string) {
		authproxy := core.NewAuthproxy(&config.Conf)
		token, err := authproxy.Auth.CtlIssueToken()
		if err != nil {
			glog.Fatal(err)
		}

		resp, err := authproxy.Resource.CtlGetCluster(token.Token)
		if err != nil {
			glog.Info(resp)
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
	},
}

var GetNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Show nodes",
	Long: `Show nodes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		authproxy := core.NewAuthproxy(&config.Conf)
		token, err := authproxy.Auth.CtlIssueToken()
		if err != nil {
			glog.Fatal(err)
		}

		resp, err := authproxy.Resource.CtlGetNode(token.Token, GetCmdClusterFlag, "%")
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
	},
}
