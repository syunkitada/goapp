package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	cmdMap := map[string]base_index_model.Cmd{}
	cmdMaps := []map[string]base_index_model.Cmd{
		base_spec.UserCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	code = base_const.CodeOk
	data = &base_spec.GetServiceIndexData{
		Index: base_index_model.Index{
			CmdMap: cmdMap,
		},
	}

	return
}

func (resolver *Resolver) GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	fmt.Println("DEBUG GetServiceDashboardIndex", input.Name)
	switch input.Name {
	case "Auth":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: base_index_model.DashboardIndex{
				DefaultRoute: map[string]interface{}{
					"Path": []string{"Home"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panels",
					Children: []interface{}{
						map[string]interface{}{
							"Name": "Home",
							"Kind": "Msg",
						},
						map[string]interface{}{
							"Name": "Piyo",
							"Kind": "Msg",
						},
					},
				},
			},
		}
	case "Home":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: base_index_model.DashboardIndex{
				DefaultRoute: map[string]interface{}{
					"Path": []string{"Home"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panels",
					Children: []interface{}{
						map[string]interface{}{
							"Name": "Home",
							"Kind": "Msg",
						},
						map[string]interface{}{
							"Name": "Piyo",
							"Kind": "Msg",
						},
					},
				},
			},
		}
	case "HomeProject":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: base_index_model.DashboardIndex{
				DefaultRoute: map[string]interface{}{
					"Path": []string{"Home"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panels",
					Children: []interface{}{
						map[string]interface{}{
							"Name": "Home",
							"Kind": "Msg",
						},
						map[string]interface{}{
							"Name": "Piyo",
							"Kind": "Msg",
						},
					},
				},
			},
		}
	}

	return
}
