package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec/genpkg"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	switch input.Name {
	case "ResourcePhysical":
		data = &base_spec.GetServiceIndexData{
			Index: base_index_model.Index{
				CmdMap: genpkg.ResourcePhysicalCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourcePhysicalAdmin":
		fmt.Println("DEBUG adminalalalal")
		data = &base_spec.GetServiceIndexData{
			Index: base_index_model.Index{
				CmdMap: genpkg.ResourcePhysicalAdminCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtual":
		data = &base_spec.GetServiceIndexData{
			Index: base_index_model.Index{
				CmdMap: genpkg.ResourceVirtualCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtualAdmin":
		data = &base_spec.GetServiceIndexData{
			Index: base_index_model.Index{
				CmdMap: genpkg.ResourceVirtualAdminCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceMonitor":
		data = &base_spec.GetServiceIndexData{
			Index: base_index_model.Index{
				CmdMap: genpkg.ResourceMonitorCmdMap,
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
			Index: base_index_model.DashboardIndex{
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						spec.DatacentersTable,
						base_index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Datacenters/:Datacenter/Resources/:Kind",
							TabParam:         "Kind",
							DataQueries:      []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
							ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.PhysicalResourcesTable,
								spec.RacksTable,
								spec.FloorsTable,
								spec.PhysicalModelsTable,
							}, // Tabs
						},
						map[string]interface{}{
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
			Index: base_index_model.DashboardIndex{
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						spec.DatacentersTable,
						base_index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Datacenters/:Datacenter/Resources/:Kind",
							TabParam:         "Kind",
							DataQueries:      []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
							ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.PhysicalResourcesTable,
								spec.RacksTable,
								spec.FloorsTable,
								spec.PhysicalModelsTable,
							}, // Tabs
						},
						map[string]interface{}{
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
			Index: base_index_model.DashboardIndex{
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						spec.RegionsTable,
						base_index_model.Tabs{
							Name:             "RegionResources",
							Kind:             "RouteTabs",
							Subname:          "RegionKind",
							Route:            "/Regions/:Region/RegionResources/:RegionKind",
							TabParam:         "RegionKind",
							ExpectedDataKeys: []string{"Clusters"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.VirtualAdminClustersTable,
								spec.RegionServicesTable,
								spec.ImagesTable,
							},
						},
						base_index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "ClusterKind",
							Route:            "/Regions/:Region/RegionResources/:RegionKind/:Cluster/Resources/:ClusterKind",
							TabParam:         "ClusterKind",
							ExpectedDataKeys: []string{"Clusters"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.ComputesTable,
							},
						},
						map[string]interface{}{
							"Name":      "Resource",
							"Subname":   "Name",
							"Route":     "/Regions/:Region/RegionResources/:RegionKind/:Cluster/Resources/:ClusterKind/:Name/:Subkind",
							"Kind":      "RoutePanes",
							"PaneParam": "ClusterKind",
							"Panes": []interface{}{
								spec.ComputesDetail,
							},
						},
					},
				},
			},
		}
		code = base_const.CodeOk

	case "ResourceVirtual":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: base_index_model.DashboardIndex{
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						map[string]interface{}{
							"Name":  "ResourceVirtual HOGE",
							"Kind":  "Msg",
							"Route": "",
						},
						map[string]interface{}{
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
			Index: base_index_model.DashboardIndex{
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						base_index_model.Table{
							Name:    "Clusters",
							Kind:    "Table",
							Route:   "",
							DataKey: "Clusters",
							Columns: []base_index_model.TableColumn{
								base_index_model.TableColumn{
									Name:        "Name",
									IsSearch:    true,
									Link:        "Clusters/:Cluster/Resources/Nodes",
									LinkParam:   "Cluster",
									LinkKey:     "Name",
									DataQueries: []string{"GetNodes"},
								},
								base_index_model.TableColumn{Name: "Region", IsSearch: true},
								base_index_model.TableColumn{Name: "Datacenter", IsSearch: true},
								base_index_model.TableColumn{Name: "Criticals"},
								base_index_model.TableColumn{Name: "Warnings"},
								base_index_model.TableColumn{Name: "Nodes"},
								base_index_model.TableColumn{Name: "Instances"},
								base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
								base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							},
						},
						base_index_model.Tabs{
							Name:             "Resources",
							Kind:             "RouteTabs",
							Subname:          "Kind",
							Route:            "/Clusters/:Cluster/Resources/:Kind",
							TabParam:         "Kind",
							DataQueries:      []string{"GetEvents", "GetEventRules", "GetNodes", "GetLogParams", "GetLogs"},
							ExpectedDataKeys: []string{"Events", "EventSilenceRules", "Nodes", "Logs"},
							IsSync:           true,
							Tabs: []interface{}{
								spec.EventsTable,
								spec.EventRulesTable,
								spec.NodesTable,
								spec.StatisticsTable,
								spec.LogsTable,
							},
						},
						map[string]interface{}{
							"Name":      "Resource",
							"Subname":   "Name",
							"Route":     "/Clusters/:Cluster/Resources/:Kind/Detail/:Name/:Subkind",
							"Kind":      "RoutePanes",
							"PaneParam": "Kind",
							"Panes": []interface{}{
								spec.NodesDetail,
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
