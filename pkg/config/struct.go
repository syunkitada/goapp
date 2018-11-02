package config

import (
	"net/http"
	"path/filepath"
)

type Config struct {
	Default   DefaultConfig
	Admin     AdminConfig
	Authproxy AuthproxyConfig
	Dashboard DashboardConfig
	Resource  ResourceConfig
}

type DefaultConfig struct {
	Name              string
	ConfigDir         string
	ConfigFile        string
	EnableTest        bool
	EnableDevelop     bool
	EnableDebug       bool
	EnableDatabaseLog bool
}

type AppConfig struct {
	Name            string
	ShutdownTimeout int
	LoopInterval    int
	Grpc            GrpcConfig
}

type AdminConfig struct {
	Username    string
	Password    string
	Secret      string
	TokenSecret string
}

type AuthproxyConfig struct {
	HttpServer HttpServerConfig
	Database   DatabaseConfig
}

type DashboardConfig struct {
	HttpServer HttpServerConfig
	BuildDir   string
}

type ResourceConfig struct {
	ApiApp         *AppConfig
	ControllerGrpc GrpcConfig
	Database       DatabaseConfig
	Cluster        ClusterConfig
	ClusterMap     map[string]*ResourceClusterConfig
}

type HttpServerConfig struct {
	Listen          string
	AllowedHosts    []string
	CertFile        string
	KeyFile         string
	GracefulTimeout int
	TestHandler     http.Handler
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

type ClusterConfig struct {
	Name string
}

type ResourceClusterConfig struct {
	ApiGrpc        GrpcConfig
	ControllerGrpc GrpcConfig
	AgentGrpc      GrpcConfig
	Database       DatabaseConfig
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
		Authproxy: AuthproxyConfig{
			HttpServer: HttpServerConfig{
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
			Database: DatabaseConfig{
				Connection: "admin:adminpass@tcp(localhost:3306)/goapp_authproxy?charset=utf8&parseTime=true",
			},
		},
		Dashboard: DashboardConfig{
			HttpServer: HttpServerConfig{
				Listen: "0.0.0.0:7000",
				AllowedHosts: []string{
					"localhost:7000",
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
