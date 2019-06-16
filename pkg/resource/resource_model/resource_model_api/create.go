package resource_model_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (modelApi *ResourceModelApi) Create(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (error, int64) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	return nil, codes.Ok
}
