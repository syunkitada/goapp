package resource_cluster_compute_agent

import (
	"github.com/golang/glog"
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
		if err := server.StartMainLoop(); err != nil {
			glog.Fatal(err)
		}

		if err := server.ServeHttp(); err != nil {
			glog.Fatal(err)
		}
	},
}
