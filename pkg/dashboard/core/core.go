package core

// import (
// 	"context"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"time"
// 
// 	"github.com/golang/glog"
// 
// 	"github.com/syunkitada/goapp/pkg/config"
// )
// 
// var (
// 	Conf = &config.Conf
// )
// 
// type Dashboard struct {
// 	Listen          string
// 	AllowedHosts    []string
// 	CertFile        string
// 	KeyFile         string
// 	GracefulTimeout time.Duration
// 	BuildDir        string
// }
// 
// func NewDashboard() *Dashboard {
// 	// TODO
// 	dashboard := &Dashboard{}
// 	return dashboard
// }
// 
// func (dashboard *Dashboard) Serv() {
// 	certPath := Conf.Path(dashboard.CertFile)
// 	keyPath := Conf.Path(dashboard.KeyFile)
// 	handler := dashboard.NewHandler()
// 
// 	s := &http.Server{
// 		Addr:           dashboard.Listen,
// 		Handler:        handler,
// 		ReadTimeout:    10 * time.Second,
// 		WriteTimeout:   10 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 
// 	go func() {
// 		// service connections
// 		if err := s.ListenAndServeTLS(certPath, keyPath); err != nil && err != http.ErrServerClosed {
// 			glog.Fatalf("listen: %s\n", err)
// 		}
// 	}()
// 
// 	// Wait for interrupt signal to gracefully shutdown the server with
// 	// a timeout of 5 seconds.
// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit
// 	glog.Info("Shutdown Server ...")
// 
// 	ctx, cancel := context.WithTimeout(context.Background(), dashboard.GracefulTimeout)
// 	defer cancel()
// 	if err := s.Shutdown(ctx); err != nil {
// 		glog.Fatal("Server Shutdown:", err)
// 	}
// 	glog.Info("Server exiting")
// 
// }
