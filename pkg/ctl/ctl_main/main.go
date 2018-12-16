package ctl_main

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/ctl/ctl_main/resource"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(resource.RootCmd)
}
