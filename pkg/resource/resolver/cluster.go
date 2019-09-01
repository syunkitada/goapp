package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetClusters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetClusters) (data *spec.GetClustersData, code uint8, err error) {
	return
}
