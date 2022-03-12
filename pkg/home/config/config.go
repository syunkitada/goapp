package config

import (
	authproxy_config "github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Config struct {
	Authproxy authproxy_config.AuthproxyConfig
	Home      HomeConfig
}

type HomeConfig struct {
	Api        base_config.AppConfig
	Controller HomeControllerConfig
}

type HomeControllerConfig struct {
	base_config.AppConfig
	SyncRegionServiceTimeout int
}

var BaseConf = base_config.Config{}

var MainConf = Config{
	Home: HomeConfig{
		Api: base_config.AppConfig{
			Name:                        "HomeApi",
			NodeServiceDownTimeDuration: 60,
		},
		Controller: HomeControllerConfig{
			AppConfig: base_config.AppConfig{
				Name: "HomeController",
			},
			SyncRegionServiceTimeout: 10,
		},
	},
}
