package monitor_alert_manager

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var RootCmd = &cobra.Command{
	Use:   "alert-manager",
	Short: "alert-manager",
	Long: `alert-manager
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewMonitorAlertManagerServer(&config.Conf)
		if err := server.StartMainLoop(); err != nil {
			logger.Fatal(server.Host, server.Name, err)
		}

		if err := server.Serve(); err != nil {
			logger.Fatal(server.Host, server.Name, err)
		}
	},
}
