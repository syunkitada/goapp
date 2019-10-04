package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetCompute(tctx *logger.TraceContext, input *spec.GetCompute, user *base_spec.UserAuthority) (data *spec.GetComputeData, code uint8, err error) {
	// var image *spec.Compute
	// if image, err = resolver.dbApi.GetCompute(tctx, input, user); err != nil {
	// 	if gorm.IsRecordNotFoundError(err) {
	// 		code = base_const.CodeOkNotFound
	// 		return
	// 	}
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	// code = base_const.CodeOk
	// data = &spec.GetComputeData{Compute: *image}
	return
}

func (resolver *Resolver) GetComputes(tctx *logger.TraceContext, input *spec.GetComputes, user *base_spec.UserAuthority) (data *spec.GetComputesData, code uint8, err error) {
	// var images []spec.Compute
	// if images, err = resolver.dbApi.GetComputes(tctx, input, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	// code = base_const.CodeOk
	// data = &spec.GetComputesData{Computes: images}
	return
}

func (resolver *Resolver) CreateCompute(tctx *logger.TraceContext, input *spec.CreateCompute, user *base_spec.UserAuthority) (data *spec.CreateComputeData, code uint8, err error) {
	// var specs []spec.Compute
	// fmt.Println("DEBUG CreateCompute", input.Spec)
	// if specs, err = resolver.ConvertToComputeSpecs(input.Spec); err != nil {
	// 	code = base_const.CodeClientBadRequest
	// 	return
	// }
	// if err = resolver.dbApi.CreateComputes(tctx, specs, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	// code = base_const.CodeOkCreated
	// data = &spec.CreateComputeData{}
	return
}

func (resolver *Resolver) UpdateCompute(tctx *logger.TraceContext, input *spec.UpdateCompute, user *base_spec.UserAuthority) (data *spec.UpdateComputeData, code uint8, err error) {
	// var specs []spec.Compute
	// if specs, err = resolver.ConvertToComputeSpecs(input.Spec); err != nil {
	// 	code = base_const.CodeClientBadRequest
	// 	return
	// }
	// if err = resolver.dbApi.UpdateComputes(tctx, specs, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	// code = base_const.CodeOkUpdated
	// data = &spec.UpdateComputeData{}
	return
}

func (resolver *Resolver) DeleteCompute(tctx *logger.TraceContext, input *spec.DeleteCompute, user *base_spec.UserAuthority) (data *spec.DeleteComputeData, code uint8, err error) {
	// if err = resolver.dbApi.DeleteCompute(tctx, input, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	// code = base_const.CodeOkDeleted
	// data = &spec.DeleteComputeData{}
	return
}

func (resolver *Resolver) DeleteComputes(tctx *logger.TraceContext, input *spec.DeleteComputes, user *base_spec.UserAuthority) (data *spec.DeleteComputesData, code uint8, err error) {
	// var specs []spec.Compute
	// if specs, err = resolver.ConvertToComputeSpecs(input.Spec); err != nil {
	// 	code = base_const.CodeClientBadRequest
	// 	return
	// }
	// if err = resolver.dbApi.DeleteComputes(tctx, specs, user); err != nil {
	// 	code = base_const.CodeServerInternalError
	// 	return
	// }
	// code = base_const.CodeOkDeleted
	// data = &spec.DeleteComputesData{}
	return
}

func (resolver *Resolver) ConvertToComputeSpecs(specStr string) (specs []spec.Compute, err error) {
	// var baseSpecs []base_spec.Spec
	// if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
	// 	return
	// }

	// for _, base := range baseSpecs {
	// 	if base.Kind != "Compute" {
	// 		continue
	// 	}
	// 	var specBytes []byte
	// 	if specBytes, err = json.Marshal(base.Spec); err != nil {
	// 		return
	// 	}
	// 	var specData spec.Compute
	// 	if err = json.Unmarshal(specBytes, &specData); err != nil {
	// 		return
	// 	}
	// 	if err = resolver.Validate.Struct(&specData); err != nil {
	// 		return
	// 	}
	// 	specs = append(specs, specData)
	// }
	return
}
