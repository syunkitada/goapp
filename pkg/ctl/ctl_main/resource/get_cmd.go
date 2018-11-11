package resource

import (
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func init() {
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

		resp, err := authproxy.Resource.CtlGetNode(token.Token)
		if err != nil {
			glog.Info(resp)
			glog.Fatal(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name"})
		for _, node := range resp.Nodes {
			table.Append([]string{
				node.Name,
			})
		}
		table.Render()
	},
}
