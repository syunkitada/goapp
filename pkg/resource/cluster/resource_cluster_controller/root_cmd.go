package resource_cluster_controller

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "cluster-controller",
	Short: "controller for management all resources",
	Long: `controller for management all resources
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewResourceClusterControllerServer(&config.Conf)
		server.StartMainLoop()
		server.Serve()
	},
}
