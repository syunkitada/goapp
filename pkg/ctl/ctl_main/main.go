package ctl_main

import (
	"os"

	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/urfave/cli"
)

var (
	Conf = &config.Conf
)

func Main() error {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "goapp-ctl"
	app.Usage = "help"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Action = func(c *cli.Context) error {
		config.Init(c)
		glog.Info("Hoge")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}
