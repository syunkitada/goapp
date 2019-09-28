package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetAllUsers(tctx *logger.TraceContext, input *base_spec.GetAllUsers, user *base_spec.UserAuthority) (data *base_spec.GetAllUsersData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetUser(tctx *logger.TraceContext, input *base_spec.GetUser, user *base_spec.UserAuthority) (data *base_spec.GetUserData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetUsers(tctx *logger.TraceContext, input *base_spec.GetUsers, user *base_spec.UserAuthority) (data *base_spec.GetUsersData, code uint8, err error) {
	return
}
