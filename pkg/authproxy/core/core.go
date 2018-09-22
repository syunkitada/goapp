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
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/health/grpc_client"
)

var (
	Conf = &config.Conf
)

type Authproxy struct {
	Listen          string
	AllowedHosts    []string
	CertFile        string
	KeyFile         string
	HealthClient    *grpc_client.HealthClient
	GracefulTimeout time.Duration
	Auth            *auth.Auth
	Dashboard       *dashboard.Dashboard
}

func NewAuthproxy() *Authproxy {
	authproxy := &Authproxy{
		Listen:          Conf.Authproxy.Listen,
		AllowedHosts:    Conf.Authproxy.AllowedHosts,
		CertFile:        Conf.Authproxy.CertFile,
		KeyFile:         Conf.Authproxy.KeyFile,
		HealthClient:    grpc_client.NewHealthClient(),
		GracefulTimeout: time.Duration(Conf.Authproxy.GracefulTimeout) * time.Second,
		Auth:            auth.NewAuth(),
		Dashboard:       dashboard.NewDashboard(),
	}
	return authproxy
}

func (authproxy *Authproxy) Serv() {
	certPath := Conf.Path(authproxy.CertFile)
	keyPath := Conf.Path(authproxy.KeyFile)
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
		if err := s.ListenAndServeTLS(certPath, keyPath); err != nil && err != http.ErrServerClosed {
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
