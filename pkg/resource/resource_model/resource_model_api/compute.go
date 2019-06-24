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

func (modelApi *ResourceModelApi) GetCompute(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var compute resource_model.Compute
	if err = db.Where(&resource_model.Compute{
		Name: resource,
	}).First(&compute).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Compute"] = compute
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetComputes(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Computes"] = computes
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateCompute(tctx *logger.TraceContext, tx *gorm.DB,
	regionService *resource_model.RegionService, spec *resource_model.RegionServiceSpec, cluster *resource_model.Cluster) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	specCompute := spec.Compute
	specBytes, err := json_utils.Marshal(spec)
	if err != nil {
		return error_utils.NewInvalidDataError("spec", spec, "Failed Marshal")
	}

	for i := 0; i < specCompute.Replicas; i++ {
		name := fmt.Sprintf("%s.r%d.%s.%s", spec.Name, i, regionService.Project, cluster.DomainSuffix)
		var data resource_model.Compute
		if err = tx.Where("name = ?", name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return err
			}

			data = resource_model.Compute{
				Project:       regionService.Project,
				Kind:          specCompute.Kind,
				Name:          name,
				RegionService: regionService.Name,
				Region:        cluster.Region,
				Cluster:       cluster.Name,
				Image:         specCompute.Image,
				Vcpus:         specCompute.Vcpus,
				Memory:        specCompute.Memory,
				Disk:          specCompute.Disk,
				Spec:          string(specBytes),
				Status:        resource_model.StatusInitializing,
				StatusReason:  "CreateCompute",
			}
			if err = tx.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (modelApi *ResourceModelApi) UpdateCompute(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteCompute(tctx *logger.TraceContext, db *gorm.DB,
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

		if err = tx.Delete(&resource_model.Compute{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) InitializeCompute(tctx *logger.TraceContext, db *gorm.DB, compute *resource_model.Compute, clusterNetworksMap map[string][]resource_model.NetworkV4) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	fmt.Println("DEBUG InitializeCompute")
	var spec resource_model.RegionServiceSpec
	if err = json_utils.Unmarshal(compute.Spec, &spec); err != nil {
		return
	}

	networkV4s, ok := clusterNetworksMap[compute.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("network")
		return
	}
	fmt.Println("DEBUG compute:", compute.Cluster, spec.Network)

	tx := db.Begin()
	defer tx.Rollback()
	switch spec.Network.Version {
	case 4:
		modelApi.AssignNetworkV4Port(tctx, tx, &spec.Network, networkV4s)
	}

	if err = modelApi.RegisterRecord(tctx, db, compute); err != nil {
		return
	}
	// Update Creating Initialized

	return
}
