package config

import (
	"github.com/urfave/cli"
	"path/filepath"
)

type Config struct {
	App               *cli.App
	Default           DefaultConfig
	Authproxy         AuthproxyConfig
	AuthproxyDatabase DatabaseConfig
	Agent             AgentConfig
	ImageDatabase     DatabaseConfig
	Admin             AdminConfig
	HealthGrpc        GrpcConfig
}

type DefaultConfig struct {
	ConfigDir  string
	ConfigFile string
}

type AuthproxyConfig struct {
	Listen       string
	Interval     int
	AllowedHosts []string
	CertFile     string
	KeyFile      string
}

type AgentConfig struct {
	ReportInterval int
}

type DatabaseConfig struct {
	Connection string
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

func newConfig(ctx *cli.Context, configDir string) *Config {
	defaultConfig := &Config{
		App: ctx.App,
		Default: DefaultConfig{
			ConfigDir:  configDir,
			ConfigFile: filepath.Join(configDir, "config.toml"),
		},
		Authproxy: AuthproxyConfig{
			Listen: "0.0.0.0:8000",
			AllowedHosts: []string{
				"localhost:8000",
			},
			Interval: 10,
			CertFile: "tls-assets/server.pem",
			KeyFile:  "tls-assets/server.key",
		},
		Agent: AgentConfig{
			ReportInterval: 10,
		},
		AuthproxyDatabase: DatabaseConfig{
			Connection: "admin:adminpass@tcp(localhost:3306)/goapp?charset=utf8&parseTime=true",
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
