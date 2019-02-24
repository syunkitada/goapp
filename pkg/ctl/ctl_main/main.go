package ctl_main

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/ctl/ctl_main/monitor"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_ctl"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.StdoutFatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(resource_ctl.RootCmd)
	rootCmd.AddCommand(monitor.RootCmd)
}
