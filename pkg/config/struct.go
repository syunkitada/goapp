package config

import (
	"net/http"
	"os"
	"strings"

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
	Resource          ResourceConfig
	Agent             AgentConfig
	ImageDatabase     DatabaseConfig
	Admin             AdminConfig
}

type DefaultConfig struct {
	ConfigDir  string
	ConfigFile string
	TestMode   bool
}

type HttpServerConfig struct {
	Listen          string
	AllowedHosts    []string
	CertFile        string
	KeyFile         string
	GracefulTimeout int
	TestHandler     http.Handler
}

type AgentConfig struct {
	ReportInterval int
}

type DatabaseConfig struct {
	Connection string
}

type DashboardConfig struct {
	HttpServerConfig
	BuildDir string
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

type ResourceConfig struct {
	Database DatabaseConfig
	Grpc     GrpcConfig
}

func newConfig(ctx *cli.Context) *Config {
	configDir := "/etc/goapp"
	dashboardBuildDir := "/opt/goapp/dashboard/build"
	if ctx.GlobalBool("use-pwd") {
		pwd := os.Getenv("PWD")
		splitedPwd := strings.Split(pwd, "/pkg/")
		configDir = splitedPwd[0] + "/testdata"
		dashboardBuildDir = splitedPwd[0] + "/dashboard/build"
	}

	defaultConfig := &Config{
		App: ctx.App,
		Default: DefaultConfig{
			ConfigDir:  configDir,
			ConfigFile: filepath.Join(configDir, "config.toml"),
			TestMode:   ctx.GlobalBool("test-mode"),
		},
		Authproxy: HttpServerConfig{
			Listen: "0.0.0.0:8000",
			AllowedHosts: []string{
				"localhost:8000",
				"192.168.10.103:3000",
				"192.168.10.103:8000",
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
			BuildDir: dashboardBuildDir,
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
		Resource: ResourceConfig{
			Grpc: GrpcConfig{
				Listen:             "localhost:13300",
				CertFile:           "server1.pem",
				KeyFile:            "server1.key",
				CaFile:             "ca.pem",
				ServerHostOverride: "x.test.youtube.com",
				Targets: []string{
					"localhost:10080",
				},
			},
			Database: DatabaseConfig{
				Connection: "admin:adminpass@tcp(localhost:3306)/resource?charset=utf8&parseTime=true",
			},
		},
	}

	return defaultConfig
}