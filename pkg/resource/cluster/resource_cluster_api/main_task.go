package resource_cluster_api

import (
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

func (server *ResourceClusterApiServer) MainTask() error {
	glog.Info("Run MainTask")
	server.UpdateNodeTask()

	return nil
}

func (server *ResourceClusterApiServer) UpdateNodeTask() error {
	request := resource_cluster_api_grpc_pb.UpdateNodeRequest{
		Name:         server.Conf.Default.Name,
		Kind:         resource_cluster_model.KindResourceApi,
		Role:         resource_cluster_model.RoleMember,
		Enable:       resource_cluster_model.StatusEnabled,
		EnableReason: "Always Enabled by UpdateNode",
		Status:       resource_cluster_model.StatusActive,
		StatusReason: "UpdateNode",
	}
	server.resourceClusterApiClient.UpdateNode(&request)

	glog.Info("UpdatedNode")
	return nil
}
