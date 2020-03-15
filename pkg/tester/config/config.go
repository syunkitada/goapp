package config

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"

	authproxy_config "github.com/syunkitada/goapp/pkg/authproxy/config"
	resource_config "github.com/syunkitada/goapp/pkg/resource/config"
)

type Config struct {
	Authproxy authproxy_config.AuthproxyConfig
	Resource  resource_config.ResourceConfig
}

var BaseConf = base_config.Config{}

var MainConf = Config{
	Authproxy: authproxy_config.MainConf.Authproxy,
	Resource:  resource_config.MainConf.Resource,
}
