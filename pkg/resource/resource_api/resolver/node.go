package resolver

import (
	"time"

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

func (resolver *Resolver) GetEvents(tctx *logger.TraceContext, input *spec.GetEvents, user *base_spec.UserAuthority) (data *spec.GetEventsData, code uint8, err error) {
	if data, err = resolver.dbApi.GetEvents(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) CreateEventRules(tctx *logger.TraceContext, input *spec.CreateEventRules, user *base_spec.UserAuthority) (data *spec.CreateEventRulesData, code uint8, err error) {
	// if data, err = resolver.dbApi.CreateEventRules(tctx, input, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) UpdateEventRules(tctx *logger.TraceContext, input *spec.UpdateEventRules, user *base_spec.UserAuthority) (data *spec.UpdateEventRulesData, code uint8, err error) {
	// if data, err = resolver.dbApi.UpdateEventRules(tctx, input, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) DeleteEventRules(tctx *logger.TraceContext, input *spec.DeleteEventRules, user *base_spec.UserAuthority) (data *spec.DeleteEventRulesData, code uint8, err error) {
	// if data, err = resolver.dbApi.DeleteEventRules(tctx, input, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEventRules(tctx *logger.TraceContext, input *spec.GetEventRules, user *base_spec.UserAuthority) (data *spec.GetEventRulesData, code uint8, err error) {
	code = base_const.CodeOk
	var eventRules = []spec.EventRule{
		spec.EventRule{
			Name:  "hoge",
			Node:  ".*",
			Kind:  "Filter",
			Until: time.Now(),
		},
		spec.EventRule{
			Name:  ".*",
			Node:  "hoge.com",
			Kind:  "Filter",
			Until: time.Now(),
		},
	}

	data = &spec.GetEventRulesData{
		EventRules: eventRules,
	}
	return
}
func (resolver *Resolver) GetStatistics(tctx *logger.TraceContext, input *spec.GetStatistics, user *base_spec.UserAuthority) (data *spec.GetStatisticsData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetStatisticsData{}
	return
}

func (resolver *Resolver) GetLogs(tctx *logger.TraceContext, input *spec.GetLogs, user *base_spec.UserAuthority) (data *spec.GetLogsData, code uint8, err error) {
	data = &spec.GetLogsData{}
	if data, err = resolver.dbApi.GetLogs(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetLogParams(tctx *logger.TraceContext, input *spec.GetLogParams, user *base_spec.UserAuthority) (data *spec.GetLogParamsData, code uint8, err error) {
	if data, err = resolver.dbApi.GetLogParams(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetTrace(tctx *logger.TraceContext, input *spec.GetTrace, user *base_spec.UserAuthority) (data *spec.GetTraceData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetTraceData{}
	return
}
