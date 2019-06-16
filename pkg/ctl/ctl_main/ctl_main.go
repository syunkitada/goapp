package ctl_main

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/config"
)

type CtlMain struct {
	name   string
	conf   *config.Config
	client *authproxy_client.AuthproxyClient
}

func New(conf *config.Config) *CtlMain {
	return &CtlMain{
		name:   "ctl",
		conf:   conf,
		client: authproxy_client.New(conf, nil),
	}
}
