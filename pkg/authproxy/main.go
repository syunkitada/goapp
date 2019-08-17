package authproxy

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

var baseConfig base_config.Config
var appConfig config.Config

var rootCmd = &cobra.Command{
	Use:   "goapp-authproxy",
	Short: "goapp-authproxy",
	Long: `goapp-authproxy
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(baseConfig)
		fmt.Println(appConfig)
		// authproxy := core.NewAuthproxy(&config.Conf)
		// authproxy.Serv()
	},
}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	base_config.InitFlags(rootCmd, &baseConfig)
}

func initConfig() {
	os.Setenv("LANG", "en_US.UTF-8")
	baseConfig.BaseDir = filepath.Join(os.Getenv("HOME"), ".goapp")
	baseConfig.LogTimeFormat = "2006-01-02T15:04:05Z09:00"
	base_config.InitConfig(&baseConfig, &appConfig)
	fmt.Println("DEBUG Config")
}

func initLogger() {
	fmt.Println("DEBUG Logger")
}
