package base

import (
	"errors"
	"time"

	"golang.org/x/net/context"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
)

type BaseClient struct {
	appConf            *config.AppConfig
	caFilePath         string
	serverHostOverride string
	timeout            time.Duration
}

func NewBaseClient(conf *config.Config, appConf *config.AppConfig) *BaseClient {
	return &BaseClient{
		appConf:            appConf,
		caFilePath:         conf.Path(appConf.CaFile),
		serverHostOverride: appConf.ServerHostOverride,
		timeout:            time.Duration(appConf.ClientTimeout) * time.Second,
	}
}

func (cli *BaseClient) NewClientConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	for _, target := range cli.appConf.Targets {
		creds, credsErr := credentials.NewClientTLSFromFile(cli.caFilePath, cli.serverHostOverride)
		if credsErr != nil {
			glog.Warning("Failed to create TLS credentials %v", credsErr)
			continue
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(target, opts...)
		if err != nil {
			glog.Warning("fail to dial: %v", err)
			continue
		}

		return conn, nil
	}

	return nil, errors.New("Failed NewGrpcConnection")
}

func (cli *BaseClient) GetContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), cli.timeout)
}
