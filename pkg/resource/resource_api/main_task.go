package resource_api

import (
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (server *ResourceApiServer) MainTask() error {
	glog.Info("Run MainTask")
	server.UpdateNodeTask()

	return nil
}

func (server *ResourceApiServer) UpdateNodeTask() error {
	request := resource_api_grpc_pb.UpdateNodeRequest{
		Name:         server.conf.Default.Name,
		Kind:         resource_model.KindResourceApi,
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
