package model_ctl

import (
	"os"

	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/model/model_manager"
	"github.com/urfave/cli"
)

var (
	Conf = &config.Conf
)

func Main() error {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "goapp-db-migrate"
	app.Usage = "help"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Action = func(c *cli.Context) error {
		config.Init(c)
		if err := model_manager.MigrateDatabase(); err != nil {
			glog.Error(err)
			os.Exit(1)
		}

		glog.Info("Success DB Migrate")
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		return err
	}

	return nil
}
