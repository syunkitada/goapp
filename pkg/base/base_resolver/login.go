package base_resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) Login(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.Login) (data *base_spec.LoginData, code uint8, err error) {
	user, code, err := resolver.dbApi.GetUserWithValidatePassword(tctx, db, input.User, input.Password)
	if err != nil {
		return
	}

	token, err := resolver.dbApi.IssueToken(user.Name)
	if err != nil {
		return
	}

	userAuthority, err := resolver.dbApi.GetUserAuthority(tctx, db, input.User)
	if err != nil {
		return
	}

	data = &base_spec.LoginData{
		Token:     token,
		Authority: *userAuthority,
	}
	return
}

func (resolver *Resolver) LoginWithToken(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (data *base_spec.LoginWithTokenData, code uint8, err error) {
	data = &base_spec.LoginWithTokenData{Authority: *user}
	if user == nil {
		code = base_const.CodeClientInvalidAuth
		err = error_utils.NewInvalidAuthError("Invalid Token")
		return
	}
	code = base_const.CodeOk
	return
}
