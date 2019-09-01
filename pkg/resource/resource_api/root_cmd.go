package resource_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/server"
)

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&config.BaseConf, &config.MainConf)
		srv.StartMainLoop()
		srv.Serve()
	},
}
