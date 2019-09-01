package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	return
}
