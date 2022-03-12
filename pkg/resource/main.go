package resource

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/ctl"
	"github.com/syunkitada/goapp/pkg/resource/resource_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller"
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

	rootCmd.AddCommand(resource_api.RootCmd)
	rootCmd.AddCommand(resource_controller.RootCmd)
	rootCmd.AddCommand(resource_cluster_api.RootCmd)
	rootCmd.AddCommand(resource_cluster_controller.RootCmd)
	rootCmd.AddCommand(resource_cluster_agent.RootCmd)

	rootCmd.AddCommand(ctl.RootCmd)
}

func initMain() {
	os.Setenv("LANG", "en_US.UTF-8")
	config.BaseConf.BaseDir = filepath.Join(os.Getenv("HOME"), ".goapp")
	config.BaseConf.LogTimeFormat = time.RFC3339
	base_config.InitConfig(&config.BaseConf, &config.MainConf)
	logger.InitLogger(&config.BaseConf)
}
