package resource_controller

import (
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (server *ResourceControllerServer) MainTask() error {
	glog.Info("Run MainTask")
	server.UpdateNode()

	return nil
}

func (server *ResourceControllerServer) UpdateNode() error {
	request := resource_api_grpc_pb.UpdateNodeRequest{
		Name:         server.Conf.Default.Name,
		Kind:         resource_model.KindResourceController,
		Role:         resource_model.RoleMember,
		Enable:       resource_model.StatusEnabled,
		EnableReason: "Always Enabled by UpdateNode",
		Status:       resource_model.StatusActive,
		StatusReason: "UpdateNode",
	}
	server.resourceApiClient.UpdateNode(&request)

	glog.Info("UpdatedNode")
	return nil
}
