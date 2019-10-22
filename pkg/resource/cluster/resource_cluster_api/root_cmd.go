package resource_cluster_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/server"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

var RootCmd = &cobra.Command{
	Use:   "cluster-api",
	Short: "cluster-api",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&config.BaseConf, &config.MainConf)
		srv.StartMainLoop()
		srv.Serve()
	},
}
