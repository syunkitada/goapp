package ctl_admin

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
	app.Name = "goapp-admin-ctl"
	app.Usage = "help"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Commands = []cli.Command{
		{
			Name:  "db-migrate",
			Usage: "db-migrate help",
			Action: func(c *cli.Context) error {
				config.Init(c)
				if err := MigrateDatabase(); err != nil {
					return err
				}

				glog.Info("Success DB Migrate")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}
