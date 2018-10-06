package resource_server

import (
	"os"

	"github.com/golang/glog"
	"github.com/urfave/cli"

	"github.com/syunkitada/goapp/pkg/config"
)

var (
	Conf = &config.Conf
)

func Main() error {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "goapp-resource-server"
	app.Usage = "goapp-resource-server"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Action = func(c *cli.Context) error {
		config.Init(c)
		resourceServer := NewResourceServer(Conf)
		return resourceServer.Serv()
	}

	if err := app.Run(os.Args); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}
