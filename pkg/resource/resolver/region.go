package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

type BaseSpec struct {
	Kind string
	Spec interface{}
}

func (resolver *Resolver) GetRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegion) (data *spec.GetRegionData, code uint8, err error) {
	data = &spec.GetRegionData{}
	return
}

func (resolver *Resolver) GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (data *spec.GetRegionsData, code uint8, err error) {
	var regions []spec.Region
	regions, err = resolver.dbApi.GetRegions(tctx, db)
	data = &spec.GetRegionsData{Regions: regions}
	return
}

func (resolver *Resolver) CreateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRegion) (data *spec.CreateRegionData, code uint8, err error) {
	fmt.Printf("DEBUG CreateRegion: %v\n", input)
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Region{}
	for _, base := range baseSpecs {
		if base.Kind != "Region" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.Region
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.CreateRegions(tctx, db, specs)
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
