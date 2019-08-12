package monitor_api

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "api",
	Short: "This api monitoring data",
	Long: `This api monitoring data
	`,
	Run: func(cmd *cobra.Command, args []string) {
		srv := NewMonitorApiServer(&config.Conf)
		srv.StartMainLoop()
		srv.Serve()
	},
}
