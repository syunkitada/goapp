package health

import (
	"os"

	"github.com/golang/glog"
	"github.com/urfave/cli"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/health/grpc_pb"
	"github.com/syunkitada/goapp/pkg/health/grpc_server"
	"github.com/syunkitada/goapp/pkg/lib/lib_grpc"
)

var (
	Conf = &config.Conf
)

func Main() error {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "goapp-health"
	app.Usage = "goapp-health"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)

	app.Action = func(c *cli.Context) error {
		config.Init(c)

		lis, grpcServer, err := lib_grpc.NewGrpcServer(&Conf.HealthGrpc)
		if err != nil {
			return err
		}

		healthServer := grpc_server.NewHealthServer()
		grpc_pb.RegisterHealthServer(grpcServer, healthServer)
		grpcServer.Serve(*lis)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		glog.Error(err)
		return err
	}

	return nil
}
