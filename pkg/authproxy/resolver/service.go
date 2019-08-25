package resolver

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateService) (data *spec.UpdateServiceData, code uint8, err error) {
	return
}
