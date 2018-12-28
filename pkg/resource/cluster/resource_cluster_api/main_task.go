package resource_cluster_api

import (
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

func (srv *ResourceClusterApiServer) MainTask(tctx *logger.TraceContext) error {
	glog.Info("Run MainTask")
	srv.UpdateNodeTask()

	return nil
}

func (srv *ResourceClusterApiServer) UpdateNodeTask() error {
	req := &resource_cluster_api_grpc_pb.UpdateNodeRequest{
		Name:         srv.conf.Default.Host,
		Kind:         resource_cluster_model.KindResourceClusterApi,
		Role:         resource_cluster_model.RoleMember,
		Status:       resource_cluster_model.StatusEnabled,
		StatusReason: "Default",
		State:        resource_cluster_model.StateUp,
		StateReason:  "UpdateNode",
	}

	if _, err := srv.resourceClusterModelApi.UpdateNode(req); err != nil {
		return err
	}

	glog.Info("UpdatedNode")
	return nil
}
