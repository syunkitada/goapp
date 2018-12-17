package base

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type BaseAppDriver interface {
	RegisterGrpcServer(*grpc.Server) error
	MainTask(string) error
}

type BaseApp struct {
	Host               string
	Name               string
	conf               *config.Config
	appConf            *config.AppConfig
	driver             BaseAppDriver
	Monitors           int
	grpcServer         *grpc.Server
	loopInterval       time.Duration
	isGracefulShutdown bool
	shutdownTimeout    time.Duration
}

func NewBaseApp(conf *config.Config, appConf *config.AppConfig) BaseApp {
	return BaseApp{
		Host:               conf.Default.Host,
		Name:               appConf.Name,
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
	var traceId string
	var startTime time.Time
	var err error
	logger.Info(app.Host, app.Name, "Start mainLoop")
	for {
		traceId = logger.NewTraceId()
		startTime = logger.StartTaskTrace(traceId, app.Host, app.Name)
		err = app.driver.MainTask(traceId)
		logger.EndTaskTrace(traceId, app.Host, app.Name, startTime, err)

		if app.isGracefulShutdown {
			logger.Info(app.Host, app.Name, "Completed graceful shutdown mainTask")
			logger.Info(app.Host, app.Name, "Starting grpcServer.GracefulStop")
			app.grpcServer.GracefulStop()
			logger.Info(app.Host, app.Name, "Completed grpcServer.GracefulStop")
			logger.Info(app.Host, app.Name, "Completed graceful shutdown")
			os.Exit(0)
		}
		time.Sleep(app.loopInterval)
	}

	logger.Info(app.Host, app.Name, "Completed mainLoop")
}

func (app *BaseApp) Serve() error {
	lis, err := net.Listen("tcp", app.appConf.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		app.conf.Path(app.appConf.CertFile),
		app.conf.Path(app.appConf.KeyFile),
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
		<-shutdown
		if err := app.gracefulShutdown(context.Background()); err != nil {
			logger.Errorf(app.Host, app.Name, "Failed gracefulShutdown: %v", err)
		}
	}()

	logger.Infof(app.Host, app.Name, "Serve: %v", app.appConf.Listen)
	if err := app.grpcServer.Serve(lis); err != nil {
		logger.Errorf(app.Host, app.Name, "Failed Serve: %v", err)
		return err
	}

	logger.Infof(app.Host, app.Name, "Completed Serve: %v", app.appConf.Listen)
	return nil
}

func (app *BaseApp) gracefulShutdown(ctx context.Context) error {
	logger.Info(app.Host, app.Name, "Starting graceful shutdown")
	ctx, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	defer cancel()

	app.isGracefulShutdown = true

	select {
	case <-ctx.Done():
		logger.Errorf(app.Host, app.Name, "Failed graceful shutdown: %v", ctx.Err())
		os.Exit(1)
	}

	return nil
}
