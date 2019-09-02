package server

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) error {
	srv.UpdateService(tctx)

	if err := srv.UpdateNodeTask(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *Server) UpdateNodeTask(tctx *logger.TraceContext) error {
	fmt.Println("DEBUG UpdateNodeTask")
	return nil
	// nodes := []resource_model.NodeSpec{
	// 	resource_model.NodeSpec{
	// 		Name:         srv.conf.Default.Host,
	// 		Kind:         resource_model.KindResourceApi,
	// 		Role:         resource_model.RoleMember,
	// 		Status:       resource_model.StatusEnabled,
	// 		StatusReason: "Default",
	// 		State:        resource_model.StateUp,
	// 		StateReason:  "UpdateNode",
	// 	},
	// }
	// specs, err := json_utils.Marshal(nodes)
	// if err != nil {
	// 	return err
	// }
	// queries := []authproxy_model.Query{
	// 	authproxy_model.Query{
	// 		Kind: "update_node",
	// 		StrParams: map[string]string{
	// 			"Specs": string(specs),
	// 		},
	// 	},
	// }
	// _, err = srv.localVirtualAction(tctx, logger.NewActionTraceContext(tctx, "system", "system", queries))
	// if err != nil {
	// 	return err
	// }
	// return nil
}
