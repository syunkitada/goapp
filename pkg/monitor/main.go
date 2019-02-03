package monitor

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_alert_manager"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api"
)

var rootCmd = &cobra.Command{}

// Main is monitor's main function
func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.StdoutFatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(monitor_api.RootCmd)
	rootCmd.AddCommand(monitor_agent.RootCmd)
	rootCmd.AddCommand(monitor_alert_manager.RootCmd)
}
