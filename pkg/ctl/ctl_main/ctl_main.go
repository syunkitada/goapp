package ctl_main

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/ctl/config"
)

type Ctl struct {
	name     string
	baseConf *base_config.Config
	mainConf *config.Config
	client   *base_client.Client
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Ctl {
	return &Ctl{
		name:     "ctl",
		baseConf: baseConf,
		mainConf: mainConf,
		client: base_client.NewClient(&base_config.ClientConfig{
			Endpoints:             []string{"https://localhost:8000"},
			TlsInsecureSkipVerify: true,
		}),
	}
}
