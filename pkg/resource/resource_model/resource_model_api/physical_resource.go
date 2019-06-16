package resource_model_api

import (
	"encoding/json"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetPhysicalResource(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("resource is None")
	}

	var physicalModel resource_model.PhysicalResource
	if err = db.Where(&resource_model.PhysicalResource{
		Name: resource,
	}).First(&physicalModel).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["PhysicalResource"] = physicalModel
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetPhysicalResources(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	datacenter, ok := query.StrParams["datacenter"]
	if !ok || datacenter == "" {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("datacenter is None")
	}

	var physicalResources []resource_model.PhysicalResource
	if err = db.Where("datacenter = ?", datacenter).Find(&physicalResources).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["PhysicalResources"] = physicalResources
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreatePhysicalResource(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.PhysicalResourceSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.PhysicalResource
		if err = tx.Where("name = ? and datacenter = ?", spec.Name, spec.Datacenter).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.PhysicalResource{
				Kind:          spec.Kind,
				Name:          spec.Name,
				Datacenter:    spec.Datacenter,
				Cluster:       spec.Cluster,
				Rack:          spec.Rack,
				PhysicalModel: spec.Model,
				RackPosition:  spec.RackPosition,
				PowerLinks:    strings.Join(spec.PowerLinks, ","),
				NetLinks:      strings.Join(spec.NetLinks, ","),
				Spec:          spec.Spec,
			}
			if err = tx.Create(&data).Error; err != nil {
				return codes.RemoteDbError, err
			}
		} else {
			err = error_utils.NewConflictAlreadyExistsError(spec.Name)
			return codes.ClientAlreadyExists, err
		}
	}

	tx.Commit()
	return codes.OkCreated, nil
}

func (modelApi *ResourceModelApi) UpdatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.PhysicalResourceSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}
		physicalModel := &resource_model.PhysicalResource{
			Kind:          spec.Kind,
			Datacenter:    spec.Datacenter,
			Cluster:       spec.Cluster,
			Rack:          spec.Rack,
			PhysicalModel: spec.Model,
			RackPosition:  spec.RackPosition,
			PowerLinks:    strings.Join(spec.PowerLinks, ","),
			NetLinks:      strings.Join(spec.NetLinks, ","),
			Spec:          spec.Spec,
		}
		if err = tx.Model(physicalModel).Where("name = ?", spec.Name).Updates(physicalModel).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeletePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NameSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		if err = tx.Delete(&resource_model.PhysicalResource{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) convertPhysicalResource(tctx *logger.TraceContext,
	physicalResource *resource_model.PhysicalResource) *resource_api_grpc_pb.PhysicalResource {
	updatedAt, err := ptypes.TimestampProto(physicalResource.Model.UpdatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", physicalResource.Model.UpdatedAt)
	}
	createdAt, err := ptypes.TimestampProto(physicalResource.Model.CreatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", physicalResource.Model.CreatedAt)
	}

	return &resource_api_grpc_pb.PhysicalResource{
		Name:      physicalResource.Name,
		Kind:      physicalResource.Kind,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
}

func (modelApi *ResourceModelApi) convertPhysicalResources(tctx *logger.TraceContext,
	physicalResources []resource_model.PhysicalResource) []*resource_api_grpc_pb.PhysicalResource {
	pbPhysicalResources := make([]*resource_api_grpc_pb.PhysicalResource, len(physicalResources))
	for i, physicalResource := range physicalResources {
		pbPhysicalResources[i] = modelApi.convertPhysicalResource(tctx, &physicalResource)
	}

	return pbPhysicalResources
}
