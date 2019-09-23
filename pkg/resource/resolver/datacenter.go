package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetDatacenter) (data *spec.GetDatacenterData, code uint8, err error) {
	var datacenter *spec.Datacenter
	if datacenter, err = resolver.dbApi.GetDatacenter(tctx, db, input); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetDatacenterData{Datacenter: *datacenter}
	return
}

func (resolver *Resolver) GetDatacenters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetDatacenters) (data *spec.GetDatacentersData, code uint8, err error) {
	var datacenters []spec.Datacenter
	if datacenters, err = resolver.dbApi.GetDatacenters(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetDatacentersData{Datacenters: datacenters}
	return
}

func (resolver *Resolver) CreateDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateDatacenter) (data *spec.CreateDatacenterData, code uint8, err error) {
	var specs []spec.Datacenter
	if specs, err = resolver.ConvertToDatacenterSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateDatacenters(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateDatacenterData{}
	return
}

func (resolver *Resolver) UpdateDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateDatacenter) (data *spec.UpdateDatacenterData, code uint8, err error) {
	var specs []spec.Datacenter
	if specs, err = resolver.ConvertToDatacenterSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateDatacenters(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateDatacenterData{}
	return
}

func (resolver *Resolver) DeleteDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteDatacenter) (data *spec.DeleteDatacenterData, code uint8, err error) {
	if err = resolver.dbApi.DeleteDatacenter(tctx, db, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteDatacenterData{}
	return
}

func (resolver *Resolver) DeleteDatacenters(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteDatacenters) (data *spec.DeleteDatacentersData, code uint8, err error) {
	var specs []spec.Datacenter
	if specs, err = resolver.ConvertToDatacenterSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteDatacenters(tctx, db, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteDatacentersData{}
	return
}

func (resolver *Resolver) ConvertToDatacenterSpecs(specStr string) (data []spec.Datacenter, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Datacenter{}
	for _, base := range baseSpecs {
		if base.Kind != "Datacenter" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.Datacenter
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
