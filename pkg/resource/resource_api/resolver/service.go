package resolver

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec/genpkg"
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

	case "ResourceVirtualAdmin":
		var regions []spec.Region
		regions, err = resolver.dbApi.GetRegions(tctx, &spec.GetRegions{}, user)
		if err != nil {
			return
		}

		data = &base_spec.GetServiceDashboardIndexData{
			Data: map[string]interface{}{
				"Regions": regions,
			},
			Index: index_model.DashboardIndex{
				View: index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						spec.RegionsTable,
						index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Regions/:Region/Resources/:Kind",
							TabParam:         "Kind",
							GetQueries:       []string{"GetRegionServices", "GetImages"},
							ExpectedDataKeys: []string{"RegionServices", "Images"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.RegionServicesTable,
								spec.ImagesTable,
							},
						},
						gin.H{
							"Name":      "Resource",
							"Subname":   "Name",
							"Route":     "/Regions/:Region/Resources/:Kind/Detail/:Name/:Subkind",
							"Kind":      "RoutePanes",
							"PaneParam": "Kind",
							"Panes": []interface{}{
								spec.ComputesDetail,
								spec.RegionServicesDetail,
								spec.ImagesDetail,
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

	case "ResourceMonitor":
		var clusters []spec.Cluster
		clusters, err = resolver.dbApi.GetClusters(tctx, &spec.GetClusters{}, user)
		if err != nil {
			return
		}

		data = &base_spec.GetServiceDashboardIndexData{
			Data: map[string]interface{}{
				"Clusters": clusters,
			},
			Index: index_model.DashboardIndex{
				View: index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						index_model.Table{
							Name:    "Clusters",
							Kind:    "Table",
							Route:   "",
							DataKey: "Clusters",
							Columns: []index_model.TableColumn{
								index_model.TableColumn{
									Name:           "Name",
									IsSearch:       true,
									Link:           "Clusters/:0/Resources/Resources",
									LinkParam:      "Cluster",
									LinkSync:       true,
									LinkGetQueries: []string{"GetNodes"},
								},
								index_model.TableColumn{Name: "Cluster", IsSearch: true},
								index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
								index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							},
						},
						index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Clusters/:Cluster/Resources/:Kind",
							TabParam:         "Kind",
							GetQueries:       []string{"GetNodes"},
							ExpectedDataKeys: []string{"Nodes"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.RegionServicesTable,
								spec.ImagesTable,
							},
						},
						gin.H{
							"Name":      "Resource",
							"Subname":   "Name",
							"Route":     "/Clusters/:Cluster/Resources/:Kind/Detail/:Name/:Subkind",
							"Kind":      "RoutePanes",
							"PaneParam": "Kind",
							"Panes": []interface{}{
								spec.ComputesDetail,
								spec.RegionServicesDetail,
								spec.ImagesDetail,
							},
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
