package server

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/resolver"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Server struct {
	base_app.BaseApp
	baseConf     *base_config.Config
	mainConf     *config.Config
	queryHandler *genpkg.QueryHandler
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	dbApi := db_api.New(baseConf, mainConf)
	resolver := resolver.New(baseConf, mainConf, dbApi)
	queryHandler := genpkg.NewQueryHandler(baseConf, &mainConf.Authproxy.Api, resolver)
	baseApp := base_app.New(baseConf, &mainConf.Authproxy.Api, dbApi, queryHandler)

	srv := &Server{
		BaseApp:  baseApp,
		baseConf: baseConf,
		mainConf: mainConf,
	}
	srv.SetDriver(srv)
	return srv
}
