package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	code = base_const.CodeClientNotFound
	return
}

func (resolver *Resolver) GetProjectServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	code = base_const.CodeClientNotFound
	return
}

func (resolver *Resolver) GetServiceDashboardIndex(tctx *logger.TraceContext,
	input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	code = base_const.CodeClientNotFound
	return
}

func (resolver *Resolver) GetProjectServiceDashboardIndex(tctx *logger.TraceContext,
	input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	code = base_const.CodeClientNotFound
	return
}
