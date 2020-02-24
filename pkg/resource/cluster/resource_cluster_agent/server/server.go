package server

import (
	"net"
	"path/filepath"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/db_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/readers"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/readers/log_reader"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/readers/system_metric_reader"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/resolver"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/spec/genpkg"
	resource_cluster_api "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/consts"
)

type Server struct {
	base_app.BaseApp
	baseConf     *base_config.Config
	clusterConf  *config.ResourceClusterConfig
	queryHandler *genpkg.QueryHandler
	apiClient    *resource_cluster_api.Client

	computeConf   config.ResourceComputeExConfig
	computeDriver compute_drivers.ComputeDriver

	reportCount int
	reportSpan  int

	dbApi *db_api.Api

	metricReaderMap map[string]readers.MetricReader

	logMap                map[string]config.ResourceLogConfig
	logReaderMap          map[string]*log_reader.LogReader
	logReaderRefreshSpan  int
	logReaderRefreshCount int
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	clusterConf, ok := mainConf.Resource.ClusterMap[mainConf.Resource.ClusterName]
	if !ok {
		logger.StdoutFatalf("cluster config is not found: cluster=%s", mainConf.Resource.ClusterName)
	}
	tctx := logger.NewTraceContext(baseConf.Host, "init")

	clusterConf.Agent.Name = consts.KindResourceClusterAgent
	resolver := resolver.New(baseConf, &clusterConf)
	queryHandler := genpkg.NewQueryHandler(baseConf, &clusterConf.Agent.AppConfig, resolver)
	dbApi := db_api.New(baseConf, &clusterConf)
	baseApp := base_app.New(baseConf, &clusterConf.Agent.AppConfig, dbApi, queryHandler)
	apiClient := resource_cluster_api.NewClient(&clusterConf.Agent.RootClient)

	srv := &Server{
		BaseApp:      baseApp,
		baseConf:     baseConf,
		clusterConf:  &clusterConf,
		queryHandler: queryHandler,
		apiClient:    apiClient,
		dbApi:        dbApi,
	}
	srv.initDb(tctx)
	srv.initReader()
	srv.initComputeDriver()
	resolver.SetComputeDriver(srv.computeDriver)

	srv.SetDriver(srv)
	return srv
}

func (srv *Server) initDb(tctx *logger.TraceContext) {
	var err error
	if err = srv.dbApi.Bootstrap(tctx, false); err != nil {
		logger.Fatalf(tctx, "Failed bootstrap: %s", err.Error())
	}
	if err = srv.dbApi.BootstrapResource(tctx, false); err != nil {
		logger.Fatalf(tctx, "Failed bootstrap: %s", err.Error())
	}
	srv.dbApi.MustOpen()
	if err = srv.SyncService(tctx, false); err != nil {
		logger.Fatalf(tctx, "Failed SyncService: %s", err.Error())
	}
}

func (srv *Server) initReader() {
	metricReaderMap := map[string]readers.MetricReader{}
	if srv.clusterConf.Agent.Metric.System.Enable {
		metricReaderMap["system"] = system_metric_reader.New(&srv.clusterConf.Agent.Metric.System)
	}

	srv.reportCount = 0
	srv.reportSpan = 10
	srv.metricReaderMap = metricReaderMap
	srv.logMap = srv.clusterConf.Agent.LogMap
	srv.logReaderMap = map[string]*log_reader.LogReader{}
	srv.logReaderRefreshSpan = 10
	srv.logReaderRefreshCount = 0
}

func (srv *Server) initComputeDriver() {
	computeConf := srv.clusterConf.Agent.Compute
	vmsDir := filepath.Join(computeConf.VarDir, "vms")
	if computeConf.VmsDir != "" {
		vmsDir = computeConf.VmsDir
	}
	imagesDir := filepath.Join(computeConf.VarDir, "images")
	if computeConf.ImagesDir != "" {
		vmsDir = computeConf.ImagesDir
	}

	computeExConf := config.ResourceComputeExConfig{
		ResourceComputeConfig: computeConf,
		ConfirmRetryInterval:  time.Duration(computeConf.ConfirmRetryInterval) * time.Second,
		VmNetnsGatewayStartIp: net.ParseIP(computeConf.VmNetnsGatewayStartIp),
		VmNetnsGatewayEndIp:   net.ParseIP(computeConf.VmNetnsGatewayEndIp),
		VmNetnsServiceIp:      net.ParseIP(computeConf.VmNetnsServiceIp),
		VmNetnsStartIp:        net.ParseIP(computeConf.VmNetnsStartIp),
		VmNetnsEndIp:          net.ParseIP(computeConf.VmNetnsEndIp),
		VmsDir:                vmsDir,
		ImagesDir:             imagesDir,
		UserdataTmpl:          filepath.Join(computeConf.ConfigDir, "user-data.tmpl"),
		VmServiceTmpl:         filepath.Join(computeConf.ConfigDir, "vm-service.tmpl"),
		VmServiceShTmpl:       filepath.Join(computeConf.ConfigDir, "vm-service.sh.tmpl"),
		SystemdDir:            "/etc/systemd/system",
	}

	os_utils.MustMkdir(computeExConf.VarDir, 0755)
	os_utils.MustMkdir(computeExConf.VmsDir, 0755)
	os_utils.MustMkdir(computeExConf.ImagesDir, 0755)

	srv.computeConf = computeExConf
	srv.computeDriver = compute_drivers.Load(&computeExConf)
}
