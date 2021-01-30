package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec/genpkg"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	code = base_const.CodeClientNotFound
	return
}

func (resolver *Resolver) GetProjectServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	switch input.Name {
	case "ResourcePhysical":
		data = &base_spec.GetServiceIndexData{
			Index: base_index_model.Index{
				CmdMap: genpkg.ResourcePhysicalCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourcePhysicalAdmin":
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
	code = base_const.CodeClientNotFound
	return
}

func (resolver *Resolver) GetProjectServiceDashboardIndex(tctx *logger.TraceContext,
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
					Kind: "Panels",
					Children: []interface{}{
						spec.DatacentersTable,
						base_index_model.Tabs{
							Name:             "Resources",
							Kind:             "Tabs",
							Subname:          "Kind",
							TabParam:         "Kind",
							DataQueries:      []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
							ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
							IsSync:           true,
							Children: []interface{}{
								spec.PhysicalResourcesTable,
								spec.RacksTable,
								spec.FloorsTable,
								spec.PhysicalModelsTable,
							}, // Tabs
						},
						map[string]interface{}{
							"Name":      "Resource",
							"Subname":   "Name",
							"Kind":      "Panes",
							"PaneParam": "Kind",
							"Children": []interface{}{
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
					Kind: "Panels",
					Children: []interface{}{
						spec.DatacentersTable,
						base_index_model.Tabs{
							Name:             "Resources",
							Kind:             "Tabs",
							Subname:          "Kind",
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
							"Kind":      "Panes",
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
				DefaultRoute: map[string]interface{}{
					"Path": []string{"Regions"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panels",
					Children: []interface{}{
						spec.RegionsTable,
						base_index_model.Tabs{
							Name:             "RegionResources",
							SubNameParamKeys: []string{"Region"},
							Kind:             "Tabs",
							Children: []interface{}{
								spec.VirtualAdminClustersTable,
								spec.RegionServicesTable,
								spec.ImagesTable,
							},
						},
						base_index_model.Tabs{
							Name:             "Resources",
							SubNameParamKeys: []string{"Cluster"},
							Kind:             "Tabs",
							Subname:          "ClusterKind",
							TabParam:         "ClusterKind",
							IsSync:           true,
							Children: []interface{}{
								spec.ComputesTable,
							},
						},
						map[string]interface{}{
							"Name":             "Resource",
							"SubNameParamKeys": []string{"Name"},
							"Kind":             "Panes",
							"Children": []interface{}{
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
					Kind: "Panels",
					Children: []interface{}{
						map[string]interface{}{
							"Name": "ResourceVirtual HOGE",
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
				DefaultRoute: map[string]interface{}{
					"Path": []string{"Clusters"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panes",
					Children: []interface{}{
						map[string]interface{}{
							"Name": "Clusters",
							"Kind": "Pane",
							"Views": []interface{}{
								map[string]interface{}{
									"Kind":  "Title",
									"Title": "Clusters",
								},
								base_index_model.Table{
									Kind:               "Table",
									DataKey:            "Clusters",
									RowsPerPageOptions: []int{10, 20, 30},
									Columns: []base_index_model.TableColumn{
										base_index_model.TableColumn{
											Name:       "Name",
											IsSearch:   true,
											LinkPath:   []string{"Clusters", "Resources", "Events"},
											LinkKeyMap: map[string]string{"Cluster": "Name"},
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
							},
							"Children": []interface{}{
								base_index_model.Tabs{
									Name:     "Resources",
									Kind:     "Tabs",
									Subname:  "Kind",
									TabParam: "Kind",
									IsSync:   true,
									Children: []interface{}{
										spec.EventsTable,
										spec.EventRulesTable,
										spec.NodesTable,
										spec.LogsTable,
									},
								},
							},
						},
					},
				},
			},
		}
		// 				base_index_model.Tabs{
		// 					Name:     "Resources",
		// 					Kind:     "Tabs",
		// 					Subname:  "Kind",
		// 					TabParam: "Kind",
		// 					IsSync:   true,
		// 					Children: []interface{}{
		// 						spec.EventsTable,
		// 						spec.EventRulesTable,
		// 						spec.NodesTable,
		// 						spec.LogsTable,
		// 					},
		// 				},
		// 				map[string]interface{}{
		// 					"Name":      "Resource",
		// 					"Subname":   "Name",
		// 					"Kind":      "Panes",
		// 					"PaneParam": "Kind",
		// 					"Children": []interface{}{
		// 						spec.NodesDetail,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
		code = base_const.CodeOk

	default:
		code = base_const.CodeClientNotFound
	}

	return
}
