package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) ReportNode(tctx *logger.TraceContext, input *api_spec.ReportNode,
	user *base_spec.UserAuthority) (data *api_spec.ReportNodeData, code uint8, err error) {
	if err = resolver.dbApi.ReportNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	if err = resolver.tsdbApi.ReportNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		fmt.Println("DEBUG error report", err)
		return
	}
	fmt.Println("DEBUG logs:", len(input.Logs))
	fmt.Println("DEBUG metrics:", len(input.Metrics))
	code = base_const.CodeOk
	data = &api_spec.ReportNodeData{}
	return
}

func (resolver *Resolver) GetNodes(tctx *logger.TraceContext, input *api_spec.GetNodes,
	user *base_spec.UserAuthority) (data *api_spec.GetNodesData, code uint8, err error) {
	var nodes []spec.Node
	if nodes, err = resolver.dbApi.GetNodes(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetNodesData{Nodes: nodes}
	return
}

func (resolver *Resolver) GetNode(tctx *logger.TraceContext, input *api_spec.GetNode,
	user *base_spec.UserAuthority) (data *api_spec.GetNodeData, code uint8, err error) {
	var node spec.Node
	if node, err = resolver.dbApi.GetNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	var metricsGroups []api_spec.MetricsGroup
	if metricsGroups, err = resolver.tsdbApi.GetNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}

	code = base_const.CodeOk
	node.MetricsGroups = metricsGroups
	data = &spec.GetNodeData{Node: node}
	return
}

func (resolver *Resolver) GetLogs(tctx *logger.TraceContext, input *api_spec.GetLogs, user *base_spec.UserAuthority) (data *api_spec.GetLogsData, code uint8, err error) {
	data, err = resolver.tsdbApi.GetLogs(tctx, input)
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetLogParams(tctx *logger.TraceContext, input *api_spec.GetLogParams, user *base_spec.UserAuthority) (data *api_spec.GetLogParamsData, code uint8, err error) {
	data, err = resolver.tsdbApi.GetLogParams(tctx, input)
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEvents(tctx *logger.TraceContext, input *api_spec.GetEvents, user *base_spec.UserAuthority) (data *api_spec.GetEventsData, code uint8, err error) {
	var getIssuedEventsData *api_spec.GetIssuedEventsData
	getIssuedEventsData, err = resolver.tsdbApi.GetIssuedEvents(tctx, &api_spec.GetIssuedEvents{})
	if err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	data = &api_spec.GetEventsData{
		Events: getIssuedEventsData.Events,
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) CreateEventRules(tctx *logger.TraceContext, input *api_spec.CreateEventRules, user *base_spec.UserAuthority) (data *api_spec.CreateEventRulesData, code uint8, err error) {
	fmt.Println("DEBUG CreateEventRules")
	// data, err = resolver.tsdbApi.CreateEventRules(tctx, input)
	// if err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) UpdateEventRules(tctx *logger.TraceContext, input *api_spec.UpdateEventRules, user *base_spec.UserAuthority) (data *api_spec.UpdateEventRulesData, code uint8, err error) {
	fmt.Println("DEBUG UpdateEventRules")
	// data, err = resolver.tsdbApi.UpdateEventRules(tctx, input)
	// if err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) DeleteEventRules(tctx *logger.TraceContext, input *api_spec.DeleteEventRules, user *base_spec.UserAuthority) (data *api_spec.DeleteEventRulesData, code uint8, err error) {
	fmt.Println("DEBUG DeleteEventRules")
	// data, err = resolver.tsdbApi.DeleteEventRules(tctx, input)
	// if err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) GetEventRules(tctx *logger.TraceContext, input *api_spec.GetEventRules, user *base_spec.UserAuthority) (data *api_spec.GetEventRulesData, code uint8, err error) {
	fmt.Println("DEBUG GetEventRules")
	// data, err = resolver.tsdbApi.GetEventRules(tctx, input)
	// if err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	code = base_const.CodeOk
	return
}
