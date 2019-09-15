package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetPhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalModel) (data *spec.GetPhysicalModelData, code uint8, err error) {
	var region *spec.PhysicalModel
	region, err = resolver.dbApi.GetPhysicalModel(tctx, db, input.Name)
	data = &spec.GetPhysicalModelData{PhysicalModel: *region}
	return
}

func (resolver *Resolver) GetPhysicalModels(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalModels) (data *spec.GetPhysicalModelsData, code uint8, err error) {
	var regions []spec.PhysicalModel
	regions, err = resolver.dbApi.GetPhysicalModels(tctx, db)
	data = &spec.GetPhysicalModelsData{PhysicalModels: regions}
	return
}

func (resolver *Resolver) CreatePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreatePhysicalModel) (data *spec.CreatePhysicalModelData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.PhysicalModel{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalModel" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.PhysicalModel
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.CreatePhysicalModels(tctx, db, specs)
	data = &spec.CreatePhysicalModelData{}
	return
}

func (resolver *Resolver) UpdatePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdatePhysicalModel) (data *spec.UpdatePhysicalModelData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.PhysicalModel{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalModel" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.PhysicalModel
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.UpdatePhysicalModels(tctx, db, specs)
	data = &spec.UpdatePhysicalModelData{}
	return
}

func (resolver *Resolver) DeletePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalModel) (data *spec.DeletePhysicalModelData, code uint8, err error) {
	err = resolver.dbApi.DeletePhysicalModel(tctx, db, input.Name)
	data = &spec.DeletePhysicalModelData{}
	return
}
