package resource_controller

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller/server"
)

var RootCmd = &cobra.Command{
	Use:   "controller",
	Short: "controller",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&config.BaseConf, &config.MainConf)
		srv.StartMainLoop()
		srv.Serve()
	},
}
