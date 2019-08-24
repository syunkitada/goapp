package resolver

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateService) (*spec.UpdateServiceData, error) {
	return nil, nil
}
