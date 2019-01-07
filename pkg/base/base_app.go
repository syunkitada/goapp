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
	MainTask(*logger.TraceContext) error
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
	var tctx *logger.TraceContext
	var startTime time.Time
	var err error
	logger.StdoutInfof("Start mainLoop")
	for {
		tctx = logger.NewTraceContext(app.Host, app.Name)
		startTime = logger.StartTrace(tctx)
		err = app.driver.MainTask(tctx)
		logger.EndTrace(tctx, startTime, err, 0)

		if app.isGracefulShutdown {
			logger.Info(tctx, "Completed graceful shutdown mainTask")
			logger.Info(tctx, "Starting grpcServer.GracefulStop")
			startTime = logger.StartTrace(tctx)
			app.grpcServer.GracefulStop()
			logger.Info(tctx, "Completed grpcServer.GracefulStop")
			logger.Info(tctx, "Completed graceful shutdown")
			logger.EndTrace(tctx, startTime, nil, 0)
			os.Exit(0)
		}

		time.Sleep(app.loopInterval)
	}

	logger.StdoutInfof("Completed mainLoop")
}

func (app *BaseApp) Serve() error {
	var err error
	tctx := logger.NewTraceContext(app.Host, app.Name)
	startTime := logger.StartTrace(tctx)
	defer func() {
		logger.EndTrace(tctx, startTime, err, 0)
	}()

	var lis net.Listener
	lis, err = net.Listen("tcp", app.appConf.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	var creds credentials.TransportCredentials
	creds, err = credentials.NewServerTLSFromFile(
		app.conf.Path(app.appConf.CertFile),
		app.conf.Path(app.appConf.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	app.grpcServer = grpc.NewServer(opts...)
	if err = app.driver.RegisterGrpcServer(app.grpcServer); err != nil {
		return err
	}

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM)
		<-shutdown
		if err := app.gracefulShutdown(context.Background()); err != nil {
			logger.Error(tctx, err, "Failed gracefulShutdown")
		}
	}()

	logger.Infof(tctx, "Serve: %v", app.appConf.Listen)
	if err := app.grpcServer.Serve(lis); err != nil {
		logger.Error(tctx, err, "Failed Serve")
		return err
	}

	logger.Infof(tctx, "Completed Serve: %v", app.appConf.Listen)
	return nil
}

func (app *BaseApp) gracefulShutdown(ctx context.Context) error {
	var err error
	tctx := logger.NewTraceContext(app.Host, app.Name)
	startTime := logger.StartTrace(tctx)
	defer func() {
		logger.EndTrace(tctx, startTime, err, 0)
	}()

	ctx, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	defer cancel()

	app.isGracefulShutdown = true

	select {
	case <-ctx.Done():
		err = ctx.Err()
		os.Exit(1)
	}

	return nil
}
