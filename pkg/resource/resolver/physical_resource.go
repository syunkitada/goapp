package resolver

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetPhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResource) (data *spec.GetPhysicalResourceData, code uint8, err error) {
	var region *spec.PhysicalResource
	region, err = resolver.dbApi.GetPhysicalResource(tctx, db, input.Name)
	data = &spec.GetPhysicalResourceData{PhysicalResource: *region}
	fmt.Println("DEBUG GetPhysicalResource", input)
	return
}

func (resolver *Resolver) GetPhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResources) (data *spec.GetPhysicalResourcesData, code uint8, err error) {
	var regions []spec.PhysicalResource
	regions, err = resolver.dbApi.GetPhysicalResources(tctx, db)
	data = &spec.GetPhysicalResourcesData{PhysicalResources: regions}
	fmt.Println("DEBUG Get PhysicalResources")
	return
}

func (resolver *Resolver) CreatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreatePhysicalResource) (data *spec.CreatePhysicalResourceData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.PhysicalResource{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalResource" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.PhysicalResource
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.CreatePhysicalResources(tctx, db, specs)
	data = &spec.CreatePhysicalResourceData{}
	return
}

func (resolver *Resolver) UpdatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdatePhysicalResource) (data *spec.UpdatePhysicalResourceData, code uint8, err error) {
	var baseSpecs []BaseSpec
	if err = json.Unmarshal([]byte(input.Spec), &baseSpecs); err != nil {
		return
	}

	specs := []spec.PhysicalResource{}
	for _, base := range baseSpecs {
		if base.Kind != "PhysicalResource" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var region spec.PhysicalResource
		if err = json.Unmarshal(specBytes, &region); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&region); err != nil {
			return
		}
		specs = append(specs, region)
	}
	err = resolver.dbApi.UpdatePhysicalResources(tctx, db, specs)
	data = &spec.UpdatePhysicalResourceData{}
	return
}

func (resolver *Resolver) DeletePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalResource) (data *spec.DeletePhysicalResourceData, code uint8, err error) {
	err = resolver.dbApi.DeletePhysicalResource(tctx, db, input.Name)
	data = &spec.DeletePhysicalResourceData{}
	return
}
