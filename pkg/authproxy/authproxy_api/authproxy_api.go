package authproxy_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/server"
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

var baseConf base_config.Config
var appConf config.Config

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "api",
	Run: func(cmd *cobra.Command, args []string) {
		srv := server.New(&baseConf, &appConf)
		srv.Serve()
	},
}
