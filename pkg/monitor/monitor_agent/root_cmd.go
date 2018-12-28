package monitor_agent

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var RootCmd = &cobra.Command{
	Use:   "agent",
	Short: "agent",
	Long: `agent
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewMonitorAgentServer(&config.Conf)
		if err := server.StartMainLoop(); err != nil {
			logger.StdoutFatal(err)
		}

		if err := server.Serve(); err != nil {
			logger.StdoutFatal(err)
		}
	},
}
