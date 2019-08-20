package server

import (
	"encoding/json"
	"net/http"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type IResolver interface {
	IssueToken(input *spec.IssueToken) (*spec.IssueTokenData, error)
}

type Server struct {
	base_app.BaseApp
	baseConf *base_config.Config
	mainConf *config.Config
	resolver IResolver
}

func New(baseConf *base_config.Config, mainConf *config.Config, resolver IResolver) *Server {
	baseApp := base_app.New(baseConf, &mainConf.Authproxy.App)

	srv := &Server{
		BaseApp:  baseApp,
		baseConf: baseConf,
		mainConf: mainConf,
		resolver: resolver,
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
				err = srv.Exec(req, rep)
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

func (srv *Server) Exec(req *base_model.Request, rep *base_model.Reply) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "IssueToken":
			var input spec.IssueToken
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, err := srv.resolver.IssueToken(&input)
			rep.Data["IssueToken"] = data
			return err
		}
	}
	return nil
}
