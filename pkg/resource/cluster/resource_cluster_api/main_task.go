package resource_cluster_api

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceClusterApiServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNodeTask(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterApiServer) UpdateNodeTask(tctx *logger.TraceContext) error {
	nodes := []resource_model.NodeSpec{
		resource_model.NodeSpec{
			Name:         srv.conf.Default.Host,
			Kind:         resource_model.KindResourceClusterApi,
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

	_, err = srv.localAction(tctx,
		logger.NewActionTraceContext(tctx, "system", "system", queries))
	if err != nil {
		return err
	}
	return nil
}
