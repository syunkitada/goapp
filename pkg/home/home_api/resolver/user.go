package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	"github.com/syunkitada/goapp/pkg/home/home_api/spec"
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

func (resolver *Resolver) UpdateUserPassword(tctx *logger.TraceContext, input *spec.UpdateUserPassword, user *base_spec.UserAuthority) (data *spec.UpdateUserPasswordData, code uint8, err error) {
	fmt.Println("DEBUG UpdateUserPassword")
	data = &spec.UpdateUserPasswordData{}
	if input.NewPassword != input.NewPasswordConfirm {
		code = base_const.CodeClientBadRequest
		err = error_utils.NewInvalidRequestError("NewPassword and NewPasswordConfirm is not equal")
		return
	}

	if _, code, err = resolver.authproxyDbApi.GetUserWithValidatePassword(tctx, user.Name, input.CurrentPassword); err != nil {
		return
	}

	if err = resolver.authproxyDbApi.UpdateUserPassword(tctx, user.Name, input.NewPassword); err != nil {
		code = base_const.CodeServerInternalError
		err = fmt.Errorf("Failed UpdateUserPassword")
	}

	return
}
