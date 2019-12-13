package tsdb_api

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/tsdb_api/drivers"
	"github.com/syunkitada/goapp/pkg/resource/config"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Api struct {
	tsdbConf    config.TimeSeriesDatabaseConfig
	baseConf    *base_config.Config
	clusterConf *config.ResourceClusterConfig
	driver      drivers.TsdbDriver
}

func New(baseConf *base_config.Config, clusterConf *config.ResourceClusterConfig) *Api {
	driver := drivers.Load(clusterConf)

	api := Api{
		tsdbConf:    clusterConf.TimeSeriesDatabase,
		baseConf:    baseConf,
		clusterConf: clusterConf,
		driver:      driver,
	}

	return &api
}

func (api *Api) ReportNode(tctx *logger.TraceContext, input *api_spec.ReportNode) error {
	return api.driver.Report(tctx, input)
}

func (api *Api) GetNode(tctx *logger.TraceContext, input *api_spec.GetNode) ([]api_spec.MetricsGroup, error) {
	return api.driver.GetNode(tctx, input)
}

func (api *Api) GetLogParams(tctx *logger.TraceContext, input *api_spec.GetLogParams) (*api_spec.GetLogParamsData, error) {
	return api.driver.GetLogParams(tctx, input)
}
