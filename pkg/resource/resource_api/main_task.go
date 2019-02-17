package resource_api

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceApiServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNodeTask(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceApiServer) UpdateNodeTask(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() {
		logger.EndTrace(tctx, startTime, err, 0)
	}()

	req := &resource_api_grpc_pb.UpdateNodeRequest{
		Node: &resource_api_grpc_pb.Node{
			Name:         srv.Host,
			Kind:         resource_model.KindResourceApi,
			Role:         resource_model.RoleMember,
			Status:       resource_model.StatusEnabled,
			StatusReason: "Default",
			State:        resource_model.StateUp,
			StateReason:  "UpdateNode",
		},
	}

	rep := &resource_api_grpc_pb.UpdateNodeReply{Tctx: logger.NewAuthproxyTraceContext(tctx, nil)}
	srv.resourceModelApi.UpdateNode(tctx, req, rep)
	if rep.Tctx.StatusCode != codes.Ok {
		return fmt.Errorf("Err=%v, StatusCode=%v", rep.Tctx.Err, rep.Tctx.StatusCode)
	}
	return nil
}
