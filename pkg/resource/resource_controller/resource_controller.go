package resource_controller

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
	"github.com/syunkitada/goapp/pkg/resource/resource_controller/resource_controller_grpc_pb"
)

type ResourceControllerServer struct {
	Conf               *config.Config
	GrpcServer         *grpc.Server
	ShutdownTimeout    time.Duration
	loopInterval       time.Duration
	isGracefulShutdown bool
}

func NewResourceControllerServer(conf *config.Config) *ResourceControllerServer {
	server := ResourceControllerServer{
		Conf:               conf,
		ShutdownTimeout:    time.Duration(10) * time.Second,
		loopInterval:       time.Duration(5) * time.Second,
		isGracefulShutdown: false,
	}
	return &server
}

func (server *ResourceControllerServer) StartMainLoop() error {
	go server.MainLoop()
	return nil
}

func (server *ResourceControllerServer) MainLoop() error {
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

func (server *ResourceControllerServer) Serv() error {
	grpcConfig := server.Conf.Resource.ControllerGrpc

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

	resource_controller_grpc_pb.RegisterResourceControllerServer(server.GrpcServer, server)
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

func (server *ResourceControllerServer) GracefulShutdown(ctx context.Context) error {
	glog.Info("Starting GracefulShutdown")
	ctx, cancel := context.WithTimeout(ctx, server.ShutdownTimeout)
	defer cancel()

	server.isGracefulShutdown = true

	select {
	case <-ctx.Done():
		glog.Warning(ctx.Err())
		os.Exit(1)
	}

	return nil
}
