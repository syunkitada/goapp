package resource_model_api

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetFloor(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("resource is None")
	}

	var floor resource_model.Floor
	if err = db.Where(&resource_model.Floor{
		Name: resource,
	}).First(&floor).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Floor"] = floor
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetFloors(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	datacenter, ok := query.StrParams["datacenter"]
	if !ok || datacenter == "" {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("datacenter is None")
	}

	var floors []resource_model.Floor
	if err = db.Where("datacenter = ?", datacenter).Find(&floors).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Floors"] = floors
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateFloor(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.FloorSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.Floor
		if err = tx.Where("name = ? and datacenter = ?", spec.Name, spec.Datacenter).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Floor{
				Kind:       spec.Kind,
				Name:       spec.Name,
				Datacenter: spec.Datacenter,
				Zone:       spec.Zone,
				Floor:      spec.Floor,
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

func (modelApi *ResourceModelApi) UpdateFloor(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.FloorSpec
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
		physicalModel := &resource_model.Floor{
			Kind:       spec.Kind,
			Datacenter: spec.Datacenter,
			Zone:       spec.Zone,
			Floor:      spec.Floor,
		}
		if err = tx.Model(physicalModel).Where("name = ?", spec.Name).Updates(physicalModel).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteFloor(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
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

		if err = tx.Delete(&resource_model.Floor{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}
