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

func (modelApi *ResourceModelApi) GetImage(tctx *logger.TraceContext,
	db *gorm.DB, query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.VirtualActionReply) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("resource is None")
	}

	var image resource_model.Image
	if err = db.Where(&resource_model.Image{
		Name: resource,
	}).First(&image).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Image = modelApi.convertImage(tctx, &image)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetImages(tctx *logger.TraceContext,
	db *gorm.DB, query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.VirtualActionReply) (int64, error) {
	var err error
	cluster, ok := query.StrParams["cluster"]
	if !ok || cluster == "" {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("cluster is None")
	}

	var images []resource_model.Image
	if err = db.Where("cluster = ?", cluster).Find(&images).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Images = modelApi.convertImages(tctx, images)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateImage(tctx *logger.TraceContext,
	db *gorm.DB, query *resource_api_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.ImageSpecData
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.Image
		if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Image{
				Kind:    spec.Kind,
				Name:    spec.Name,
				Cluster: spec.Cluster,
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

func (modelApi *ResourceModelApi) UpdateImage(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.ImageSpecData
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
		physicalModel := &resource_model.Image{
			Kind:    spec.Kind,
			Cluster: spec.Cluster,
		}
		if err = tx.Model(physicalModel).Where("name = ?", spec.Name).Updates(physicalModel).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteImage(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (int64, error) {
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

		if err = tx.Delete(&resource_model.Image{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) convertImage(tctx *logger.TraceContext,
	image *resource_model.Image) *resource_api_grpc_pb.Image {
	updatedAt, err := ptypes.TimestampProto(image.Model.UpdatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", image.Model.UpdatedAt)
	}
	createdAt, err := ptypes.TimestampProto(image.Model.CreatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", image.Model.CreatedAt)
	}

	return &resource_api_grpc_pb.Image{
		Name:      image.Name,
		Kind:      image.Kind,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
}

func (modelApi *ResourceModelApi) convertImages(tctx *logger.TraceContext,
	images []resource_model.Image) []*resource_api_grpc_pb.Image {
	pbImages := make([]*resource_api_grpc_pb.Image, len(images))
	for i, image := range images {
		pbImages[i] = modelApi.convertImage(tctx, &image)
	}

	return pbImages
}
