package resource_model_api

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetImage(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var image resource_model.Image
	if err = db.Where(&resource_model.Image{
		Name: resource,
	}).First(&image).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Image"] = image
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetImages(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	var images []resource_model.Image
	if err = db.Find(&images).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Images"] = images
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateImage(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.ImageSpec
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

		var specBytes []byte
		if specBytes, err = json_utils.Marshal(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		var data resource_model.Image
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Image{
				Kind:        spec.Kind,
				Name:        spec.Name,
				Description: spec.Description,
				Region:      spec.Region,
				Status:      resource_model.StatusActive,
				Spec:        string(specBytes),
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

func (modelApi *ResourceModelApi) UpdateImage(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.ImageSpec
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
		image := &resource_model.Image{
			Kind:        spec.Kind,
			Description: spec.Description,
		}
		if err = tx.Model(image).Where("name = ?", spec.Name).Updates(image).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteImage(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
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
