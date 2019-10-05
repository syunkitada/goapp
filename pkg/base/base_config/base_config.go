package base_config

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
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
	Name                     string   `validate:"required"`
	ClientTimeout            int      `validate:"required"`
	ShutdownTimeout          int      `validate:"required"`
	LoopInterval             int      `validate:"required"`
	Listen                   string   `validate:"required"`
	HttpListen               string   `validate:"required"`
	Endpoints                []string `validate:"required"`
	CertFile                 string   `validate:"required"`
	KeyFile                  string   `validate:"required"`
	CaFile                   string   `validate:"required"`
	AccessControlAllowOrigin string   `validate:"required"`
	ServerHostOverride       string   `validate:"required"`
	Targets                  []string `validate:"required"`
	Labels                   []string `validate:"required"`
	NodeDownTimeDuration     int      `validate:"required"`
	Database                 DatabaseConfig
	Auth                     AuthConfig
	RootClient               ClientConfig
}

type ClientConfig struct {
	User                  string   `validate:"required"`
	Password              string   `validate:"required"`
	Service               string   `validate:"required"`
	Endpoints             []string `validate:"required"`
	TlsInsecureSkipVerify bool
	LocalHandler          http.Handler
}

type DatabaseConfig struct {
	Connection string `validate:"required"`
}

type AuthConfig struct {
	Secrets             []string `validate:"required"`
	DefaultUsers        []AuthUser
	DefaultRoles        []AuthRole
	DefaultProjects     []AuthProject
	DefaultProjectRoles []AuthProjectRole
	DefaultServices     []AuthService
}

type AuthUser struct {
	Name     string
	Password string
	Roles    []string
}

type AuthRole struct {
	Name    string
	Project string
}

type AuthProject struct {
	Name        string
	ProjectRole string
}

type AuthProjectRole struct {
	Name string
}

type AuthService struct {
	Name            string
	Scope           string
	SyncRootCluster bool
	ProjectRoles    []string
}

func InitFlags(rootCmd *cobra.Command, conf *Config) {
	rootCmd.PersistentFlags().StringVar(&conf.BaseDir, "base-dir", "", "application base directory")
	rootCmd.PersistentFlags().StringVar(&conf.ConfigDir, "config-dir", "", "config directory")
	rootCmd.PersistentFlags().StringVar(&conf.ConfigFile, "config-file", "", "config file")
	rootCmd.PersistentFlags().StringVar(&conf.LogDir, "log-dir", "", "log directory")
	rootCmd.PersistentFlags().StringVar(&conf.TmpDir, "tmp-dir", "", "tmp directory")
	rootCmd.PersistentFlags().StringVar(&conf.VarDir, "var-dir", "", "var directory")
	rootCmd.PersistentFlags().BoolVar(&conf.EnableDebug, "debug", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVar(&conf.EnableDevelop, "develop", false, "enable develop mode")
	rootCmd.PersistentFlags().BoolVar(&conf.EnableDatabaseLog, "database-log", false, "enable database logging")
}

func InitConfig(conf *Config, appConf interface{}) {
	mustMkdir(conf.BaseDir, 0755)

	if conf.ConfigDir == "" {
		conf.ConfigDir = filepath.Join(conf.BaseDir, "etc")
	}
	mustMkdir(conf.ConfigDir, 0755)

	if conf.ConfigFile == "" {
		conf.ConfigFile = filepath.Join(conf.ConfigDir, "config.yaml")
	}

	if conf.TmpDir == "" {
		conf.TmpDir = filepath.Join(conf.BaseDir, "tmp")
	}
	mustMkdir(conf.TmpDir, 0755)

	if conf.LogDir == "" {
		conf.LogDir = filepath.Join(conf.BaseDir, "log")
	}
	mustMkdir(conf.LogDir, 0755)

	if conf.VarDir == "" {
		conf.VarDir = filepath.Join(conf.BaseDir, "var")
	}
	mustMkdir(conf.VarDir, 0755)

	if conf.Host == "" {
		host, err := os.Hostname()
		if err != nil {
			log.Fatalf("Failed os.Hostname(): %v", err)
		}
		conf.Host = host
	}

	mustLoadConf(conf.ConfigFile, appConf)
}

func mustMkdir(path string, perm os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, perm); err != nil {
			log.Fatalf("Failed Mkdir: path=%s, err=%v", path, err)
		}
	} else if err != nil {
		log.Fatalf("Failed Mkdir: path=%s, err=%v", path, err)
	}
}

func mustLoadConf(filePath string, data interface{}) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed ReadFile: path=%s, err=%v", filePath, err)
	}
	err = yaml.Unmarshal(bytes, data)
}
