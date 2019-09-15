package resolver

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	cmdMap := map[string]index_model.Cmd{}
	cmdMaps := []map[string]index_model.Cmd{
		base_spec.UserCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	code = base_const.CodeOk
	data = &base_spec.GetServiceIndexData{
		Index: index_model.Index{
			CmdMap: cmdMap,
		},
	}

	return
}

func (resolver *Resolver) GetServiceDashboardIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceDashboardIndex) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	data = &base_spec.GetServiceDashboardIndexData{
		Index: index_model.DashboardIndex{
			View: index_model.Panels{
				Name: "Root",
				Kind: "RoutePanels",
				Panels: []interface{}{
					gin.H{
						"Name":  "Hoge",
						"Kind":  "Msg",
						"Route": "",
					},
					gin.H{
						"Name":  "Piyo",
						"Kind":  "Msg",
						"Route": "/Piyo",
					},
				},
			},
		},
	}

	return
}
