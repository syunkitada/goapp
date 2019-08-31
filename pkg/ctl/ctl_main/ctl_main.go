package ctl_main

import (
	"github.com/syunkitada/goapp/pkg/authproxy/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/ctl/config"
)

type Ctl struct {
	name     string
	baseConf *base_config.Config
	mainConf *config.Config
	client   *genpkg.Client
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Ctl {
	return &Ctl{
		name:     "ctl",
		baseConf: baseConf,
		mainConf: mainConf,
		client: genpkg.NewClient(&base_config.ClientConfig{
			Targets:               []string{"https://localhost:8000"},
			TlsInsecureSkipVerify: true,
		}),
	}
}
