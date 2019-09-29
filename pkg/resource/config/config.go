package config

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Config struct {
	Resource ResourceConfig
}

type ResourceConfig struct {
	Api        base_config.AppConfig
	Controller base_config.AppConfig
}

var BaseConf = base_config.Config{}

var MainConf = Config{
	Resource: ResourceConfig{
		Api: base_config.AppConfig{
			Name: "ResourceApi",
		},
		Controller: base_config.AppConfig{
			Name: "ReosurceController",
		},
	},
}
