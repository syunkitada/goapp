package tsdb_api

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/tsdb_api/drivers"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
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

func (api *Api) SetFilterEventRules(tctx *logger.TraceContext, eventRules []db_model.EventRule) {
	api.driver.SetFilterEventRules(tctx, eventRules)
}

func (api *Api) ReportNode(tctx *logger.TraceContext, input *api_spec.ReportNode) error {
	return api.driver.Report(tctx, input)
}

func (api *Api) GetNode(tctx *logger.TraceContext, input *api_spec.GetNodeMetrics) ([]api_spec.MetricsGroup, error) {
	return api.driver.GetNode(tctx, input)
}

func (api *Api) GetLogParams(tctx *logger.TraceContext, input *api_spec.GetLogParams) (*api_spec.GetLogParamsData, error) {
	return api.driver.GetLogParams(tctx, input)
}

func (api *Api) GetLogs(tctx *logger.TraceContext, input *api_spec.GetLogs) (*api_spec.GetLogsData, error) {
	return api.driver.GetLogs(tctx, input)
}

func (api *Api) GetEvents(tctx *logger.TraceContext, input *api_spec.GetEvents) (data *api_spec.GetEventsData, err error) {
	return api.driver.GetEvents(tctx, input)
}

func (api *Api) IssueEvent(tctx *logger.TraceContext, input *api_spec.IssueEvent) (err error) {
	return api.driver.IssueEvent(tctx, input)
}

func (api *Api) GetIssuedEvents(tctx *logger.TraceContext, input *api_spec.GetIssuedEvents) (data *api_spec.GetIssuedEventsData, err error) {
	return api.driver.GetIssuedEvents(tctx, input)
}
