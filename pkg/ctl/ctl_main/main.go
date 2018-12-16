package ctl_main

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/ctl/ctl_main/resource"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(config.Conf.Default.Host, "ctl", err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(resource.RootCmd)
}
