package lib_grpc

import (
	"errors"
	"net"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
)

var (
	Conf = &config.Conf
)

func NewGrpcServer(conf *config.GrpcConfig) (*net.Listener, *grpc.Server, error) {
	lis, err := net.Listen("tcp", conf.Listen)
	if err != nil {
		return nil, nil, err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(Conf.Path(conf.CertFile), Conf.Path(conf.KeyFile))
	if err != nil {
		return nil, nil, err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	return &lis, grpcServer, nil
}

func NewClientConnection(conf *config.GrpcConfig) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	caFile := Conf.Path(conf.CaFile)
	targets := conf.Targets
	for _, target := range targets {
		creds, credsErr := credentials.NewClientTLSFromFile(caFile, conf.ServerHostOverride)
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
