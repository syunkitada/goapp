package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRack) (data *spec.GetRackData, code uint8, err error) {
	var region *spec.Rack
	region, err = resolver.dbApi.GetRack(tctx, db, input.Name)
	data = &spec.GetRackData{Rack: *region}
	return
}

func (resolver *Resolver) GetRacks(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRacks) (data *spec.GetRacksData, code uint8, err error) {
	var regions []spec.Rack
	regions, err = resolver.dbApi.GetRacks(tctx, db)
	data = &spec.GetRacksData{Racks: regions}
	return
}

func (resolver *Resolver) CreateRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRack) (data *spec.CreateRackData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Rack{}
	for _, base := range baseSpecs {
		if base.Kind != "Rack" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.Rack
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.CreateRacks(tctx, db, specs)
	data = &spec.CreateRackData{}
	return
}

func (resolver *Resolver) UpdateRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateRack) (data *spec.UpdateRackData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Rack{}
	for _, base := range baseSpecs {
		if base.Kind != "Rack" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.Rack
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.UpdateRacks(tctx, db, specs)
	data = &spec.UpdateRackData{}
	return
}

func (resolver *Resolver) DeleteRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRack) (data *spec.DeleteRackData, code uint8, err error) {
	err = resolver.dbApi.DeleteRack(tctx, db, input.Name)
	data = &spec.DeleteRackData{}
	return
}
