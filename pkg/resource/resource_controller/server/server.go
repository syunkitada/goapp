package server

import (
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller/resolver"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller/spec/genpkg"
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
	queryHandler := genpkg.NewQueryHandler(baseConf, &mainConf.Resource.Controller, resolver)
	baseApp := base_app.New(baseConf, &mainConf.Resource.Controller, dbApi, queryHandler)

	srv := &Server{
		BaseApp:      baseApp,
		baseConf:     baseConf,
		mainConf:     mainConf,
		queryHandler: queryHandler,
	}
	srv.SetDriver(srv)
	return srv
}
