package config

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Config struct {
	Resource ResourceConfig
}

type ResourceConfig struct {
	App base_config.AppConfig
}

var (
	BaseConf base_config.Config
	MainConf Config
)
