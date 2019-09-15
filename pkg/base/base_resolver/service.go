package base_resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.UpdateService) (data *base_spec.UpdateServiceData, code uint8, err error) {
	if err = resolver.dbApi.CreateOrUpdateService(tctx, db, input); err != nil {
		return
	}
	data = &base_spec.UpdateServiceData{}
	code = base_const.CodeOk
	return
}
