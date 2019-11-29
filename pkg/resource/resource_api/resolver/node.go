package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetNodes(tctx *logger.TraceContext, input *spec.GetNodes, user *base_spec.UserAuthority) (data *spec.GetNodesData, code uint8, err error) {
	var nodes []spec.Node
	if nodes, err = resolver.dbApi.GetNodes(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetNodesData{Nodes: nodes}
	return
}

func (resolver *Resolver) GetNode(tctx *logger.TraceContext, input *spec.GetNode, user *base_spec.UserAuthority) (data *spec.GetNodeData, code uint8, err error) {
	var node *spec.Node
	if node, err = resolver.dbApi.GetNode(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetNodeData{Node: *node}
	return
}

func (resolver *Resolver) GetAlerts(tctx *logger.TraceContext, input *spec.GetAlerts, user *base_spec.UserAuthority) (data *spec.GetAlertsData, code uint8, err error) {
	var alerts = []spec.ResourceAlert{
		spec.ResourceAlert{
			Name:    "hoge",
			Time:    "timestamp",
			Level:   "Critical",
			Handler: "handlerhoge",
			Msg:     "critical on host",
			Tag:     map[string]string{},
		},
		spec.ResourceAlert{
			Name:    "piyo",
			Time:    "timestamp",
			Level:   "Warning",
			Handler: "handlerhoge",
			Msg:     "critical on host",
			Tag:     map[string]string{},
		},
	}
	code = base_const.CodeOk
	data = &spec.GetAlertsData{
		Alerts: alerts,
	}
	return
}
func (resolver *Resolver) GetAlertRules(tctx *logger.TraceContext, input *spec.GetAlertRules, user *base_spec.UserAuthority) (data *spec.GetAlertRulesData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetAlertRulesData{}
	return
}
func (resolver *Resolver) GetStatistics(tctx *logger.TraceContext, input *spec.GetStatistics, user *base_spec.UserAuthority) (data *spec.GetStatisticsData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetStatisticsData{}
	return
}

func (resolver *Resolver) GetLogs(tctx *logger.TraceContext, input *spec.GetLogs, user *base_spec.UserAuthority) (data *spec.GetLogsData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetLogsData{}
	return
}

func (resolver *Resolver) GetTrace(tctx *logger.TraceContext, input *spec.GetTrace, user *base_spec.UserAuthority) (data *spec.GetTraceData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetTraceData{}
	return
}
