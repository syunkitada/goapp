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

func (srv *ResourceApiServer) UpdateNodeTask() error {
	var err error
	req := &resource_api_grpc_pb.UpdateNodeRequest{
		Name:         srv.conf.Default.Name,
		Kind:         resource_model.KindResourceApi,
		Role:         resource_model.RoleMember,
		Status:       resource_model.StatusEnabled,
		StatusReason: "Always Enabled",
		State:        resource_model.StateUp,
		StateReason:  "UpdateNode",
	}

	rep, err := srv.resourceModelApi.UpdateNode(req)
	if err != nil {
		return nil
	}
	glog.Info(rep)
	glog.Info("UpdatedNodeTask")
	return nil
}
