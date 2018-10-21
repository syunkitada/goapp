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
	ConfigDir         string
	ConfigFile        string
	EnableTest        bool
	EnableDevelop     bool
	EnableDebug       bool
	EnableDatabaseLog bool
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
	Database       DatabaseConfig
	ApiGrpc        GrpcConfig
	ControllerGrpc GrpcConfig
	Region         RegionConfig
	RegionMap      map[string]*ResourceRegionConfig
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

type RegionConfig struct {
	Name string
}

type ResourceRegionConfig struct {
	Database DatabaseConfig
	ApiGrpc  GrpcConfig
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
		Resource: ResourceConfig{
			ApiGrpc: GrpcConfig{
				Listen:             "localhost:13300",
				CertFile:           "server1.pem",
				KeyFile:            "server1.key",
				CaFile:             "ca.pem",
				ServerHostOverride: "x.test.youtube.com",
				Targets: []string{
					"localhost:13300",
				},
			},
			ControllerGrpc: GrpcConfig{
				Listen:             "localhost:13301",
				CertFile:           "server1.pem",
				KeyFile:            "server1.key",
				CaFile:             "ca.pem",
				ServerHostOverride: "x.test.youtube.com",
				Targets: []string{
					"localhost:13301",
				},
			},
			Database: DatabaseConfig{
				Connection: "admin:adminpass@tcp(localhost:3306)/goapp_resource?charset=utf8&parseTime=true",
			},
			Region: RegionConfig{
				Name: "region1",
			},
			RegionMap: map[string]*ResourceRegionConfig{
				"region1": &ResourceRegionConfig{
					ApiGrpc: GrpcConfig{
						Listen:             "localhost:13310",
						CertFile:           "server1.pem",
						KeyFile:            "server1.key",
						CaFile:             "ca.pem",
						ServerHostOverride: "x.test.youtube.com",
						Targets: []string{
							"localhost:13310",
						},
					},
				},
			},
		},
	}

	return config
}

func (conf *Config) Path(path string) string {
	return filepath.Join(conf.Default.ConfigDir, path)
}
