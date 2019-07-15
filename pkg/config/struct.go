package config

import (
	"net/http"
	"path/filepath"
)

type Config struct {
	Default   DefaultConfig
	Admin     AdminConfig
	Ctl       CtlConfig
	Authproxy AuthproxyConfig
	Dashboard DashboardConfig
	Resource  ResourceConfig
	Monitor   MonitorConfig
}

type DefaultConfig struct {
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
	CertFile           string
	KeyFile            string
	CaFile             string
	ServerHostOverride string
	Targets            []string
	Labels             []string
}

type AdminConfig struct {
	Username    string
	Password    string
	Secret      string
	TokenSecret string
}

type CtlConfig struct {
	Username string
	Password string
	Project  string
	ApiUrl   string
}

type AuthproxyConfig struct {
	HttpServer HttpServerConfig
	Database   DatabaseConfig
}

type DashboardConfig struct {
	HttpServer HttpServerConfig
	BuildDir   string
}

type HttpServerConfig struct {
	Listen                   string
	AllowedHosts             []string
	AccessControlAllowOrigin string
	CertFile                 string
	KeyFile                  string
	GracefulTimeout          int
	TestHandler              http.Handler
}

type DatabaseConfig struct {
	Connection string
}

func newConfig(defaultConfig *DefaultConfig) *Config {
	config := &Config{
		Default: *defaultConfig,
		Admin: AdminConfig{
			Username:    "admin",
			Password:    "changeme",
			Secret:      "changeme",
			TokenSecret: "changeme",
		},
		Ctl: CtlConfig{
			Username: "admin",
			Password: "changeme",
			Project:  "admin",
			ApiUrl:   "https://127.0.0.1:8000",
		},
		Authproxy: AuthproxyConfig{
			HttpServer: HttpServerConfig{
				Listen:                   "0.0.0.0:8000",
				AllowedHosts:             []string{"127.0.0.1:8000"},
				AccessControlAllowOrigin: "127.0.0.1:3000",
				CertFile:                 "tls-assets/server.pem",
				KeyFile:                  "tls-assets/server.key",
				GracefulTimeout:          10,
			},
			Database: DatabaseConfig{
				Connection: "goapp:goapppass@tcp(127.0.0.1:3306)/goapp_authproxy?charset=utf8&parseTime=true",
			},
		},
		Dashboard: DashboardConfig{
			HttpServer: HttpServerConfig{
				Listen: "0.0.0.0:7000",
				AllowedHosts: []string{
					"127.0.0.1:7000",
					"192.168.10.103:7000",
				},
				CertFile:        "tls-assets/server.pem",
				KeyFile:         "tls-assets/server.key",
				GracefulTimeout: 10,
			},
			BuildDir: filepath.Join(configDir, "dashboard/build"),
		},
	}

	return config
}

func (conf *Config) Path(path string) string {
	return filepath.Join(conf.Default.ConfigDir, path)
}

func (conf *Config) LogPath(path string) string {
	return filepath.Join(conf.Default.LogDir, path)
}
