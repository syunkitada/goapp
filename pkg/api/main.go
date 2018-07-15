package api

import (
	"os"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/urfave/cli"

	"github.com/golang/glog"
	"net/http"
	"time"

	// "github.com/syunkitada/goapp/pkg/api/services"
	"github.com/syunkitada/goapp/testdata"
)

var (
	Conf = &config.Conf
)

func Main() error {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "sample-simple-app"
	app.Usage = "sample-simple-app"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Action = func(c *cli.Context) error {
		config.Init(c)
		Serv()
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}

func Serv() {
	certPath := testdata.Path("tls-assets/server.pem")
	keyPath := testdata.Path("tls-assets/server.key")
	// api := NewAPI()

	handler := NewHandler()
	// handler.GET("/ping", debug.Ping)

	s := &http.Server{
		Addr:           ":8000",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// s.Handle("/hello", AddMiddleware(api.Hello,
	// 	api.SecureHeaders))
	// s.Handle("/tokens", api.Tokens)

	// s.Handle("/users", AddMiddleware(api.Users,
	// 	api.Authenticate,
	// 	api.Authorize(services.Permission("user_modify")),
	// 	api.SecureHeaders))

	err := s.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		glog.Fatal("ListenAndServe: ", err)
	}
}

// AddMiddleware adds middleware to a Handler
func AddMiddleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}
