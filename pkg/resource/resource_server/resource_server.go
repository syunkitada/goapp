package resource_server

import (
	"net"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/grpc_pb"
)

type ResourceServer struct {
	Conf *config.Config
}

func NewResourceServer(conf *config.Config) *ResourceServer {
	server := ResourceServer{
		Conf: conf,
	}
	return &server
}

func (resourceServer *ResourceServer) Status(ctx context.Context, statusRequest *grpc_pb.StatusRequest) (*grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}

func (resourceServer *ResourceServer) Serv() error {
	grpcConfig := resourceServer.Conf.Resource.Grpc

	lis, err := net.Listen("tcp", grpcConfig.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		resourceServer.Conf.Path(grpcConfig.CertFile),
		resourceServer.Conf.Path(grpcConfig.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	grpc_pb.RegisterHealthServer(grpcServer, resourceServer)
	grpcServer.Serve(lis)

	return nil
}
