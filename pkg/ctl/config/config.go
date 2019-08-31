package config

import "github.com/syunkitada/goapp/pkg/base/base_config"

type Config struct {
	Ctl CtlConfig
}

type CtlConfig struct {
	Project string
}

var (
	BaseConf base_config.Config
	MainConf Config
)
