package server

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type Server struct {
	base_app.BaseApp
	baseConf *base_config.Config
	mainConf *config.Config
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	baseApp := base_app.New(baseConf, &mainConf.Authproxy.App)

	srv := &Server{
		BaseApp:  baseApp,
		baseConf: baseConf,
		mainConf: mainConf,
	}
	handler := srv.NewHandler()
	srv.SetHandler(handler)
	return srv
}

func (srv *Server) NewHandler() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		var err error
		tctx, req, rep, startTime, err := srv.Start(r)
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

		fmt.Println(req)

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	return handler
}