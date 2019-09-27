package base_app

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type BaseAppDriver interface {
	MainTask(tctx *logger.TraceContext) (err error)
}

type BaseApp struct {
	host               string
	name               string
	conf               *base_config.Config
	appConf            *base_config.AppConfig
	driver             BaseAppDriver
	loopInterval       time.Duration
	isGracefulShutdown bool
	server             *http.Server
	shutdownTimeout    time.Duration
	rootClient         *base_client.Client
	dbApi              base_db_api.IApi
	serviceMap         map[string]spec_model.ServiceRouter
	queryHandler       IQueryHandler
}

type IQueryHandler interface {
	Exec(tctx *logger.TraceContext, userAuthority *base_spec.UserAuthority, httpReq *http.Request, rw http.ResponseWriter, req *base_model.Request, rep *base_model.Response) error
}

func New(conf *base_config.Config, appConf *base_config.AppConfig, dbApi base_db_api.IApi, queryHandler IQueryHandler) BaseApp {
	if appConf.LoopInterval == 0 {
		appConf.LoopInterval = 5
	}
	if appConf.CertFile == "" {
		appConf.CertFile = "server.pem"
	}
	if appConf.KeyFile == "" {
		appConf.KeyFile = "server.key"
	}

	dbApi.MustOpen()

	return BaseApp{
		host:               conf.Host,
		name:               appConf.Name,
		conf:               conf,
		appConf:            appConf,
		loopInterval:       time.Duration(appConf.LoopInterval) * time.Second,
		isGracefulShutdown: false,
		shutdownTimeout:    time.Duration(appConf.ShutdownTimeout) * time.Second,
		rootClient:         base_client.NewClient(&appConf.RootClient),
		dbApi:              dbApi,
		queryHandler:       queryHandler,
	}
}

func (app *BaseApp) Bootstrap() {
	app.dbApi.MustOpen()
}

func (app *BaseApp) MustClose() {
}

func (app *BaseApp) SetDriver(driver BaseAppDriver) {
	app.driver = driver
}

func (app *BaseApp) SyncService(tctx *logger.TraceContext) (err error) {
	queries := []base_client.Query{}
	var data *base_spec.GetServicesData
	if data, err = app.dbApi.GetServices(tctx, &base_spec.GetServices{}); err != nil {
		return
	}

	serviceMap := map[string]spec_model.ServiceRouter{}
	for _, service := range data.Services {
		var queryMap map[string]spec_model.QueryModel
		if err = json.Unmarshal([]byte(service.QueryMap), &queryMap); err != nil {
			return
		}
		serviceMap[service.Name] = spec_model.ServiceRouter{
			Token:     service.Token,
			Endpoints: strings.Split(service.Endpoints, ","),
			QueryMap:  queryMap,
		}

		if service.SyncRootCluster {
			var token string
			token, err = app.dbApi.IssueToken(service.Name)
			if err != nil {
				return
			}
			queries = append(queries, base_client.Query{
				Name: "UpdateService",
				Data: base_spec.UpdateService{
					Name:            service.Name,
					Token:           token,
					Scope:           service.Scope,
					Endpoints:       app.appConf.Endpoints,
					ProjectRoles:    strings.Split(service.ProjectRoles, ","),
					QueryMap:        queryMap,
					SyncRootCluster: false,
				},
			})
		}
	}
	app.serviceMap = serviceMap

	if len(queries) > 0 {
		// var data *base_spec.UpdateServiceData,
		if _, err = app.rootClient.UpdateServices(tctx, queries); err != nil {
			return
		}
	}
	return
}

func (app *BaseApp) StartMainLoop() {
	go app.mainLoop()
}

func (app *BaseApp) mainLoop() {
	var tctx *logger.TraceContext
	var startTime time.Time
	var err error
	logger.StdoutInfof("Start mainLoop")
	for {
		tctx = logger.NewTraceContext(app.host, app.name)
		app.mainTask(tctx)
		if app.isGracefulShutdown {
			logger.Info(tctx, "Completed graceful shutdown mainTask")
			logger.Info(tctx, "Starting graceful shutdown server")
			startTime = logger.StartTrace(tctx)

			ctx, cancel := context.WithTimeout(context.Background(), app.shutdownTimeout)
			defer cancel()
			if err = app.server.Shutdown(ctx); err != nil {
				logger.Fatalf(tctx, "Failed graceful shutdown: %v", err)
			}

			app.dbApi.MustClose()
			logger.Info(tctx, "Completed graceful shutdown")
			logger.EndTrace(tctx, startTime, nil, 0)
			os.Exit(0)
		}
		time.Sleep(app.loopInterval)
	}
}

func (app *BaseApp) mainTask(tctx *logger.TraceContext) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 0) }()

	if err = app.SyncService(tctx); err != nil {
		return
	}

	if err = app.driver.MainTask(tctx); err != nil {
		return
	}
}

func (app *BaseApp) NewTraceContext() *logger.TraceContext {
	return logger.NewTraceContext(app.host, app.name)
}

func (app *BaseApp) Serve() {
	var err error
	tctx := logger.NewTraceContext(app.host, app.name)
	startTime := logger.StartTrace(tctx)
	defer func() {
		logger.EndTrace(tctx, startTime, err, 0)
	}()

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM)
		<-shutdown
		if err = app.gracefulShutdown(context.Background()); err != nil {
			logger.Error(tctx, err, "Failed gracefulShutdown")
		}
	}()

	app.server = &http.Server{
		Addr:           app.appConf.Listen,
		Handler:        app.NewHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	logger.Infof(tctx, "Serve: %v", app.appConf.Listen)

	certFile := filepath.Join(app.conf.ConfigDir, app.appConf.CertFile)
	keyFile := filepath.Join(app.conf.ConfigDir, app.appConf.KeyFile)
	err = app.server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		logger.Fatalf(tctx, "Failed ListenAndServeTLS: certFile=%s, keyFile=%s, %v", certFile, keyFile, err)
	}

	logger.Infof(tctx, "Completed Serve: %v", app.appConf.Listen)
}

func (app *BaseApp) gracefulShutdown(ctx context.Context) error {
	var err error
	tctx := logger.NewTraceContext(app.host, app.name)
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
