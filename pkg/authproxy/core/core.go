package core

import (
	"net/http"
	"time"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/health/grpc_client"
)

var (
	Conf = &config.Conf
)

type Authproxy struct {
	Listen       string
	AllowedHosts []string
	CertFile     string
	KeyFile      string
	HealthClient *grpc_client.HealthClient
}

func NewAuthproxy() *Authproxy {
	authproxy := &Authproxy{
		Listen:       Conf.Authproxy.Listen,
		AllowedHosts: Conf.Authproxy.AllowedHosts,
		CertFile:     Conf.Authproxy.CertFile,
		KeyFile:      Conf.Authproxy.KeyFile,
		HealthClient: grpc_client.NewHealthClient(),
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

	err := s.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		glog.Fatal("ListenAndServe: ", err)
	}
}
