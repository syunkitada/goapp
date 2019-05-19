package resource_model_api

import (
	"encoding/json"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetRack(tctx *logger.TraceContext,
	db *gorm.DB, query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.PhysicalActionReply) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("resource is None")
	}

	var floor resource_model.Rack
	if err = db.Where(&resource_model.Rack{
		Name: resource,
	}).First(&floor).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Rack = modelApi.convertRack(tctx, &floor)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetRacks(tctx *logger.TraceContext,
	db *gorm.DB, query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.PhysicalActionReply) (int64, error) {
	var err error
	datacenter, ok := query.StrParams["datacenter"]
	if !ok || datacenter == "" {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("datacenter is None")
	}

	var floors []resource_model.Rack
	if err = db.Where("datacenter = ?", datacenter).Find(&floors).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Racks = modelApi.convertRacks(tctx, floors)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateRack(tctx *logger.TraceContext,
	db *gorm.DB, query *resource_api_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.RackSpecData
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.Rack
		if err = tx.Where("name = ? and datacenter = ?", spec.Name, spec.Datacenter).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Rack{
				Kind:       spec.Kind,
				Name:       spec.Name,
				Datacenter: spec.Datacenter,
				Floor:      spec.Floor,
				Unit:       spec.Unit,
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
	return codes.Ok, nil
}

func (modelApi *ResourceModelApi) UpdateRack(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.RackSpecData
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
		rack := &resource_model.Rack{
			Kind:       spec.Kind,
			Datacenter: spec.Datacenter,
			Floor:      spec.Floor,
			Unit:       spec.Unit,
		}
		if err = tx.Model(rack).Where("name = ?", spec.Name).Updates(rack).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteRack(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (int64, error) {
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

		if err = tx.Delete(&resource_model.Rack{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) convertRack(tctx *logger.TraceContext,
	floor *resource_model.Rack) *resource_api_grpc_pb.Rack {
	updatedAt, err := ptypes.TimestampProto(floor.Model.UpdatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", floor.Model.UpdatedAt)
	}
	createdAt, err := ptypes.TimestampProto(floor.Model.CreatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", floor.Model.CreatedAt)
	}

	return &resource_api_grpc_pb.Rack{
		Name:      floor.Name,
		Kind:      floor.Kind,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
}

func (modelApi *ResourceModelApi) convertRacks(tctx *logger.TraceContext,
	floors []resource_model.Rack) []*resource_api_grpc_pb.Rack {
	pbRacks := make([]*resource_api_grpc_pb.Rack, len(floors))
	for i, floor := range floors {
		pbRacks[i] = modelApi.convertRack(tctx, &floor)
	}

	return pbRacks
}
