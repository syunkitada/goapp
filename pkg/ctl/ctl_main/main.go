package ctl_main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/ctl/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var flagMap map[string]string

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		ctl := New(&config.BaseConf, &config.MainConf)
		if err := ctl.index(args); err != nil {
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
	rootCmd.Flags().SetInterspersed(false)

	base_config.InitFlags(rootCmd, &config.BaseConf)
	cobra.OnInitialize(initMain)
}

func initMain() {
	os.Setenv("LANG", "en_US.UTF-8")
	config.BaseConf.BaseDir = filepath.Join(os.Getenv("HOME"), ".goapp")
	config.BaseConf.LogTimeFormat = "2006-01-02T15:04:05Z09:00"
	base_config.InitConfig(&config.BaseConf, &config.MainConf)
	logger.InitLogger(&config.BaseConf)
}
