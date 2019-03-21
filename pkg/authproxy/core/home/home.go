package home

import (
	"github.com/syunkitada/goapp/pkg/config"
)

type Home struct {
	host string
	name string
	path string
	conf *config.Config
}

func New(conf *config.Config) *Home {
	srv := Home{
		host: conf.Default.Host,
		name: "authproxy:home",
		path: "/Home",
		conf: conf,
	}
	return &srv
}

func (srv *Home) GetPath() string {
	return srv.path
}
