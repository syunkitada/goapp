package server

import (
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Server struct {
	base_app.BaseApp
	baseConfig *base_config.Config
	mainConfig *config.Config
}

func New(baseConfig *base_config.Config, mainConfig *config.Config) *Server {
	baseApp := base_app.New(baseConfig, &mainConfig.Authproxy.App)
	// TODO
	// baseApp.SetHandler()

	return &Server{
		BaseApp:    baseApp,
		baseConfig: baseConfig,
		mainConfig: mainConfig,
	}
}
