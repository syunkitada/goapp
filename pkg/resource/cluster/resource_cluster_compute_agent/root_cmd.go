package resource_cluster_compute_agent

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "cluster-compute-agent",
	Short: "agent for management resource",
	Long: `agent for management resource
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := New(&config.Conf)
		server.StartMainLoop()
		server.ServeHttp()
	},
}
