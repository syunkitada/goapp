package server

import (
	"encoding/json"
	"net/http"

	"github.com/syunkitada/goapp/pkg/authproxy/autogen"
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/resolver"
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type Server struct {
	base_app.BaseApp
	baseConf     *base_config.Config
	mainConf     *config.Config
	queryHandler *autogen.QueryHandler
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	baseApp := base_app.New(baseConf, &mainConf.Authproxy.App)
	resolver := resolver.New(baseConf, mainConf)
	queryHandler := autogen.NewQueryHandler(baseConf, mainConf, resolver)

	srv := &Server{
		BaseApp:      baseApp,
		baseConf:     baseConf,
		mainConf:     mainConf,
		queryHandler: queryHandler,
	}
	handler := srv.NewHandler()
	srv.SetHandler(handler)
	return srv
}

func (srv *Server) NewHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		var err error
		tctx, service, req, rep, startTime, err := srv.Start(r)
		defer func() { srv.End(tctx, startTime, err) }()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			var bytes []byte
			bytes, err = json.Marshal(&rep)
			if err != nil {
				logger.Error(tctx, err, "Failed json.Marshal")
				return
			}
			w.Write(bytes)
			return
		}

		for _, endpoint := range service.Endpoints {
			if endpoint == "self" {
				err = srv.queryHandler.Exec(tctx, req, rep)
				break
			}

			// TODO Routing
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		repBytes, err := json.Marshal(&rep)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(repBytes)
	})

	return handler
}
