package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRack) (data *spec.GetRackData, code uint8, err error) {
	var rack *spec.Rack
	if rack, err = resolver.dbApi.GetRack(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRackData{Rack: *rack}
	return
}

func (resolver *Resolver) GetRacks(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRacks) (data *spec.GetRacksData, code uint8, err error) {
	var racks []spec.Rack
	if racks, err = resolver.dbApi.GetRacks(tctx, db); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetRacksData{Racks: racks}
	return
}

func (resolver *Resolver) CreateRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRack) (data *spec.CreateRackData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.Rack{}
	for _, base := range baseSpecs {
		if base.Kind != "Rack" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.Rack
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

	if err = resolver.dbApi.CreateRacks(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateRackData{}
	return
}

func (resolver *Resolver) UpdateRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateRack) (data *spec.UpdateRackData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}

	specs := []spec.Rack{}
	for _, base := range baseSpecs {
		if base.Kind != "Rack" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			code = base_const.CodeClientBadRequest
			return
		}
		var specData spec.Rack
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
	if err = resolver.dbApi.UpdateRacks(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateRackData{}
	return
}

func (resolver *Resolver) DeleteRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRack) (data *spec.DeleteRackData, code uint8, err error) {
	if err = resolver.dbApi.DeleteRack(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteRackData{}
	return
}
