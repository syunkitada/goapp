package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_api"
	"github.com/syunkitada/goapp/pkg/resource/resolver"
	"github.com/syunkitada/goapp/pkg/resource/spec/genpkg"
)

type Server struct {
	base_app.BaseApp
	baseConf     *base_config.Config
	mainConf     *config.Config
	queryHandler *genpkg.QueryHandler
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	dbApi := db_api.New(baseConf, mainConf)
	baseApp := base_app.New(baseConf, &mainConf.Resource.App, dbApi)
	resolver := resolver.New(baseConf, mainConf)
	queryHandler := genpkg.NewQueryHandler(baseConf, &mainConf.Resource.App, dbApi, resolver)

	srv := &Server{
		BaseApp:      baseApp,
		baseConf:     baseConf,
		mainConf:     mainConf,
		queryHandler: queryHandler,
	}
	handler := srv.NewHandler()
	srv.SetHandler(handler)
	srv.SetDriver(srv)
	return srv
}

func (srv *Server) NewHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		var err error
		tctx, service, rawReq, req, rep, startTime, err := srv.Start(r)
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

		statusCode := 0
		var repBytes []byte
		for _, endpoint := range service.Endpoints {
			if endpoint == "" {
				fmt.Println("DEBUG service")
				if err = srv.queryHandler.Exec(tctx, req, rep); err != nil {
					continue
				}
				repBytes, err = json.Marshal(&rep)
				break
			}

			if repBytes, statusCode, err = srv.Proxy(tctx, endpoint, rawReq); err != nil {
				fmt.Println("DEBUG Failed err", err)
				continue
			} else {
				fmt.Println("DEBUG success")
				break
			}
		}

		if statusCode != 0 {
			w.WriteHeader(statusCode)
		} else {
			if err == nil {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		w.Write(repBytes)
	})

	return handler
}
