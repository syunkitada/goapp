package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	"github.com/syunkitada/goapp/pkg/home/home_api/spec"
)

func (resolver *Resolver) GetProjectUsers(tctx *logger.TraceContext, input *spec.GetProjectUsers, user *base_spec.UserAuthority) (data *spec.GetProjectUsersData, code uint8, err error) {
	var tmpUsers []base_db_model.CustomUser
	if tmpUsers, err = resolver.authproxyDbApi.GetUsers(tctx, user.ProjectName); err != nil {
		code = base_const.CodeServerInternalError
		err = fmt.Errorf("Failed GetUsers")
		return
	}

	var users []spec.User
	for _, user := range tmpUsers {
		users = append(users, spec.User{Name: user.Name, RoleName: user.RoleName})
	}
	data = &spec.GetProjectUsersData{
		Users: users,
	}
	code = base_const.CodeOk
	return
}

func (resolver *Resolver) UpdateUserPassword(tctx *logger.TraceContext, input *spec.UpdateUserPassword, user *base_spec.UserAuthority) (data *spec.UpdateUserPasswordData, code uint8, err error) {
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

	code = base_const.CodeOk
	return
}
