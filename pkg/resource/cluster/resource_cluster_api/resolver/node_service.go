package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetNodeServices(tctx *logger.TraceContext, input *api_spec.GetNodeServices,
	user *base_spec.UserAuthority) (data *api_spec.GetNodeServicesData, code uint8, err error) {
	var nodes []base_spec.NodeService
	if nodes, err = resolver.dbApi.GetNodeServices(tctx, &base_spec.GetNodeServices{}, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &api_spec.GetNodeServicesData{NodeServices: nodes}
	return
}

func (resolver *Resolver) SyncNodeService(tctx *logger.TraceContext, input *api_spec.SyncNodeService,
	user *base_spec.UserAuthority) (data *api_spec.SyncNodeServiceData, code uint8, err error) {
	var nodeTask *api_spec.NodeServiceTask
	if nodeTask, err = resolver.dbApi.SyncNodeService(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &api_spec.SyncNodeServiceData{
		Task: *nodeTask,
	}
	return
}

func (resolver *Resolver) ReportNodeServiceTask(tctx *logger.TraceContext, input *api_spec.ReportNodeServiceTask,
	user *base_spec.UserAuthority) (data *api_spec.ReportNodeServiceTaskData, code uint8, err error) {
	if err = resolver.dbApi.ReportNodeServiceTask(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &api_spec.ReportNodeServiceTaskData{}
	return
}
