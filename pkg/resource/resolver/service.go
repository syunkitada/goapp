package resolver

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
	"github.com/syunkitada/goapp/pkg/resource/spec/genpkg"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	switch input.Name {
	case "ResourcePhysical":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourcePhysicalCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourcePhysicalAdmin":
		fmt.Println("DEBUG adminalalalal")
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourcePhysicalAdminCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtual":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourceVirtualCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtualAdmin":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourceVirtualAdminCmdMap,
			},
		}
		code = base_const.CodeOk
	default:
		code = base_const.CodeClientNotFound
	}

	return
}

func (resolver *Resolver) GetServiceDashboardIndex(tctx *logger.TraceContext,
	input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	// TODO
	switch input.Name {
	case "ResourcePhysicalAdmin":
		var datacenters []spec.Datacenter
		datacenters, err = resolver.dbApi.GetDatacenters(tctx, &spec.GetDatacenters{}, user)
		if err != nil {
			return
		}

		data = &base_spec.GetServiceDashboardIndexData{
			Data: map[string]interface{}{
				"Datacenters": datacenters,
			},
			Index: index_model.DashboardIndex{
				View: index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						spec.DatacentersTable,
						index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Datacenters/:Datacenter/Resources/:Kind",
							TabParam:         "Kind",
							GetQueries:       []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
							ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.PhysicalResourcesTable,
								spec.RacksTable,
								spec.FloorsTable,
								spec.PhysicalModelsTable,
							}, // Tabs
						},
						gin.H{
							"Name":      "Resource",
							"Subname":   "Name",
							"Route":     "/Datacenters/:Datacenter/Resources/:Kind/Detail/:Name/:Subkind",
							"Kind":      "RoutePanes",
							"PaneParam": "Kind",
							"Panes": []interface{}{
								spec.PhysicalModelsDetail,
								spec.PhysicalResourcesDetail,
							},
						},
					},
				},
			},
		}
		code = base_const.CodeOk
	case "ResourcePhysical":
		var datacenters []spec.Datacenter
		datacenters, err = resolver.dbApi.GetDatacenters(tctx, &spec.GetDatacenters{}, user)
		if err != nil {
			return
		}

		data = &base_spec.GetServiceDashboardIndexData{
			Data: map[string]interface{}{
				"Datacenters": datacenters,
			},
			Index: index_model.DashboardIndex{
				View: index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						spec.DatacentersTable,
						index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Datacenters/:Datacenter/Resources/:Kind",
							TabParam:         "Kind",
							GetQueries:       []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
							ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.PhysicalResourcesTable,
								spec.RacksTable,
								spec.FloorsTable,
								spec.PhysicalModelsTable,
							}, // Tabs
						},
						gin.H{
							"Name":      "Resource",
							"Subname":   "Name",
							"Route":     "/Datacenters/:Datacenter/Resources/:Kind/Detail/:Name/:Subkind",
							"Kind":      "RoutePanes",
							"PaneParam": "Kind",
							"Panes": []interface{}{
								spec.PhysicalModelsDetail,
								spec.PhysicalResourcesDetail,
							},
						},
					},
				},
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtual":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: index_model.DashboardIndex{
				View: index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						gin.H{
							"Name":  "ResourceVirtual HOGE",
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
		code = base_const.CodeOk
	default:
		code = base_const.CodeClientNotFound
	}

	return
}
