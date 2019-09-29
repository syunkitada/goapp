package base_app

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (app *BaseApp) SyncNodeByDb(tctx *logger.TraceContext, spec interface{}) (err error) {
	data := &base_spec.UpdateNode{
		Node: base_spec.Node{
			Name:         app.conf.Host,
			Kind:         app.appConf.Name,
			Role:         base_const.RoleMember,
			Status:       base_const.StatusEnabled,
			StatusReason: "Default",
			State:        base_const.StateUp,
			StateReason:  "SyncNode",
			Spec:         spec,
		},
	}

	err = app.dbApi.CreateOrUpdateNode(tctx, data)
	fmt.Println("SyncedNode", data)
	return
}

func (app *BaseApp) SyncNodeRole(tctx *logger.TraceContext) (role string, err error) {
	var nodes []base_db_model.Node
	nodes, err = app.dbApi.SyncNodeRole(tctx, app.appConf.Name)

	existsSelfNode := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != app.appConf.Name {
			continue
		}
		if node.Name == app.conf.Host && node.Status == base_const.StatusEnabled && node.State == base_const.StateUp {
			existsSelfNode = true
			role = node.Role
		}
		if node.Status == base_const.StatusEnabled && node.State == base_const.StateUp {
			if node.Role == base_const.RoleLeader {
				existsActiveLeader = true
			}
		}
	}

	if !existsSelfNode {
		err = fmt.Errorf("This node is not activated")
		return
	}

	if !existsActiveLeader {
		err = fmt.Errorf("Active Leader is not exists, after ReassignNode")
		return
	}
	return
}
