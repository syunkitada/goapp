package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetAllUsers(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetAllUsers) (data *base_spec.GetAllUsersData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetUser(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetUser) (data *base_spec.GetUserData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetUsers(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetUsers) (data *base_spec.GetUsersData, code uint8, err error) {
	return
}
