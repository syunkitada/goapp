package resource_cluster_model_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_utils"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (modelApi *ResourceClusterModelApi) Action(tctx *logger.TraceContext,
	req *authproxy_grpc_pb.ActionRequest, rep *authproxy_grpc_pb.ActionReply) {
	var err error
	var statusCode int64
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	response := authproxy_model.ActionResponse{
		Tctx: *req.Tctx,
	}

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		authproxy_utils.MergeResponse(rep, &response, nil, err, codes.RemoteDbError)
		return
	}
	defer modelApi.close(tctx, db)

	data := map[string]interface{}{}
	tmpStatusCode := codes.Unknown
	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "update_node":
			tmpStatusCode, err = modelApi.UpdateNode(tctx, db, query, data)
		case "get_computes":
			tmpStatusCode, err = modelApi.GetComputes(tctx, db, req, query, data)
		case "create_compute":
			tmpStatusCode, err = modelApi.CreateCompute(tctx, db, req, query)
		}

		if err != nil {
			break
		}

		if tmpStatusCode > statusCode {
			statusCode = tmpStatusCode
		}
	}

	authproxy_utils.MergeResponse(rep, &response, data, err, statusCode)
}
