package base_config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
)

type Config struct {
	BaseDir           string
	Host              string
	ConfigDir         string
	ConfigFile        string
	TmpDir            string
	VarDir            string
	LogDir            string
	LogTimeFormat     string
	EnableTest        bool
	EnableDevelop     bool
	EnableDebug       bool
	EnableDatabaseLog bool
}

type AppConfig struct {
	Name               string
	ClientTimeout      int
	ShutdownTimeout    int
	LoopInterval       int
	Listen             string
	HttpListen         string
	CertFile           string
	KeyFile            string
	CaFile             string
	ServerHostOverride string
	Targets            []string
	Labels             []string
}

func InitFlags(rootCmd *cobra.Command, config *Config) {
	rootCmd.PersistentFlags().StringVar(&config.BaseDir, "base-dir", "", "application base directory")
	rootCmd.PersistentFlags().StringVar(&config.ConfigDir, "config-dir", "", "config directory")
	rootCmd.PersistentFlags().StringVar(&config.ConfigFile, "config-file", "", "config file")
	rootCmd.PersistentFlags().StringVar(&config.LogDir, "log-dir", "", "log directory")
	rootCmd.PersistentFlags().StringVar(&config.TmpDir, "tmp-dir", "", "tmp directory")
	rootCmd.PersistentFlags().StringVar(&config.VarDir, "var-dir", "", "var directory")
	rootCmd.PersistentFlags().BoolVar(&config.EnableDebug, "debug", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVar(&config.EnableDevelop, "develop", false, "enable develop mode")
	rootCmd.PersistentFlags().BoolVar(&config.EnableDatabaseLog, "database-log", true, "enable database logging")
}

func InitConfig(config *Config, appConfig interface{}) {
	os_utils.MustMkdir(config.BaseDir, 0755)

	if config.ConfigDir == "" {
		config.ConfigDir = filepath.Join(config.BaseDir, "etc")
	}
	os_utils.MustMkdir(config.ConfigDir, 0755)

	if config.ConfigFile == "" {
		config.ConfigFile = filepath.Join(config.ConfigDir, "config.yaml")
	}
	os_utils.MustMkdir(config.ConfigFile, 0755)

	if config.TmpDir == "" {
		config.TmpDir = filepath.Join(config.BaseDir, "tmp")
	}
	os_utils.MustMkdir(config.TmpDir, 0755)

	if config.LogDir == "" {
		config.LogDir = filepath.Join(config.BaseDir, "log")
	}
	os_utils.MustMkdir(config.LogDir, 0755)

	if config.VarDir == "" {
		config.VarDir = filepath.Join(config.BaseDir, "var")
	}
	os_utils.MustMkdir(config.VarDir, 0755)

	if config.Host == "" {
		host, err := os.Hostname()
		if err != nil {
			logger.StdoutFatal(err)
		}
		config.Host = host
	}

	if err := json_utils.ReadFile(config.ConfigFile, appConfig); err != nil {
		logger.StdoutFatalf("Failed ReadFile: path=%s, err=%v", config.ConfigFile, err)
	}
}
