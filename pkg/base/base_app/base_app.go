package base_app

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/net/context"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type BaseAppDriver interface {
	MainTask(*logger.TraceContext, *gorm.DB) error
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
	handler            http.Handler
	shutdownTimeout    time.Duration
	rootClient         *base_client.Client
	dbApi              base_db_api.IApi
	serviceMap         map[string]base_model.ServiceRouter
}

func New(conf *base_config.Config, appConf *base_config.AppConfig, dbApi base_db_api.IApi) BaseApp {
	if appConf.LoopInterval == 0 {
		appConf.LoopInterval = 5
	}
	if appConf.CertFile == "" {
		appConf.CertFile = "server.pem"
	}
	if appConf.KeyFile == "" {
		appConf.KeyFile = "server.key"
	}

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
	}
}

func (app *BaseApp) SetHandler(handler http.Handler) {
	app.handler = handler
}

func (app *BaseApp) SetDriver(driver BaseAppDriver) {
	app.driver = driver
}

func (app *BaseApp) SyncService(tctx *logger.TraceContext, db *gorm.DB) (err error) {
	queries := []base_client.Query{}
	var data *base_spec.GetServicesData
	if data, err = app.dbApi.GetServices(tctx, db, &base_spec.GetServices{}); err != nil {
		return
	}

	serviceMap := map[string]base_model.ServiceRouter{}
	for _, service := range data.Services {
		var queryMap map[string]base_model.QueryModel
		if err = json.Unmarshal([]byte(service.QueryMap), &queryMap); err != nil {
			return
		}
		serviceMap[service.Name] = base_model.ServiceRouter{
			Endpoints: strings.Split(service.Endpoints, ","),
			QueryMap:  queryMap,
		}

		if service.SyncRootCluster {
			queries = append(queries, base_client.Query{
				Name: "UpdateService",
				Data: base_spec.UpdateService{
					Name:            service.Name,
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

	var db *gorm.DB
	if db, err = app.dbApi.Open(tctx); err != nil {
		return
	}
	defer app.dbApi.Close(tctx, db)
	if err = app.SyncService(tctx, db); err != nil {
		return
	}

	if err = app.driver.MainTask(tctx, db); err != nil {
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
		Handler:        app.handler,
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

func (app *BaseApp) Proxy(tctx *logger.TraceContext, endpoint string, rawReq []byte) (repBytes []byte, statusCode int, err error) {
	var httpResp *http.Response
	reqBuffer := bytes.NewBuffer(rawReq)
	var httpReq *http.Request
	if httpReq, err = http.NewRequest("POST", endpoint+"/q", reqBuffer); err != nil {
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Transport: tr,
	}

	httpResp, err = httpClient.Do(httpReq)
	if err != nil {
		return
	}

	defer func() {
		if tmpErr := httpResp.Body.Close(); tmpErr != nil {
			logger.Errorf(tctx, err, "Failed httpResp.Body.Close()")
		}
	}()
	statusCode = httpResp.StatusCode
	repBytes, err = ioutil.ReadAll(httpResp.Body)
	return
}
