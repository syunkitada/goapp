package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetRegion(tctx *logger.TraceContext, input *spec.GetRegion, user *base_spec.UserAuthority) (data *spec.GetRegionData, code uint8, err error) {
	var region *spec.Region
	if region, err = resolver.dbApi.GetRegion(tctx, input, user); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionData{Region: *region}
	return
}

func (resolver *Resolver) GetRegions(tctx *logger.TraceContext, input *spec.GetRegions, user *base_spec.UserAuthority) (data *spec.GetRegionsData, code uint8, err error) {
	var regions []spec.Region
	if regions, err = resolver.dbApi.GetRegions(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionsData{Regions: regions}
	return
}

func (resolver *Resolver) CreateRegion(tctx *logger.TraceContext, input *spec.CreateRegion, user *base_spec.UserAuthority) (data *spec.CreateRegionData, code uint8, err error) {
	var specs []spec.Region
	if specs, err = resolver.ConvertToRegionSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateRegions(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRegionData{}
	return
}

func (resolver *Resolver) UpdateRegion(tctx *logger.TraceContext, input *spec.UpdateRegion, user *base_spec.UserAuthority) (data *spec.UpdateRegionData, code uint8, err error) {
	var specs []spec.Region
	if specs, err = resolver.ConvertToRegionSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateRegions(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRegionData{}
	return
}

func (resolver *Resolver) DeleteRegion(tctx *logger.TraceContext, input *spec.DeleteRegion, user *base_spec.UserAuthority) (data *spec.DeleteRegionData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRegion(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionData{}
	return
}

func (resolver *Resolver) DeleteRegions(tctx *logger.TraceContext, input *spec.DeleteRegions, user *base_spec.UserAuthority) (data *spec.DeleteRegionsData, code uint8, err error) {
	var specs []spec.Region
	if specs, err = resolver.ConvertToRegionSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteRegions(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionsData{}
	return
}

func (resolver *Resolver) ConvertToRegionSpecs(specStr string) (specs []spec.Region, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	for _, base := range baseSpecs {
		if base.Kind != "Region" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.Region
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
