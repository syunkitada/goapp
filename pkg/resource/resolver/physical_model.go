package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetPhysicalModel(tctx *logger.TraceContext, input *spec.GetPhysicalModel, user *base_spec.UserAuthority) (data *spec.GetPhysicalModelData, code uint8, err error) {
	var physicalModel *spec.PhysicalModel
	if physicalModel, err = resolver.dbApi.GetPhysicalModel(tctx, input, user); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetPhysicalModelData{PhysicalModel: *physicalModel}
	return
}

func (resolver *Resolver) GetPhysicalModels(tctx *logger.TraceContext, input *spec.GetPhysicalModels, user *base_spec.UserAuthority) (data *spec.GetPhysicalModelsData, code uint8, err error) {
	var physicalModels []spec.PhysicalModel
	if physicalModels, err = resolver.dbApi.GetPhysicalModels(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetPhysicalModelsData{PhysicalModels: physicalModels}
	return
}

func (resolver *Resolver) CreatePhysicalModel(tctx *logger.TraceContext, input *spec.CreatePhysicalModel, user *base_spec.UserAuthority) (data *spec.CreatePhysicalModelData, code uint8, err error) {
	var specs []spec.PhysicalModel
	if specs, err = resolver.ConvertToPhysicalModelSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreatePhysicalModels(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreatePhysicalModelData{}
	return
}

func (resolver *Resolver) UpdatePhysicalModel(tctx *logger.TraceContext, input *spec.UpdatePhysicalModel, user *base_spec.UserAuthority) (data *spec.UpdatePhysicalModelData, code uint8, err error) {
	var specs []spec.PhysicalModel
	if specs, err = resolver.ConvertToPhysicalModelSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdatePhysicalModels(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdatePhysicalModelData{}
	return
}

func (resolver *Resolver) DeletePhysicalModel(tctx *logger.TraceContext, input *spec.DeletePhysicalModel, user *base_spec.UserAuthority) (data *spec.DeletePhysicalModelData, code uint8, err error) {
	if err = resolver.dbApi.DeletePhysicalModel(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeletePhysicalModelData{}
	return
}

func (resolver *Resolver) DeletePhysicalModels(tctx *logger.TraceContext, input *spec.DeletePhysicalModels, user *base_spec.UserAuthority) (data *spec.DeletePhysicalModelsData, code uint8, err error) {
	var specs []spec.PhysicalModel
	if specs, err = resolver.ConvertToPhysicalModelSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeletePhysicalModels(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeletePhysicalModelsData{}
	return
}

func (resolver *Resolver) ConvertToPhysicalModelSpecs(specStr string) (specs []spec.PhysicalModel, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	for _, base := range baseSpecs {
		if base.Kind != "PhysicalModel" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.PhysicalModel
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
