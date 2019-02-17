package resource_cluster_api

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (srv *ResourceClusterApiServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNodeTask(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterApiServer) UpdateNodeTask(tctx *logger.TraceContext) error {
	req := &resource_cluster_api_grpc_pb.UpdateNodeRequest{
		Tctx: logger.NewAuthproxyTraceContext(tctx, nil),
		Node: &resource_cluster_api_grpc_pb.Node{
			Node: &resource_api_grpc_pb.Node{
				Name:         srv.conf.Default.Host,
				Kind:         resource_cluster_model.KindResourceClusterApi,
				Role:         resource_cluster_model.RoleMember,
				Status:       resource_cluster_model.StatusEnabled,
				StatusReason: "Default",
				State:        resource_cluster_model.StateUp,
				StateReason:  "UpdateNode",
			},
		},
	}

	rep := &resource_cluster_api_grpc_pb.UpdateNodeReply{Tctx: logger.NewAuthproxyTraceContext(tctx, nil)}
	srv.resourceClusterModelApi.UpdateNode(tctx, req, rep)
	if rep.Tctx.StatusCode != codes.Ok {
		return fmt.Errorf("Err=%v, StatusCode=%v", rep.Tctx.Err, rep.Tctx.StatusCode)
	}

	return nil
}
