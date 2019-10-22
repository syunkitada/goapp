package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetNodes(tctx *logger.TraceContext, input *api_spec.GetNodes,
	user *base_spec.UserAuthority) (data *api_spec.GetNodesData, code uint8, err error) {
	var nodes []base_spec.Node
	if nodes, err = resolver.dbApi.GetNodes(tctx, &base_spec.GetNodes{}, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &api_spec.GetNodesData{Nodes: nodes}
	return
}

func (resolver *Resolver) SyncNode(tctx *logger.TraceContext, input *api_spec.SyncNode,
	user *base_spec.UserAuthority) (data *api_spec.SyncNodeData, code uint8, err error) {
	var nodeTask *api_spec.NodeTask
	if nodeTask, err = resolver.dbApi.SyncNode(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &api_spec.SyncNodeData{
		Task: *nodeTask,
	}
	return
}

func (resolver *Resolver) ReportNodeTask(tctx *logger.TraceContext, input *api_spec.ReportNodeTask,
	user *base_spec.UserAuthority) (data *api_spec.ReportNodeTaskData, code uint8, err error) {
	if err = resolver.dbApi.ReportNodeTask(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &api_spec.ReportNodeTaskData{}
	return
}
