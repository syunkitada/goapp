package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetFloor) (data *spec.GetFloorData, code uint8, err error) {
	var region *spec.Floor
	region, err = resolver.dbApi.GetFloor(tctx, db, input.Name)
	data = &spec.GetFloorData{Floor: *region}
	return
}

func (resolver *Resolver) GetFloors(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetFloors) (data *spec.GetFloorsData, code uint8, err error) {
	var regions []spec.Floor
	regions, err = resolver.dbApi.GetFloors(tctx, db)
	data = &spec.GetFloorsData{Floors: regions}
	return
}

func (resolver *Resolver) CreateFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateFloor) (data *spec.CreateFloorData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Floor{}
	for _, base := range baseSpecs {
		if base.Kind != "Floor" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.Floor
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.CreateFloors(tctx, db, specs)
	data = &spec.CreateFloorData{}
	return
}

func (resolver *Resolver) UpdateFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateFloor) (data *spec.UpdateFloorData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Floor{}
	for _, base := range baseSpecs {
		if base.Kind != "Floor" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.Floor
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.UpdateFloors(tctx, db, specs)
	data = &spec.UpdateFloorData{}
	return
}

func (resolver *Resolver) DeleteFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteFloor) (data *spec.DeleteFloorData, code uint8, err error) {
	err = resolver.dbApi.DeleteFloor(tctx, db, input.Name)
	data = &spec.DeleteFloorData{}
	return
}
