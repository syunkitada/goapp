package base

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
)

type BaseAppDriver interface {
	RegisterGrpcServer(*grpc.Server) error
	MainTask() error
}

type BaseApp struct {
	conf               *config.Config
	appConf            *config.AppConfig
	driver             BaseAppDriver
	grpcServer         *grpc.Server
	loopInterval       time.Duration
	isGracefulShutdown bool
	shutdownTimeout    time.Duration
}

func NewBaseApp(conf *config.Config, appConf *config.AppConfig) BaseApp {
	return BaseApp{
		conf:               conf,
		appConf:            appConf,
		loopInterval:       time.Duration(appConf.LoopInterval) * time.Second,
		isGracefulShutdown: false,
		shutdownTimeout:    time.Duration(appConf.ShutdownTimeout) * time.Second,
	}
}

func (app *BaseApp) RegisterDriver(driver BaseAppDriver) {
	app.driver = driver
}

func (app *BaseApp) StartMainLoop() error {
	go app.mainLoop()
	return nil
}

func (app *BaseApp) mainLoop() {
	glog.Info("Starting mainLoop")
	for {
		if err := app.driver.MainTask(); err != nil {
			glog.Warning(err)
		}

		if app.isGracefulShutdown {
			glog.Info("Completed graceful shutdown mainTask")
			glog.Info("Starting grpcServer.GracefulStop")
			app.grpcServer.GracefulStop()
			glog.Info("Completed grpcServer.GracefulStop")
			glog.Info("Completed graceful shutdown")
			os.Exit(0)
		}
		glog.Infof("Completed mainTask, and sleep %v", app.loopInterval)
		time.Sleep(app.loopInterval)
	}

	glog.Info("Completed mainLoop")
}

func (app *BaseApp) Serve() error {
	lis, err := net.Listen("tcp", app.appConf.Grpc.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		app.conf.Path(app.appConf.Grpc.CertFile),
		app.conf.Path(app.appConf.Grpc.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	app.grpcServer = grpc.NewServer(opts...)
	if err := app.driver.RegisterGrpcServer(app.grpcServer); err != nil {
		return err
	}

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM)
		signal.Notify(shutdown, syscall.SIGINT)
		<-shutdown
		if err := app.gracefulShutdown(context.Background()); err != nil {
			glog.Errorf("Failed gracefulShutdown: %v\n", err)
		}
	}()

	glog.Infof("Serve: %v", app.appConf.Grpc.Listen)
	if err := app.grpcServer.Serve(lis); err != nil {
		glog.Infof("Failed Serve: %v\n", err)
	}

	glog.Infof("Completed Serve")
	return nil
}

func (app *BaseApp) gracefulShutdown(ctx context.Context) error {
	glog.Info("Starting graceful shutdown")
	ctx, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	defer cancel()

	app.isGracefulShutdown = true

	select {
	case <-ctx.Done():
		glog.Warning(ctx.Err())
		os.Exit(1)
	}

	return nil
}
