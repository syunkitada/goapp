package resource_model_api

import (
	"encoding/json"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) CreatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (error, int64) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Specs")
		return error_utils.NewInvalidRequestError("NotFound Specs"), codes.ClientBadRequest
	}

	var specs []resource_model.PhysicalResourceSpecData
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return err, codes.ClientBadRequest
	}
	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.PhysicalResource
		if err = tx.Where("name = ? and datacenter = ?", spec.Name, spec.Datacenter).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return err, codes.RemoteDbError
			}

			data = resource_model.PhysicalResource{
				Kind:          spec.Kind,
				Name:          spec.Name,
				Datacenter:    spec.Datacenter,
				Cluster:       spec.Cluster,
				Rack:          spec.Rack,
				PhysicalModel: spec.Model,
				RackPosition:  spec.RackPosition,
				PowerLinks:    strings.Join(spec.PowerLinks, ","),
				NetLinks:      strings.Join(spec.NetLinks, ","),
				Spec:          spec.Spec,
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
