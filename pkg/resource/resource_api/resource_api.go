package resource_api

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type ResourceApiServer struct {
	Conf               *config.Config
	GrpcServer         *grpc.Server
	ShutdownTimeout    time.Duration
	loopInterval       time.Duration
	isGracefulShutdown bool
	resourceModelApi   *resource_model_api.ResourceModelApi
	resourceApiClient  *resource_api_client.ResourceApiClient
}

func NewResourceApiServer(conf *config.Config) *ResourceApiServer {
	server := ResourceApiServer{
		Conf:               conf,
		ShutdownTimeout:    time.Duration(10) * time.Second,
		loopInterval:       time.Duration(5) * time.Second,
		isGracefulShutdown: false,
		resourceModelApi:   resource_model_api.NewResourceModelApi(conf),
		resourceApiClient:  resource_api_client.NewResourceApiClient(conf),
	}
	return &server
}

func (server *ResourceApiServer) StartMainLoop() error {
	go server.MainLoop()
	return nil
}

func (server *ResourceApiServer) MainLoop() error {
	glog.Info("Starting MainLoop")
	for {
		if err := server.MainTask(); err != nil {
			glog.Warning(err)
		}

		if server.isGracefulShutdown {
			glog.Info("End MainLoop on GracefulShutdown")
			glog.Info("Start GrpcServer.GracefulStop")
			server.GrpcServer.GracefulStop()
			glog.Info("Success GrpcServer.GracefulStop")
			glog.Info("Success GracefulShutdown")
			os.Exit(0)
		}
		glog.Infof("Success MainTask, and sleep %v", server.loopInterval)
		time.Sleep(server.loopInterval)
	}
	return nil
}

func (server *ResourceApiServer) Serv() error {
	grpcConfig := server.Conf.Resource.ApiGrpc

	lis, err := net.Listen("tcp", grpcConfig.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		server.Conf.Path(grpcConfig.CertFile),
		server.Conf.Path(grpcConfig.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	server.GrpcServer = grpc.NewServer(opts...)

	resource_api_grpc_pb.RegisterResourceApiServer(server.GrpcServer, server)
	glog.Infof("Serve: %v", grpcConfig.Listen)

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM)
		signal.Notify(shutdown, syscall.SIGINT)
		<-shutdown
		if err := server.GracefulShutdown(context.Background()); err != nil {
			glog.Errorf("App Shutdown: %v\n", err)
		}
	}()

	if err := server.GrpcServer.Serve(lis); err != nil {
		glog.Infof("App Serv Failed: %v\n", err)
	}

	glog.Infof("End Serv")
	return nil
}

func (server *ResourceApiServer) GracefulShutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, server.ShutdownTimeout)
	defer cancel()

	go func() {
		glog.Info("Start GracefulStop")
		server.GrpcServer.GracefulStop()
		glog.Info("Success GracefulStop")
		os.Exit(0)
	}()

	select {
	case <-ctx.Done():
		glog.Warning(ctx.Err())
		os.Exit(1)
	}

	return nil
}

func (srv *ResourceApiServer) Status(ctx context.Context, statusRequest *resource_api_grpc_pb.StatusRequest) (*resource_api_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

func (srv *ResourceApiServer) GetNode(ctx context.Context, req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	glog.Info("GetNode")
	var err error
	var rep *resource_api_grpc_pb.GetNodeReply
	if rep, err = srv.resourceModelApi.GetNode(req); err != nil {
		glog.Error(err)
	}
	return rep, err
}

func (srv *ResourceApiServer) UpdateNode(ctx context.Context, req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	glog.Info("UpdateNode")
	if err = srv.resourceModelApi.UpdateNode(req); err != nil {
		glog.Error(err)
	}
	return &resource_api_grpc_pb.UpdateNodeReply{}, err
}
