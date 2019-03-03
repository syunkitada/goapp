package resource_model_api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) Create(tctx *logger.TraceContext, req *resource_api_grpc_pb.ActionRequest, rep *resource_api_grpc_pb.ActionReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var specs []resource_model.ResourceSpec
	if err = json.Unmarshal([]byte(req.Spec), &specs); err != nil {
		logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
		return
	}

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
		return
	}
	defer func() { err = db.Close() }()

	tx := db.Begin()
	defer tx.Rollback()

	switch req.Tctx.ActionName {
	case "CreatePhysicalResource":
		for _, spec := range specs {
			specBytes, err := json_utils.Marshal(&spec)
			if err != nil {
				logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
				return
			}

			switch spec.Kind {
			case resource_model.ResourceKindDatacenter:
				var rspec resource_model.DatacenterSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
					return
				}

				var data resource_model.Datacenter
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}

					data = resource_model.Datacenter{
						Kind:         rspec.Spec.Kind,
						Name:         rspec.Name,
						Region:       rspec.Spec.Region,
						DomainSuffix: rspec.Spec.DomainSuffix,
						Spec:         string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}
				} else {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientAlreadyExists, rspec.Name)
					return
				}

			case resource_model.ResourceKindCluster:
				var rspec resource_model.ClusterSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
					return
				}

				var data resource_model.Cluster
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}

					data = resource_model.Cluster{
						Kind:         rspec.Spec.Kind,
						Name:         rspec.Name,
						Datacenter:   rspec.Spec.Datacenter,
						DomainSuffix: rspec.Spec.DomainSuffix,
						Spec:         string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}
				} else {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientAlreadyExists, rspec.Name)
					return
				}

			case resource_model.ResourceKindFloor:
				var rspec resource_model.FloorSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
					return
				}

				var data resource_model.Floor
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
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
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}
				} else {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientAlreadyExists, rspec.Name)
					return
				}

			case resource_model.ResourceKindRack:
				var rspec resource_model.RackSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
					return
				}

				var data resource_model.Rack
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}

					data = resource_model.Rack{
						Kind:       rspec.Spec.Kind,
						Name:       rspec.Name,
						Datacenter: rspec.Spec.Datacenter,
						Floor:      rspec.Spec.Floor,
						Spec:       string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}
				} else {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientAlreadyExists, rspec.Name)
					return
				}

			case resource_model.ResourceKindPhysicalResource:
				var rspec resource_model.PhysicalResourceSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
					return
				}

				var data resource_model.PhysicalResource
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
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
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}
				} else {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientAlreadyExists, rspec.Name)
					return
				}

			case resource_model.ResourceKindPhysicalModel:
				var rspec resource_model.PhysicalModelSpec
				if err = json.Unmarshal(specBytes, &rspec); err != nil {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientBadRequest, err)
					return
				}

				var data resource_model.PhysicalModel
				if err = tx.Where("name = ?", rspec.Name).First(&data).Error; err != nil {
					if !gorm.IsRecordNotFoundError(err) {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}

					data = resource_model.PhysicalModel{
						Kind:        rspec.Spec.Kind,
						Name:        rspec.Name,
						Description: rspec.Spec.Description,
						Unit:        uint8(rspec.Spec.Unit),
						Spec:        string(specBytes),
					}
					if err = tx.Create(&data).Error; err != nil {
						logger.SetErrorTraceContext(rep.Tctx, codes.RemoteDbError, err)
						return
					}
				} else {
					logger.SetErrorTraceContext(rep.Tctx, codes.ClientAlreadyExists, rspec.Name)
					return
				}

			}

		}

	case "CreateVirtualResource":
		fmt.Println("TODO")

	}

	tx.Commit()

	// var compute resource_model.Compute
	// if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&compute).Error; err != nil {
	// 	if !gorm.IsRecordNotFoundError(err) {
	// 		rep.Tctx.Err = err.Error()
	// 		rep.Tctx.StatusCode = codes.RemoteDbError
	// 		return
	// 	}

	// 	compute = resource_model.Compute{
	// 		Cluster:      spec.Cluster,
	// 		Kind:         spec.Kind,
	// 		Name:         spec.Name,
	// 		Domain:       spec.Spec.Domain,
	// 		Spec:         req.Spec,
	// 		Status:       resource_model.StatusCreating,
	// 		StatusReason: fmt.Sprintf("CreateCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
	// 	}
	// 	if err = tx.Create(&compute).Error; err != nil {
	// 		rep.Tctx.Err = err.Error()
	// 		rep.Tctx.StatusCode = codes.RemoteDbError
	// 		return
	// 	}
	// } else {
	// 	rep.Tctx.Err = fmt.Sprintf("Already Exists: cluster=%v, name=%v", spec.Cluster, spec.Name)
	// 	rep.Tctx.StatusCode = codes.ClientAlreadyExists
	// 	return
	// }
	// tx.Commit()

	// computePb, err := modelApi.convertCompute(&compute)
	// if err != nil {
	// 	rep.Tctx.Err = err.Error()
	// 	rep.Tctx.StatusCode = codes.ServerInternalError
	// 	return
	// }

	// rep.Computes = []*resource_api_grpc_pb.Compute{computePb}
	// rep.Tctx.StatusCode = codes.Ok
	return
}
