package server

import (
	authproxy_config "github.com/syunkitada/goapp/pkg/authproxy/config"
	authproxy_db_api "github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/home/config"
	"github.com/syunkitada/goapp/pkg/home/db_api"
	"github.com/syunkitada/goapp/pkg/home/home_api/resolver"
	"github.com/syunkitada/goapp/pkg/home/home_api/spec/genpkg"
)

type Server struct {
	base_app.BaseApp
	baseConf       *base_config.Config
	mainConf       *config.Config
	queryHandler   *genpkg.QueryHandler
	authproxyDbApi *authproxy_db_api.Api
	dbApi          *db_api.Api
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	dbApi := db_api.New(baseConf, mainConf)
	authproxyDbApi := authproxy_db_api.New(baseConf, &authproxy_config.Config{Authproxy: mainConf.Authproxy})
	resolver := resolver.New(baseConf, mainConf, dbApi, authproxyDbApi)
	queryHandler := genpkg.NewQueryHandler(baseConf, &mainConf.Home.Api, resolver)
	baseApp := base_app.New(baseConf, &mainConf.Home.Api, dbApi, queryHandler)

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
