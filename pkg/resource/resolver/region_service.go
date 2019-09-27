package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRegionService(tctx *logger.TraceContext, input *spec.GetRegionService) (data *spec.GetRegionServiceData, code uint8, err error) {
	var regionService *spec.RegionService
	if regionService, err = resolver.dbApi.GetRegionService(tctx, input); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionServiceData{RegionService: *regionService}
	return
}

func (resolver *Resolver) GetRegionServices(tctx *logger.TraceContext, input *spec.GetRegionServices) (data *spec.GetRegionServicesData, code uint8, err error) {
	var regionServices []spec.RegionService
	if regionServices, err = resolver.dbApi.GetRegionServices(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionServicesData{RegionServices: regionServices}
	return
}

func (resolver *Resolver) CreateRegionService(tctx *logger.TraceContext, input *spec.CreateRegionService) (data *spec.CreateRegionServiceData, code uint8, err error) {
	var specs []spec.RegionService
	if specs, err = resolver.ConvertToRegionServiceSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateRegionServices(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRegionServiceData{}
	return
}

func (resolver *Resolver) UpdateRegionService(tctx *logger.TraceContext, input *spec.UpdateRegionService) (data *spec.UpdateRegionServiceData, code uint8, err error) {
	var specs []spec.RegionService
	if specs, err = resolver.ConvertToRegionServiceSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateRegionServices(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRegionServiceData{}
	return
}

func (resolver *Resolver) DeleteRegionService(tctx *logger.TraceContext, input *spec.DeleteRegionService) (data *spec.DeleteRegionServiceData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRegionService(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionServiceData{}
	return
}

func (resolver *Resolver) DeleteRegionServices(tctx *logger.TraceContext, input *spec.DeleteRegionServices) (data *spec.DeleteRegionServicesData, code uint8, err error) {
	var specs []spec.RegionService
	if specs, err = resolver.ConvertToRegionServiceSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteRegionServices(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionServicesData{}
	return
}

func (resolver *Resolver) ConvertToRegionServiceSpecs(specStr string) (data []spec.RegionService, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	specs := []spec.RegionService{}
	for _, base := range baseSpecs {
		if base.Kind != "RegionService" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.RegionService
		if err = json.Unmarshal(specBytes, &specData); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&specData); err != nil {
			return
		}
		specs = append(specs, specData)
	}
	return
}
