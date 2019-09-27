package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRack(tctx *logger.TraceContext, input *spec.GetRack) (data *spec.GetRackData, code uint8, err error) {
	var rack *spec.Rack
	if rack, err = resolver.dbApi.GetRack(tctx, input); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRackData{Rack: *rack}
	return
}

func (resolver *Resolver) GetRacks(tctx *logger.TraceContext, input *spec.GetRacks) (data *spec.GetRacksData, code uint8, err error) {
	var racks []spec.Rack
	if racks, err = resolver.dbApi.GetRacks(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRacksData{Racks: racks}
	return
}

func (resolver *Resolver) CreateRack(tctx *logger.TraceContext, input *spec.CreateRack) (data *spec.CreateRackData, code uint8, err error) {
	var specs []spec.Rack
	if specs, err = resolver.ConvertToRackSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateRacks(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRackData{}
	return
}

func (resolver *Resolver) UpdateRack(tctx *logger.TraceContext, input *spec.UpdateRack) (data *spec.UpdateRackData, code uint8, err error) {
	var specs []spec.Rack
	if specs, err = resolver.ConvertToRackSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateRacks(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRackData{}
	return
}

func (resolver *Resolver) DeleteRack(tctx *logger.TraceContext, input *spec.DeleteRack) (data *spec.DeleteRackData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRack(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRackData{}
	return
}

func (resolver *Resolver) DeleteRacks(tctx *logger.TraceContext, input *spec.DeleteRacks) (data *spec.DeleteRacksData, code uint8, err error) {
	var specs []spec.Rack
	if specs, err = resolver.ConvertToRackSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteRacks(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRacksData{}
	return
}

func (resolver *Resolver) ConvertToRackSpecs(specStr string) (data []spec.Rack, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Rack{}
	for _, base := range baseSpecs {
		if base.Kind != "Rack" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.Rack
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
