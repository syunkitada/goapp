package resource_model_api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) Create(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (error, int64) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpec, ok := query.StrParams["Spec"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Spec")
		return error_utils.NewInvalidRequestError("NotFound Spec"), codes.ClientBadRequest
	}

	var specs []resource_model.ResourceSpec
	if err = json.Unmarshal([]byte(strSpec), &specs); err != nil {
		return err, codes.ClientBadRequest
	}

	tx := db.Begin()
	defer tx.Rollback()

	switch query.Kind {
	case "CreatePhysicalResource":
		for _, spec := range specs {
			specBytes, err := json_utils.Marshal(&spec)
			if err != nil {
				return err, codes.ClientBadRequest
			}

			switch spec.Kind {
			case resource_model.ResourceKindDatacenter:
				var rspec resource_model.DatacenterSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					return err, codes.ClientBadRequest
				}

				var data resource_model.Datacenter
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						return err, codes.RemoteDbError
					}

					data = resource_model.Datacenter{
						Kind:         rspec.Spec.Kind,
						Name:         rspec.Name,
						Region:       rspec.Spec.Region,
						DomainSuffix: rspec.Spec.DomainSuffix,
						Spec:         string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						return err, codes.RemoteDbError
					}
				} else {
					err = error_utils.NewConflictAlreadyExistsError(rspec.Name)
					return err, codes.ClientAlreadyExists
				}

			case resource_model.ResourceKindCluster:
				var rspec resource_model.ClusterSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					return err, codes.ClientBadRequest
				}

				var data resource_model.Cluster
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						return err, codes.RemoteDbError
					}

					data = resource_model.Cluster{
						Kind:         rspec.Spec.Kind,
						Name:         rspec.Name,
						Datacenter:   rspec.Spec.Datacenter,
						DomainSuffix: rspec.Spec.DomainSuffix,
						Spec:         string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						return err, codes.RemoteDbError
					}
				} else {
					err = error_utils.NewConflictAlreadyExistsError(rspec.Name)
					return err, codes.ClientAlreadyExists
				}

			case resource_model.ResourceKindFloor:
				var rspec resource_model.FloorSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					return err, codes.ClientBadRequest
				}

				var data resource_model.Floor
				if err = tx.Where("name = ? and datacenter = ?", rspec.Name, rspec.Spec.Datacenter).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						return err, codes.RemoteDbError
					}

					data = resource_model.Floor{
						Kind:       rspec.Spec.Kind,
						Name:       rspec.Name,
						Datacenter: rspec.Spec.Datacenter,
						Zone:       rspec.Spec.Zone,
						Floor:      rspec.Spec.Floor,
						Spec:       string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						return err, codes.RemoteDbError
					}
				} else {
					err = error_utils.NewConflictAlreadyExistsError(rspec.Name)
					return err, codes.ClientAlreadyExists
				}

			case resource_model.ResourceKindRack:
				var rspec resource_model.RackSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					return err, codes.ClientBadRequest
				}

				var data resource_model.Rack
				if err = tx.Where("name = ? and datacenter = ?", rspec.Name, rspec.Spec.Datacenter).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						return err, codes.RemoteDbError
					}

					data = resource_model.Rack{
						Kind:       rspec.Spec.Kind,
						Name:       rspec.Name,
						Datacenter: rspec.Spec.Datacenter,
						Floor:      rspec.Spec.Floor,
						Spec:       string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						return err, codes.RemoteDbError
					}
				} else {
					err = error_utils.NewConflictAlreadyExistsError(rspec.Name)
					return err, codes.ClientAlreadyExists
				}

			case resource_model.ResourceKindPhysicalResource:
				var rspec resource_model.PhysicalResourceSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					return err, codes.ClientBadRequest
				}

				var data resource_model.PhysicalResource
				if err = tx.Where("name = ? and datacenter = ?", rspec.Name, rspec.Spec.Datacenter).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						return err, codes.RemoteDbError
					}

					data = resource_model.PhysicalResource{
						Kind:          rspec.Spec.Kind,
						Name:          rspec.Name,
						Datacenter:    rspec.Spec.Datacenter,
						Cluster:       rspec.Spec.Cluster,
						Rack:          rspec.Spec.Rack,
						PhysicalModel: rspec.Spec.Model,
						RackPosition:  rspec.Spec.RackPosition,
						PowerLinks:    strings.Join(rspec.Spec.PowerLinks, ","),
						NetLinks:      strings.Join(rspec.Spec.NetLinks, ","),
						Spec:          string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						return err, codes.RemoteDbError
					}
				} else {
					err = error_utils.NewConflictAlreadyExistsError(rspec.Name)
					return err, codes.ClientAlreadyExists
				}

			case resource_model.ResourceKindPhysicalModel:
				var rspec resource_model.PhysicalModelSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					return err, codes.ClientBadRequest
				}

				var data resource_model.PhysicalModel
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						return err, codes.RemoteDbError
					}

					data = resource_model.PhysicalModel{
						Kind:        rspec.Spec.Kind,
						Name:        rspec.Name,
						Description: rspec.Spec.Description,
						Unit:        uint8(rspec.Spec.Unit),
						Spec:        string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						return err, codes.RemoteDbError
					}
				} else {
					err = error_utils.NewConflictAlreadyExistsError(rspec.Name)
					return err, codes.ClientAlreadyExists
				}

			}

		}

	case "CreateVirtualResource":
		fmt.Println("TODO")

	}

	tx.Commit()

	return nil, codes.Ok
}
