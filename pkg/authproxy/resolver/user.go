package resolver

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetAllUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetAllUsers) (data *spec.GetAllUsersData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetUser(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUser) (data *spec.GetUserData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUsers) (data *spec.GetUsersData, code uint8, err error) {
	return
}
