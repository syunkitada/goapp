package config

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Config struct {
	Authproxy AuthproxyConfig
}

type AuthproxyConfig struct {
	Api base_config.AppConfig
}

var (
	BaseConf base_config.Config
)

var MainConf = Config{
	Authproxy: AuthproxyConfig{
		Api: base_config.AppConfig{
			Name:                 "AuthproxyApi",
			NodeDownTimeDuration: 60,
		},
	},
}
