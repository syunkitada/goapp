package resolver

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetAllUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetAllUsers) (*spec.GetAllUsersData, error) {
	return nil, nil
}

func (resolver *Resolver) GetUser(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUser) (*spec.GetUserData, error) {
	return nil, nil
}

func (resolver *Resolver) GetUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUsers) (*spec.GetUsersData, error) {
	return nil, nil
}
