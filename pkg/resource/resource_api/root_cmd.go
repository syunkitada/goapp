package resource_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "This is api for controlle all resources",
	Long: `This is api for controlle all resources
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewResourceApiServer(&config.Conf)
		if err := server.StartMainLoop(); err != nil {
			logger.Fatal(server.Host, server.Name, err)
		}

		if err := server.Serve(); err != nil {
			logger.Fatal(server.Host, server.Name, err)
		}
	},
}
