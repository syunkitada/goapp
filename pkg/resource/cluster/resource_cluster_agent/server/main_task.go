package server

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncNode(tctx); err != nil {
		return
	}
	return
}

func (srv *Server) SyncNode(tctx *logger.TraceContext) (err error) {
	fmt.Println("DEBUG SyncNode")
	nodeSpec := resource_api_spec.NodeSpec{}
	queries := []base_client.Query{
		base_client.Query{
			Name: "SyncNode",
			Data: resource_api_spec.SyncNode{
				Node: base_spec.Node{
					Name:         srv.baseConf.Host,
					Kind:         srv.clusterConf.Agent.Name,
					Role:         base_const.RoleMember,
					Status:       base_const.StatusEnabled,
					StatusReason: "Default",
					State:        base_const.StateUp,
					StateReason:  "SyncNode",
					Spec:         nodeSpec,
				},
			},
		},
	}

	var syncNodeData *api_spec.SyncNodeData
	if syncNodeData, err = srv.apiClient.ResourceVirtualAdminSyncNode(tctx, queries); err != nil {
		return
	}
	fmt.Println("DEBUG SyncedNode", syncNodeData)
	return
}
