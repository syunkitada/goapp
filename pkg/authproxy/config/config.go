package config

import "github.com/syunkitada/goapp/pkg/base/base_config"

type Config struct {
	Authproxy AuthproxyConfig
}

type AuthproxyConfig struct {
	App base_config.AppConfig
}
