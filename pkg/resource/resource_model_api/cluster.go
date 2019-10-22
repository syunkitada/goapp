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

func (modelApi *ResourceModelApi) GetCluster(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var cluster resource_model.Cluster
	if err = db.Where(&resource_model.Cluster{
		Name: resource,
	}).First(&cluster).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Cluster"] = cluster
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetClusters(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	var clusters []resource_model.Cluster
	if err = db.Find(&clusters).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Clusters"] = clusters
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateCluster(tctx *logger.TraceContext, db *gorm.DB,
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

	var specs []resource_model.ClusterSpec
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

		var datacenter resource_model.Datacenter
		if err = tx.Where("name = ?", spec.Datacenter).First(&datacenter).Error; err != nil {
			return codes.RemoteDbError, err
		}

		var data resource_model.Cluster
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Cluster{
				Kind:         spec.Kind,
				Name:         spec.Name,
				Description:  spec.Description,
				Datacenter:   spec.Datacenter,
				Region:       datacenter.Region,
				DomainSuffix: spec.DomainSuffix,
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

func (modelApi *ResourceModelApi) UpdateCluster(tctx *logger.TraceContext, db *gorm.DB,
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

	var specs []resource_model.ClusterSpec
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
		datacenter := &resource_model.Cluster{
			Kind:         spec.Kind,
			Description:  spec.Description,
			Datacenter:   spec.Datacenter,
			DomainSuffix: spec.DomainSuffix,
		}
		if err = tx.Model(datacenter).Where("name = ?", spec.Name).Updates(datacenter).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteCluster(tctx *logger.TraceContext, db *gorm.DB,
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

		if err = tx.Delete(&resource_model.Cluster{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}
