package resource_cluster_agent

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/server"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

var RootCmd = &cobra.Command{
	Use:   "cluster-agent",
	Short: "cluster-agent",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&config.BaseConf, &config.MainConf)
		srv.StartMainLoop()
		srv.Serve()
	},
}
