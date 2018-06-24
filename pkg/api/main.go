package api

import (
	"os"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/urfave/cli"

	"github.com/golang/glog"
	"net/http"

	"github.com/syunkitada/goapp/pkg/api/services"
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
	api := NewAPI(certPath, keyPath)

	http.Handle("/hello", api.Hello)
	http.Handle("/tokens", api.Tokens)

	http.Handle("/users", AddMiddleware(api.Users,
		api.Authenticate,
		api.Authorize(services.Permission("user_modify"))))

	err := http.ListenAndServeTLS(":8000", certPath, keyPath, nil)
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
