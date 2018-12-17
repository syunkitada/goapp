package resource_cluster_agent

import (
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

func (srv *ResourceClusterAgentServer) MainTask(traceId string) error {
	glog.Info("Run MainTask")
	if err := srv.UpdateNode(); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterAgentServer) UpdateNode() error {
	req := resource_cluster_api_grpc_pb.UpdateNodeRequest{
		Name:          srv.conf.Default.Host,
		Kind:          resource_cluster_model.KindResourceClusterAgent,
		Role:          resource_cluster_model.RoleMember,
		Status:        resource_cluster_model.StatusEnabled,
		StatusReason:  "Default",
		State:         resource_cluster_model.StateUp,
		StateReason:   "UpdateNode",
		ComputeDriver: srv.conf.Resource.Node.Compute.Driver,
	}
	if _, err := srv.resourceClusterApiClient.UpdateNode(&req); err != nil {
		return err
	}

	glog.Info("UpdatedNode")
	return nil
}
