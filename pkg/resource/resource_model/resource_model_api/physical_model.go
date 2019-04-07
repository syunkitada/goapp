package resource_model_api

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) CreatePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (error, int64) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Specs")
		return error_utils.NewInvalidRequestError("NotFound Specs"), codes.ClientBadRequest
	}
	fmt.Println("DEBUG CreatePhysicalModel1")

	var specs []resource_model.PhysicalModelSpecData
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return err, codes.ClientBadRequest
	}
	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()
	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return err, codes.ClientBadRequest
		}

		var data resource_model.PhysicalModel
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return err, codes.RemoteDbError
			}

			data = resource_model.PhysicalModel{
				Kind:        spec.Kind,
				Name:        spec.Name,
				Description: spec.Description,
				Unit:        spec.Unit,
			}
			if err = tx.Create(&data).Error; err != nil {
				return err, codes.RemoteDbError
			}
		} else {
			err = error_utils.NewConflictAlreadyExistsError(spec.Name)
			return err, codes.ClientAlreadyExists
		}
	}

	tx.Commit()

	return nil, codes.Ok
}

func (modelApi *ResourceModelApi) convertPhysicalModels(tctx *logger.TraceContext, physicalModels []resource_model.PhysicalModel) []*resource_api_grpc_pb.PhysicalModel {
	pbPhysicalModels := make([]*resource_api_grpc_pb.PhysicalModel, len(physicalModels))
	for i, physicalModel := range physicalModels {
		updatedAt, err := ptypes.TimestampProto(physicalModel.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", physicalModel.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(physicalModel.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", physicalModel.Model.CreatedAt)
			continue
		}

		pbPhysicalModels[i] = &resource_api_grpc_pb.PhysicalModel{
			Name:      physicalModel.Name,
			Kind:      physicalModel.Kind,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbPhysicalModels
}
