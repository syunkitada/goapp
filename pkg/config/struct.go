package config

import (
	"github.com/urfave/cli"
)

type Config struct {
	App      *cli.App
	Api      ApiConfig
	Agent    AgentConfig
	Database DatabaseConfig
	Admin    AdminConfig
}

type ApiConfig struct {
	Listen   string
	Interval int
}

type AgentConfig struct {
	ReportInterval int
}

type DatabaseConfig struct {
	Connection string
}

type AdminConfig struct {
	Username    string
	Password    string
	Secret      string
	TokenSecret string
}

func newConfig(ctx *cli.Context) *Config {
	defaultConfig := &Config{
		App: ctx.App,
		Api: ApiConfig{
			Listen:   "0.0.0.0:8080",
			Interval: 10,
		},
		Agent: AgentConfig{
			ReportInterval: 10,
		},
		Database: DatabaseConfig{
			Connection: "admin:adminpass@tcp(localhost:3306)/goapp?charset=utf8&parseTime=true",
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
