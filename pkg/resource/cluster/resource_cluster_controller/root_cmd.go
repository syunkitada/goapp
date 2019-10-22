package resource_cluster_controller

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller/server"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

var RootCmd = &cobra.Command{
	Use:   "cluster-controller",
	Short: "cluster-controller",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&config.BaseConf, &config.MainConf)
		srv.StartMainLoop()
		srv.Serve()
	},
}
