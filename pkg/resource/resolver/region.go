package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegion) (data *spec.GetRegionData, code uint8, err error) {
	data = &spec.GetRegionData{}
	return
}

func (resolver *Resolver) GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (data *spec.GetRegionsData, code uint8, err error) {
	data = &spec.GetRegionsData{}
	return
}

func (resolver *Resolver) CreateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRegion) (data *spec.CreateRegionData, code uint8, err error) {
	data = &spec.CreateRegionData{}
	return
}

func (resolver *Resolver) UpdateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateRegion) (data *spec.UpdateRegionData, code uint8, err error) {
	data = &spec.UpdateRegionData{}
	return
}

func (resolver *Resolver) DeleteRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRegion) (data *spec.DeleteRegionData, code uint8, err error) {
	data = &spec.DeleteRegionData{}
	return
}
