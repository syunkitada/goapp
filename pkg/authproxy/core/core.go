package core

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model/authproxy_model_api"
	"github.com/syunkitada/goapp/pkg/authproxy/core/auth"
	"github.com/syunkitada/goapp/pkg/authproxy/core/dashboard"
	"github.com/syunkitada/goapp/pkg/authproxy/core/monitor"
	"github.com/syunkitada/goapp/pkg/config"
)

type Authproxy struct {
	host              string
	name              string
	conf              *config.Config
	Listen            string
	AllowedHosts      []string
	CertFilePath      string
	KeyFilePath       string
	GracefulTimeout   time.Duration
	AuthproxyModelApi *authproxy_model_api.AuthproxyModelApi
	Token             *auth.Token
	Auth              *auth.Auth
	Dashboard         *dashboard.Dashboard
	Monitor           *monitor.Monitor
}

func NewAuthproxy(conf *config.Config) *Authproxy {
	authproxyModelApi := authproxy_model_api.NewAuthproxyModelApi(conf)
	token := auth.NewToken(conf, authproxyModelApi)

	authproxy := &Authproxy{
		host:              conf.Default.Host,
		name:              "authproxy",
		conf:              conf,
		Listen:            conf.Authproxy.HttpServer.Listen,
		AllowedHosts:      conf.Authproxy.HttpServer.AllowedHosts,
		CertFilePath:      conf.Path(conf.Authproxy.HttpServer.CertFile),
		KeyFilePath:       conf.Path(conf.Authproxy.HttpServer.KeyFile),
		GracefulTimeout:   time.Duration(conf.Authproxy.HttpServer.GracefulTimeout) * time.Second,
		AuthproxyModelApi: authproxyModelApi,
		Token:             token,
		Auth:              auth.NewAuth(conf, authproxyModelApi, token),
		Dashboard:         dashboard.NewDashboard(conf, authproxyModelApi, token),
		Monitor:           monitor.NewMonitor(conf),
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
