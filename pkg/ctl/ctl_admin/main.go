package ctl_admin

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/ctl/ctl_admin/resource"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var RootCmd = &cobra.Command{
	Use:   "goapp-adminctl",
	Short: "goapp-adminctl is command line interface for running command to API",
	Long: `goapp-adminctl is command line interface for running command to API
	`,
}

func Main() {
	if err := RootCmd.Execute(); err != nil {
		logger.StdoutFatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(RootCmd)

	RootCmd.AddCommand(resource.RootCmd)
}
