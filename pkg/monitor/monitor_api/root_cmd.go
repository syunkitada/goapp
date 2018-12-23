package monitor_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "This api monitoring data",
	Long: `This api monitoring data
	`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := NewMonitorApiServer(&config.Conf)
		if err := srv.StartMainLoop(); err != nil {
			logger.Fatal(srv.Host, srv.Name, err)
		}

		if err := srv.Serve(); err != nil {
			logger.Fatal(srv.Host, srv.Name, err)
		}
	},
}
