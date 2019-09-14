package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetDatacenter) (data *spec.GetDatacenterData, code uint8, err error) {
	var region *spec.Datacenter
	region, err = resolver.dbApi.GetDatacenter(tctx, db, input.Name)
	data = &spec.GetDatacenterData{Datacenter: *region}
	return
}

func (resolver *Resolver) GetDatacenters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetDatacenters) (data *spec.GetDatacentersData, code uint8, err error) {
	var regions []spec.Datacenter
	regions, err = resolver.dbApi.GetDatacenters(tctx, db)
	data = &spec.GetDatacentersData{Datacenters: regions}
	return
}

func (resolver *Resolver) CreateDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateDatacenter) (data *spec.CreateDatacenterData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
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
		var region spec.Datacenter
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.CreateDatacenters(tctx, db, specs)
	data = &spec.CreateDatacenterData{}
	return
}

func (resolver *Resolver) UpdateDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateDatacenter) (data *spec.UpdateDatacenterData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
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
		var region spec.Datacenter
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.UpdateDatacenters(tctx, db, specs)
	data = &spec.UpdateDatacenterData{}
	return
}

func (resolver *Resolver) DeleteDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteDatacenter) (data *spec.DeleteDatacenterData, code uint8, err error) {
	err = resolver.dbApi.DeleteDatacenter(tctx, db, input.Name)
	data = &spec.DeleteDatacenterData{}
	return
}
