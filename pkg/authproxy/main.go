package authproxy

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_api"
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/ctl"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		logger.StdoutFatalf("Failed Execute: %v", err)
	}
}

func init() {
	base_config.InitFlags(rootCmd, &config.BaseConf)
	cobra.OnInitialize(initMain)

	rootCmd.AddCommand(authproxy_api.RootCmd)
	rootCmd.AddCommand(ctl.RootCmd)
}

func initMain() {
	os.Setenv("LANG", "en_US.UTF-8")
	config.BaseConf.BaseDir = filepath.Join(os.Getenv("HOME"), ".goapp")
	config.BaseConf.LogTimeFormat = "2006-01-02T15:04:05Z09:00"
	base_config.InitConfig(&config.BaseConf, &config.MainConf)
	logger.InitLogger(&config.BaseConf)
}
