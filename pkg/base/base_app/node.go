package base_app

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (app *BaseApp) SyncNodeServiceByDb(tctx *logger.TraceContext, spec interface{}) (err error) {
	data := &base_spec.UpdateNodeService{
		NodeService: base_spec.NodeService{
			Name:         app.conf.Host,
			Kind:         app.appConf.Name,
			Role:         base_const.RoleMember,
			Status:       base_const.StatusEnabled,
			StatusReason: "Default",
			State:        base_const.StateUp,
			StateReason:  "SyncNodeService",
			Spec:         spec,
		},
	}

	err = app.dbApi.CreateOrUpdateNodeService(tctx, data)
	return
}

func (app *BaseApp) SyncNodeServiceRole(tctx *logger.TraceContext) (role string, err error) {
	var nodes []base_db_model.NodeService
	nodes, err = app.dbApi.SyncNodeServiceRole(tctx, app.appConf.Name)

	existsSelfNodeService := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != app.appConf.Name {
			continue
		}
		if node.Name == app.conf.Host && node.Status == base_const.StatusEnabled && node.State == base_const.StateUp {
			existsSelfNodeService = true
			role = node.Role
		}
		if node.Status == base_const.StatusEnabled && node.State == base_const.StateUp {
			if node.Role == base_const.RoleLeader {
				existsActiveLeader = true
			}
		}
	}

	if !existsSelfNodeService {
		err = fmt.Errorf("This node is not activated")
		return
	}

	if !existsActiveLeader {
		err = fmt.Errorf("Active Leader is not exists, after ReassignNodeService")
		return
	}
	return
}

func (app *BaseApp) SyncNodeServiceState(tctx *logger.TraceContext) (err error) {
	err = app.dbApi.SyncNodeServiceState(tctx)
	return
}
