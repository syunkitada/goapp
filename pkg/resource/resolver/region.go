package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRegion(tctx *logger.TraceContext, input *spec.GetRegion) (data *spec.GetRegionData, code uint8, err error) {
	var region *spec.Region
	if region, err = resolver.dbApi.GetRegion(tctx, input); err != nil {
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

func (resolver *Resolver) GetRegions(tctx *logger.TraceContext, input *spec.GetRegions) (data *spec.GetRegionsData, code uint8, err error) {
	var regions []spec.Region
	fmt.Println("DEBUG GEtREgions")
	if regions, err = resolver.dbApi.GetRegions(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRegionsData{Regions: regions}
	return
}

func (resolver *Resolver) CreateRegion(tctx *logger.TraceContext, input *spec.CreateRegion) (data *spec.CreateRegionData, code uint8, err error) {
	var specs []spec.Region
	if specs, err = resolver.ConvertToRegionSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateRegions(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRegionData{}
	return
}

func (resolver *Resolver) UpdateRegion(tctx *logger.TraceContext, input *spec.UpdateRegion) (data *spec.UpdateRegionData, code uint8, err error) {
	var specs []spec.Region
	if specs, err = resolver.ConvertToRegionSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateRegions(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRegionData{}
	return
}

func (resolver *Resolver) DeleteRegion(tctx *logger.TraceContext, input *spec.DeleteRegion) (data *spec.DeleteRegionData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRegion(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionData{}
	return
}

func (resolver *Resolver) DeleteRegions(tctx *logger.TraceContext, input *spec.DeleteRegions) (data *spec.DeleteRegionsData, code uint8, err error) {
	var specs []spec.Region
	if specs, err = resolver.ConvertToRegionSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteRegions(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRegionsData{}
	return
}

func (resolver *Resolver) ConvertToRegionSpecs(specStr string) (data []spec.Region, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
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
