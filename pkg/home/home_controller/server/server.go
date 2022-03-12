package server

import (
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/home/config"
	"github.com/syunkitada/goapp/pkg/home/db_api"
	"github.com/syunkitada/goapp/pkg/home/home_controller/resolver"
	"github.com/syunkitada/goapp/pkg/home/home_controller/spec/genpkg"
)

type Server struct {
	base_app.BaseApp
	baseConf     *base_config.Config
	mainConf     *config.Config
	queryHandler *genpkg.QueryHandler
	dbApi        *db_api.Api
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	dbApi := db_api.New(baseConf, mainConf)
	resolver := resolver.New(baseConf, mainConf, dbApi)
	queryHandler := genpkg.NewQueryHandler(baseConf, &mainConf.Home.Controller.AppConfig, resolver)
	baseApp := base_app.New(baseConf, &mainConf.Home.Controller.AppConfig, dbApi, queryHandler)

	srv := &Server{
		BaseApp:      baseApp,
		baseConf:     baseConf,
		mainConf:     mainConf,
		queryHandler: queryHandler,
		dbApi:        dbApi,
	}
	srv.SetDriver(srv)
	return srv
}
