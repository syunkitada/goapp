package monitor

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_alert_manager"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_proxy"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(config.Conf.Default.Host, "monitor", err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(monitor_proxy.RootCmd)
	rootCmd.AddCommand(monitor_agent.RootCmd)
	rootCmd.AddCommand(monitor_alert_manager.RootCmd)
}