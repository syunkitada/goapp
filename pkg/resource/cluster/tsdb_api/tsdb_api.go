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

func (api *Api) ReportResource(tctx *logger.TraceContext, input *api_spec.ReportResource) (err error) {
	return api.driver.Report(tctx, input)
}
