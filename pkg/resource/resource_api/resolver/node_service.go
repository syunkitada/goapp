package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetNodeServices(tctx *logger.TraceContext, input *spec.GetNodeServices, user *base_spec.UserAuthority) (data *spec.GetNodeServicesData, code uint8, err error) {
	var nodes []base_spec.NodeService
	if nodes, err = resolver.dbApi.GetNodeServices(tctx, &base_spec.GetNodeServices{}, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetNodeServicesData{NodeServices: nodes}
	return
}
