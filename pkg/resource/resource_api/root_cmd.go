package resource_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "This is api for controlle all resources",
	Long: `This is api for controlle all resources
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewResourceApiServer(&config.Conf)
		server.StartMainLoop()
		server.Serve()
	},
}
