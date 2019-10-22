package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetRegionService(tctx *logger.TraceContext, input *spec.GetRegionService, user *base_spec.UserAuthority) (data *spec.GetRegionServiceData, code uint8, err error) {
	var regionService *spec.RegionService
	if regionService, err = resolver.dbApi.GetRegionService(tctx, input, user); err != nil {
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

func (resolver *Resolver) GetRegionServices(tctx *logger.TraceContext, input *spec.GetRegionServices, user *base_spec.UserAuthority) (data *spec.GetRegionServicesData, code uint8, err error) {
	var regionServices []spec.RegionService
	if regionServices, err = resolver.dbApi.GetRegionServices(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionServicesData{RegionServices: regionServices}
	return
}

func (resolver *Resolver) CreateRegionService(tctx *logger.TraceContext, input *spec.CreateRegionService, user *base_spec.UserAuthority) (data *spec.CreateRegionServiceData, code uint8, err error) {
	var specs []spec.RegionService
	if specs, err = resolver.ConvertToRegionServiceSpecs(input.Spec); err != nil {
		fmt.Println(err)
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateRegionServices(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRegionServiceData{}
	return
}

func (resolver *Resolver) UpdateRegionService(tctx *logger.TraceContext, input *spec.UpdateRegionService, user *base_spec.UserAuthority) (data *spec.UpdateRegionServiceData, code uint8, err error) {
	var specs []spec.RegionService
	if specs, err = resolver.ConvertToRegionServiceSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateRegionServices(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRegionServiceData{}
	return
}

func (resolver *Resolver) DeleteRegionService(tctx *logger.TraceContext, input *spec.DeleteRegionService, user *base_spec.UserAuthority) (data *spec.DeleteRegionServiceData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRegionService(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionServiceData{}
	return
}

func (resolver *Resolver) DeleteRegionServices(tctx *logger.TraceContext, input *spec.DeleteRegionServices, user *base_spec.UserAuthority) (data *spec.DeleteRegionServicesData, code uint8, err error) {
	var specs []spec.RegionService
	if specs, err = resolver.ConvertToRegionServiceSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteRegionServices(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionServicesData{}
	return
}

func (resolver *Resolver) ConvertToRegionServiceSpecs(specStr string) (specs []spec.RegionService, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

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
