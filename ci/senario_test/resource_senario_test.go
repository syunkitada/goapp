package senario_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/spf13/cobra"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/tester"

	"github.com/syunkitada/goapp/pkg/tester/config"
)

func TestResourceSenario(t *testing.T) {
	appName := "goapp-test"
	configFile := "config-test.yaml"
	os.Args = []string{appName, "--config-file", configFile, "--test"}

	var rootCmd = &cobra.Command{
		Use: appName,
		Run: func(cmd *cobra.Command, args []string) {
			tester := tester.NewTester(t, &config.BaseConf, &config.MainConf)
			tester.TestResourceSenario()
		},
	}

	base_config.InitFlags(rootCmd, &config.BaseConf)
	cobra.OnInitialize(func() {
		os.Setenv("LANG", "en_US.UTF-8")
		config.BaseConf.BaseDir = filepath.Join(os.Getenv("HOME"), ".goapp")
		config.BaseConf.LogTimeFormat = time.RFC3339
		base_config.InitConfig(&config.BaseConf, &config.MainConf)
		logger.InitLogger(&config.BaseConf)
	})

	if err := rootCmd.Execute(); err != nil {
		logger.Logger.Fatalf("Failed Execute: err=%s", err.Error())
	}
}
