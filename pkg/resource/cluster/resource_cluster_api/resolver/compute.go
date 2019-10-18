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

func (resolver *Resolver) GetCompute(tctx *logger.TraceContext, input *spec.GetCompute, user *base_spec.UserAuthority) (data *spec.GetComputeData, code uint8, err error) {
	var regionService *spec.Compute
	if regionService, err = resolver.dbApi.GetCompute(tctx, input); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetComputeData{Compute: *regionService}
	return
}

func (resolver *Resolver) GetComputes(tctx *logger.TraceContext, input *spec.GetComputes, user *base_spec.UserAuthority) (data *spec.GetComputesData, code uint8, err error) {
	var computes []spec.Compute
	if computes, err = resolver.dbApi.GetComputes(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetComputesData{Computes: computes}
	return
}

func (resolver *Resolver) CreateCompute(tctx *logger.TraceContext, input *spec.CreateCompute, user *base_spec.UserAuthority) (data *spec.CreateComputeData, code uint8, err error) {
	fmt.Println("DEBUG input", input.Spec)
	var specs []spec.RegionServiceComputeSpec
	if specs, err = resolver.ConvertToComputeSpecs(input.Spec); err != nil {
		fmt.Println(err)
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateComputes(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateComputeData{}
	return
}

func (resolver *Resolver) UpdateCompute(tctx *logger.TraceContext, input *spec.UpdateCompute, user *base_spec.UserAuthority) (data *spec.UpdateComputeData, code uint8, err error) {
	var specs []spec.RegionServiceComputeSpec
	if specs, err = resolver.ConvertToComputeSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateComputes(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateComputeData{}
	return
}

func (resolver *Resolver) DeleteCompute(tctx *logger.TraceContext, input *spec.DeleteCompute, user *base_spec.UserAuthority) (data *spec.DeleteComputeData, code uint8, err error) {
	if err = resolver.dbApi.DeleteCompute(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteComputeData{}
	return
}

func (resolver *Resolver) DeleteComputes(tctx *logger.TraceContext, input *spec.DeleteComputes, user *base_spec.UserAuthority) (data *spec.DeleteComputesData, code uint8, err error) {
	var specs []spec.RegionServiceComputeSpec
	if specs, err = resolver.ConvertToComputeSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteComputes(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteComputesData{}
	return
}

func (resolver *Resolver) ConvertToComputeSpecs(specStr string) (specs []spec.RegionServiceComputeSpec, err error) {
	err = json.Unmarshal([]byte(specStr), &specs)
	return
}
