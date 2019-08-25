package authproxy_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/server"
	"github.com/syunkitada/goapp/pkg/authproxy/config"
)

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&config.BaseConf, &config.MainConf)
		srv.Serve()
	},
}
