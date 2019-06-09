package ctl_main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var flagMap map[string]string

var rootCmd = &cobra.Command{
	Use:                "",
	Short:              "",
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := New(&config.Conf)
		if err := ctl.Index(args); err != nil {
			fmt.Println(err)
		}
	},
}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.StdoutFatal(err)
	}
}

func init() {
	// flagMap
	// rootCmd.PersistentFlags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

	cobra.OnInitialize(config.InitConfig, logger.Init)
	config.InitFlags(rootCmd)

	// rootCmd.AddCommand(resource_ctl.RootCmd)
}
