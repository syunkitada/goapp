package resource_model_api

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetPhysicalModel(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var physicalModel resource_model.PhysicalModel
	if err = db.Where(&resource_model.PhysicalModel{
		Name: resource,
	}).First(&physicalModel).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["PhysicalModel"] = physicalModel
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetPhysicalModels(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	var physicalModels []resource_model.PhysicalModel
	if err = db.Find(&physicalModels).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["PhysicalModels"] = physicalModels
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreatePhysicalModel(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.PhysicalModelSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestError("Specs is empty")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		var data resource_model.PhysicalModel
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.PhysicalModel{
				Kind:        spec.Kind,
				Name:        spec.Name,
				Description: spec.Description,
				Unit:        spec.Unit,
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

func (modelApi *ResourceModelApi) UpdatePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.PhysicalModelSpec
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
		physicalModel := &resource_model.PhysicalModel{
			Kind:        spec.Kind,
			Description: spec.Description,
			Unit:        spec.Unit,
		}
		if err = tx.Model(physicalModel).Where("name = ?", spec.Name).Updates(physicalModel).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeletePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
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

		if err = tx.Delete(&resource_model.PhysicalModel{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}
