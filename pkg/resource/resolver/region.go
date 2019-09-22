package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

type BaseSpec struct {
	Kind string
	Spec interface{}
}

func (resolver *Resolver) GetRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegion) (data *spec.GetRegionData, code uint8, err error) {
	var region *spec.Region
	if region, err = resolver.dbApi.GetRegion(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionData{Region: *region}
	return
}

func (resolver *Resolver) GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (data *spec.GetRegionsData, code uint8, err error) {
	var regions []spec.Region
	if regions, err = resolver.dbApi.GetRegions(tctx, db); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionsData{Regions: regions}
	return
}

func (resolver *Resolver) CreateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRegion) (data *spec.CreateRegionData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.Region{}
	for _, base := range baseSpecs {
		if base.Kind != "Region" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.Region
		if err = json.Unmarshal(specBytes, &specData); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		if err = resolver.Validate.Struct(&specData); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		specs = append(specs, specData)
	}

	if err = resolver.dbApi.CreateRegions(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRegionData{}
	return
}

func (resolver *Resolver) UpdateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateRegion) (data *spec.UpdateRegionData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.Region{}
	for _, base := range baseSpecs {
		if base.Kind != "Region" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.Region
		if err = json.Unmarshal(specBytes, &specData); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		if err = resolver.Validate.Struct(&specData); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		specs = append(specs, specData)
	}
	if err = resolver.dbApi.UpdateRegions(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRegionData{}
	return
}

func (resolver *Resolver) DeleteRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRegion) (data *spec.DeleteRegionData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRegion(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionData{}
	return
}
