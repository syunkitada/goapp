package resource_model_api

import (
	"encoding/json"
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/codes"
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
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.ClientBadRequest
	}

	fmt.Println(specs)
	fmt.Println("DEBUGlalala")
	for _, spec := range specs {
		switch spec.Kind {
		case resource_model.ResourceKindPhysicalModel:
			fmt.Println(string(spec.Spec))
		}
	}

	// var db *gorm.DB
	// if db, err = modelApi.open(tctx); err != nil {
	// 	rep.Tctx.Err = err.Error()
	// 	rep.Tctx.StatusCode = codes.RemoteDbError
	// 	return
	// }
	// defer func() { err = db.Close() }()

	// spec, statusCode, err := modelApi.validateComputeSpec(db, req.Spec)
	// if err != nil {
	// 	rep.Tctx.Err = err.Error()
	// 	rep.Tctx.StatusCode = statusCode
	// 	return
	// }

	// var compute resource_model.Compute
	// tx := db.Begin()
	// defer tx.Rollback()
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
