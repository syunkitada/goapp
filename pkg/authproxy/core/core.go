package core

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/core/auth"
	"github.com/syunkitada/goapp/pkg/authproxy/core/dashboard"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/health/grpc_client"
)

type Authproxy struct {
	Conf            *config.Config
	Listen          string
	AllowedHosts    []string
	CertFilePath    string
	KeyFilePath     string
	HealthClient    *grpc_client.HealthClient
	GracefulTimeout time.Duration
	ModelApi        *model_api.ModelApi
	Token           *auth.Token
	Auth            *auth.Auth
	Dashboard       *dashboard.Dashboard
}

func NewAuthproxy(conf *config.Config) *Authproxy {
	modelApi := model_api.NewModelApi(conf)
	token := auth.NewToken(conf, modelApi)

	authproxy := &Authproxy{
		Conf:            conf,
		Listen:          conf.Authproxy.Listen,
		AllowedHosts:    conf.Authproxy.AllowedHosts,
		CertFilePath:    conf.Path(conf.Authproxy.CertFile),
		KeyFilePath:     conf.Path(conf.Authproxy.KeyFile),
		HealthClient:    grpc_client.NewHealthClient(),
		GracefulTimeout: time.Duration(conf.Authproxy.GracefulTimeout) * time.Second,
		Token:           token,
		Auth:            auth.NewAuth(conf, modelApi, token),
		Dashboard:       dashboard.NewDashboard(conf, modelApi, token),
	}

	if conf.Default.TestMode {
		conf.Authproxy.TestHandler = authproxy.NewHandler()
	}

	return authproxy
}

func (authproxy *Authproxy) Serv() {
	handler := authproxy.NewHandler()

	s := &http.Server{
		Addr:           authproxy.Listen,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		// service connections
		if err := s.ListenAndServeTLS(authproxy.CertFilePath, authproxy.KeyFilePath); err != nil && err != http.ErrServerClosed {
			glog.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	glog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), authproxy.GracefulTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		glog.Fatal("Server Shutdown:", err)
	}
	glog.Info("Server exiting")
}
