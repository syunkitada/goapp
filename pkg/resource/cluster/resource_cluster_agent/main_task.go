package resource_cluster_agent

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceClusterAgentServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNode(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterAgentServer) UpdateNode(tctx *logger.TraceContext) error {
	nodes := []resource_model.NodeSpec{
		resource_model.NodeSpec{
			Name:         srv.conf.Default.Host,
			Kind:         resource_model.KindResourceClusterAgent,
			Role:         resource_model.RoleMember,
			Status:       resource_model.StatusEnabled,
			StatusReason: "Default",
			State:        resource_model.StateUp,
			StateReason:  "UpdateNode",
		},
	}
	specs, err := json_utils.Marshal(nodes)
	if err != nil {
		return err
	}
	queries := []authproxy_model.Query{
		authproxy_model.Query{
			Kind: "update_node",
			StrParams: map[string]string{
				"Specs": string(specs),
			},
		},
	}
	if _, err := srv.resourceClusterApiClient.Action(
		logger.NewActionTraceContext(tctx, "system", "system", queries)); err != nil {
		return err
	}
	return nil
}
