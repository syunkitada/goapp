package home

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/home/config"
	"github.com/syunkitada/goapp/pkg/home/ctl"
	"github.com/syunkitada/goapp/pkg/home/home_api"
	"github.com/syunkitada/goapp/pkg/home/home_controller"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.StdoutFatal(err)
	}
}

func init() {
	base_config.InitFlags(rootCmd, &config.BaseConf)
	cobra.OnInitialize(initMain)

	rootCmd.AddCommand(home_api.RootCmd)
	rootCmd.AddCommand(home_controller.RootCmd)

	rootCmd.AddCommand(ctl.RootCmd)
}

func initMain() {
	os.Setenv("LANG", "en_US.UTF-8")
	config.BaseConf.BaseDir = filepath.Join(os.Getenv("HOME"), ".goapp")
	// config.BaseConf.LogTimeFormat = "2006-01-02T15:04:05Z09:00"
	config.BaseConf.LogTimeFormat = time.RFC3339
	base_config.InitConfig(&config.BaseConf, &config.MainConf)
	logger.InitLogger(&config.BaseConf)
}
