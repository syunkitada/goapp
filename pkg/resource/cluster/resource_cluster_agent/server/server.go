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

	clusterConf.Agent.Name = consts.KindResourceClusterAgent
	resolver := resolver.New(baseConf, &clusterConf)
	queryHandler := genpkg.NewQueryHandler(baseConf, &clusterConf.Agent.AppConfig, resolver)
	baseApp := base_app.New(baseConf, &clusterConf.Agent.AppConfig, nil, queryHandler)
	apiClient := resource_cluster_api.NewClient(&clusterConf.Agent.RootClient)

	computeConf := clusterConf.Agent.Compute
	computeExConf := config.ResourceComputeExConfig{
		ResourceComputeConfig: computeConf,
		ConfirmRetryInterval:  time.Duration(computeConf.ConfirmRetryInterval) * time.Second,
		VmNetnsGatewayStartIp: net.ParseIP(computeConf.VmNetnsGatewayStartIp),
		VmNetnsGatewayEndIp:   net.ParseIP(computeConf.VmNetnsGatewayEndIp),
		VmNetnsServiceIp:      net.ParseIP(computeConf.VmNetnsServiceIp),
		VmNetnsStartIp:        net.ParseIP(computeConf.VmNetnsStartIp),
		VmNetnsEndIp:          net.ParseIP(computeConf.VmNetnsEndIp),
		VmsDir:                filepath.Join(computeConf.VarDir, "vms"),
		ImagesDir:             filepath.Join(computeConf.VarDir, "images"),
		UserdataTmpl:          filepath.Join(computeConf.ConfigDir, "user-data.tmpl"),
		VmServiceTmpl:         filepath.Join(computeConf.ConfigDir, "vm-service.tmpl"),
		VmServiceShTmpl:       filepath.Join(computeConf.ConfigDir, "vm-service.sh.tmpl"),
		SystemdDir:            "/etc/systemd/system",
	}

	os_utils.MustMkdir(computeExConf.VarDir, 0755)
	os_utils.MustMkdir(computeExConf.VmsDir, 0755)
	os_utils.MustMkdir(computeExConf.ImagesDir, 0755)

	computeDriver := compute_drivers.Load(&computeExConf)

	// init metric readers
	metricReaderMap := map[string]readers.MetricReader{}
	if clusterConf.Agent.Metric.System.Enable {
		metricReaderMap["system"] = system_metric_reader.New(&clusterConf.Agent.Metric.System)
	}

	srv := &Server{
		BaseApp:       baseApp,
		baseConf:      baseConf,
		clusterConf:   &clusterConf,
		queryHandler:  queryHandler,
		apiClient:     apiClient,
		computeConf:   computeExConf,
		computeDriver: computeDriver,

		reportCount:           0,
		reportSpan:            10,
		metricReaderMap:       metricReaderMap,
		logMap:                clusterConf.Agent.LogMap,
		logReaderMap:          map[string]*log_reader.LogReader{},
		logReaderRefreshSpan:  10,
		logReaderRefreshCount: 0,
	}
	srv.SetDriver(srv)
	return srv
}
