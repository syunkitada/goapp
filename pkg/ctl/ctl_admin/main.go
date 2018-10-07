package ctl_admin

import (
	"os"

	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/urfave/cli"

	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
)

var (
	Conf = &config.Conf
)

type AdminCtl struct {
	Conf     *config.Config
	ModelApi *model_api.ModelApi
}

func NewAdminCtl(conf *config.Config) *AdminCtl {
	adminCtl := AdminCtl{
		Conf:     conf,
		ModelApi: model_api.NewModelApi(conf),
	}

	return &adminCtl
}

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
				adminCtl := NewAdminCtl(Conf)
				if err := adminCtl.MigrateDatabase(); err != nil {
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
