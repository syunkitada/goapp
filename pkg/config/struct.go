package config

import (
	"os"

	"github.com/urfave/cli"
	"path/filepath"
)

type Config struct {
	App               *cli.App
	Default           DefaultConfig
	Authproxy         HttpServerConfig
	AuthproxyDatabase DatabaseConfig
	Dashboard         DashboardConfig
	HealthGrpc        GrpcConfig
	Agent             AgentConfig
	ImageDatabase     DatabaseConfig
	Admin             AdminConfig
}

type DefaultConfig struct {
	ConfigDir  string
	ConfigFile string
}

type HttpServerConfig struct {
	Listen          string
	AllowedHosts    []string
	CertFile        string
	KeyFile         string
	GracefulTimeout int
}

type AgentConfig struct {
	ReportInterval int
}

type DatabaseConfig struct {
	Connection string
}

type DashboardConfig struct {
	HttpServerConfig
	StaticDir    string
	TemplatesDir string
}

type GrpcConfig struct {
	Listen             string
	CertFile           string
	KeyFile            string
	CaFile             string
	ServerHostOverride string
	Targets            []string
}

type AdminConfig struct {
	Username    string
	Password    string
	Secret      string
	TokenSecret string
}

func newConfig(ctx *cli.Context) *Config {
	configDir := "/etc/goapp"
	dashboardStaticDir := "/opt/goapp/htdocs/static"
	dashboardTemplatesDir := "/opt/goapp/htdocs/templates/**/*"
	if ctx.GlobalBool("use-pwd") {
		pwd := os.Getenv("PWD")
		configDir = pwd + "/testdata"
		dashboardStaticDir = pwd + "/htdocs/static"
		dashboardTemplatesDir = pwd + "/htdocs/templates/**/*"
	}

	defaultConfig := &Config{
		App: ctx.App,
		Default: DefaultConfig{
			ConfigDir:  configDir,
			ConfigFile: filepath.Join(configDir, "config.toml"),
		},
		Authproxy: HttpServerConfig{
			Listen: "0.0.0.0:8000",
			AllowedHosts: []string{
				"localhost:8000",
				"192.168.10.103:7000",
			},
			CertFile:        "tls-assets/server.pem",
			KeyFile:         "tls-assets/server.key",
			GracefulTimeout: 10,
		},
		AuthproxyDatabase: DatabaseConfig{
			Connection: "admin:adminpass@tcp(localhost:3306)/goapp?charset=utf8&parseTime=true",
		},
		Dashboard: DashboardConfig{
			HttpServerConfig: HttpServerConfig{
				Listen: "0.0.0.0:7000",
				AllowedHosts: []string{
					"localhost:7000",
					"192.168.10.103:7000",
				},
				CertFile:        "tls-assets/server.pem",
				KeyFile:         "tls-assets/server.key",
				GracefulTimeout: 10,
			},
			StaticDir:    dashboardStaticDir,
			TemplatesDir: dashboardTemplatesDir,
		},
		Agent: AgentConfig{
			ReportInterval: 10,
		},
		ImageDatabase: DatabaseConfig{
			Connection: "admin:adminpass@tcp(localhost:3306)/goapp?charset=utf8&parseTime=true",
		},
		HealthGrpc: GrpcConfig{
			Listen:             "localhost:10080",
			CertFile:           "server1.pem",
			KeyFile:            "server1.key",
			CaFile:             "ca.pem",
			ServerHostOverride: "x.test.youtube.com",
			Targets: []string{
				"localhost:10080",
			},
		},
		Admin: AdminConfig{
			Username:    "admin",
			Password:    "changeme",
			Secret:      "changeme",
			TokenSecret: "changeme",
		},
	}

	return defaultConfig
}
