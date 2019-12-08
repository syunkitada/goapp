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
	code = base_const.CodeOk
	data = &api_spec.GetLogsData{}
	return
}

func (resolver *Resolver) GetLogParams(tctx *logger.TraceContext, input *api_spec.GetLogParams, user *base_spec.UserAuthority) (data *api_spec.GetLogParamsData, code uint8, err error) {
	code = base_const.CodeOk
	data = &api_spec.GetLogParamsData{
		LogNodes: []string{"piyohoge", "piyo"},
		LogApps:  []string{"hogeapp", "piyoapp"},
	}
	return
}
