package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetPhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResource) (data *spec.GetPhysicalResourceData, code uint8, err error) {
	var region *spec.PhysicalResource
	if region, err = resolver.dbApi.GetPhysicalResource(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetPhysicalResourceData{PhysicalResource: *region}
	return
}

func (resolver *Resolver) GetPhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResources) (data *spec.GetPhysicalResourcesData, code uint8, err error) {
	var regions []spec.PhysicalResource
	if regions, err = resolver.dbApi.GetPhysicalResources(tctx, db); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetPhysicalResourcesData{PhysicalResources: regions}
	return
}

func (resolver *Resolver) CreatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreatePhysicalResource) (data *spec.CreatePhysicalResourceData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.PhysicalResource{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalResource" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.PhysicalResource
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

	if err = resolver.dbApi.CreatePhysicalResources(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreatePhysicalResourceData{}
	return
}

func (resolver *Resolver) UpdatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdatePhysicalResource) (data *spec.UpdatePhysicalResourceData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.PhysicalResource{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalResource" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.PhysicalResource
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
	if err = resolver.dbApi.UpdatePhysicalResources(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdatePhysicalResourceData{}
	return
}

func (resolver *Resolver) DeletePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalResource) (data *spec.DeletePhysicalResourceData, code uint8, err error) {
	if err = resolver.dbApi.DeletePhysicalResource(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeletePhysicalResourceData{}
	return
}

func (resolver *Resolver) DeletePhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalResources) (data *spec.DeletePhysicalResourcesData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.PhysicalResource{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalResource" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.PhysicalResource
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
	if err = resolver.dbApi.DeletePhysicalResources(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeletePhysicalResourcesData{}
	return
}
