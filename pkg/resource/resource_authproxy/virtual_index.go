package resource_authproxy

import (
	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

func (resource *Resource) getVirtualIndex() interface{} {
	return index_model.Index{
		SyncDelay: 20000,
		View: index_model.Panels{
			Name: "Root",
			Kind: "RoutePanels",
			Panels: []interface{}{
				index_model.Table{
					Name:    "Clusters",
					Kind:    "Table",
					Route:   "",
					Subname: "cluster",
					DataKey: "Clusters",
					Columns: []index_model.TableColumn{
						index_model.TableColumn{
							Name:      "Name",
							IsSearch:  true,
							Link:      "Clusters/:0/Resources/Computes",
							LinkParam: "cluster",
							LinkSync:  true,
							LinkGetQueries: []string{
								"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
						},
						index_model.TableColumn{Name: "Datacenter", IsSearch: true},
						index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
						index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
					},
				},
				index_model.Tabs{
					Name:             "Resources",
					Kind:             "RouteTabs",
					Subname:          "kind",
					Route:            "/Clusters/:datacenter/Resources/:kind",
					TabParam:         "kind",
					GetQueries:       []string{"GetComputes", "GetImages"},
					ExpectedDataKeys: []string{"Computes", "Images"},
					IsSync:           true,
					Tabs: []interface{}{
						index_model.Table{
							Name:    "Computes",
							Route:   "Computes",
							Kind:    "Table",
							DataKey: "Computes",
							Actions: []index_model.Action{
								index_model.Action{
									Name: "Create", Icon: "Create", Kind: "Form",
									DataKind: "Compute",
									Fields: []index_model.Field{
										index_model.Field{Name: "Name", Kind: "text",
											Require: true, Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$"},
										index_model.Field{Name: "Kind", Kind: "select", Require: true,
											Options: []string{
												"Server", "Pdu", "RackSpineRouter",
												"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
											}},
										index_model.Field{Name: "Rack", Kind: "select", Require: true,
											DataKey: "Racks"},
										index_model.Field{Name: "Model", Kind: "select", Require: true,
											DataKey: "PhysicalModels"},
									},
								},
							},
							SelectActions: []index_model.Action{
								index_model.Action{Name: "Delete", Icon: "Delete",
									Kind:      "Form",
									DataKind:  "Compute",
									SelectKey: "Name",
								},
							},
							ColumnActions: []index_model.Action{
								index_model.Action{Name: "Detail", Icon: "Detail"},
								index_model.Action{Name: "Update", Icon: "Update"},
							},
							Columns: []index_model.TableColumn{
								index_model.TableColumn{
									Name: "Name", IsSearch: true,
									Link:           "Clusters/:datacenter/Resources/Computes/Detail/:0/View",
									LinkParam:      "resource",
									LinkSync:       false,
									LinkGetQueries: []string{"GetCompute"},
								},
								index_model.TableColumn{Name: "Kind"},
								index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
								index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
								index_model.TableColumn{Name: "Action", Kind: "Action"},
							},
						},
						index_model.Table{
							Name:    "Images",
							Route:   "/Images",
							Kind:    "Table",
							DataKey: "Images",
							SelectActions: []index_model.Action{
								index_model.Action{
									Name:      "Delete",
									Icon:      "Delete",
									Kind:      "Form",
									DataKind:  "Image",
									SelectKey: "Name",
								},
							},
							Columns: []index_model.TableColumn{
								index_model.TableColumn{
									Name: "Name", IsSearch: true,
									Link:           "Clusters/:datacenter/Resources/Images/Detail/:0/View",
									LinkParam:      "resource",
									LinkSync:       false,
									LinkGetQueries: []string{"GetImage"},
								},
								index_model.TableColumn{Name: "Kind"},
								index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
								index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							},
						},
					}, // Tabs
				},
				gin.H{
					"Name":      "Resource",
					"Subname":   "resource",
					"Route":     "/Clusters/:datacenter/Resources/:kind/Detail/:resource/:subkind",
					"Kind":      "RoutePanes",
					"PaneParam": "kind",
					"Panes": []interface{}{
						gin.H{
							"Name":            "Computes",
							"Kind":            "RouteTabs",
							"RouteParamKey":   "kind",
							"RouteParamValue": "Computes",
							"Route":           "/Clusters/:datacenter/Resources/Computes/Detail/:resource/:subkind",
							"TabParam":        "subkind",
							"GetQueries": []string{
								"GetCompute",
								"GetComputes", "GetImages"},
							"ExpectedDataKeys": []string{"Compute"},
							"IsSync":           true,
							"Tabs": []interface{}{
								gin.H{
									"Name":    "View",
									"Route":   "/View",
									"Kind":    "View",
									"DataKey": "Compute",
									"Fields": []interface{}{
										gin.H{"Name": "Name", "Kind": "text"},
										gin.H{"Name": "Kind", "Kind": "select"},
									},
								},
								gin.H{
									"Name":         "Edit",
									"Route":        "/Edit",
									"Kind":         "Form",
									"DataKey":      "Compute",
									"SubmitAction": "Update",
									"Icon":         "Update",
									"Fields": []interface{}{
										gin.H{"Name": "Name", "Kind": "text", "Require": true,
											"Updatable": false,
											"Min":       5, "Max": 200, "RegExp": "^[0-9a-zA-Z]+$",
											"RegExpMsg": "Please enter alphanumeric characters."},
										gin.H{"Name": "Kind", "Kind": "select", "Require": true,
											"Updatable": true,
											"Options": []string{
												"Server", "Pdu", "RackSpineRouter",
												"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
											}},
									},
								},
							},
						},
						gin.H{
							"Name":            "Images",
							"Kind":            "RouteTabs",
							"RouteParamKey":   "kind",
							"RouteParamValue": "Images",
							"Route":           "/Clusters/:datacenter/Resources/Images/Detail/:resource/:subkind",
							"TabParam":        "subkind",
							"GetQueries": []string{
								"GetImage",
								"GetComputes", "GetImages"},
							"ExpectedDataKeys": []string{"Image"},
							"IsSync":           true,
							"Tabs": []interface{}{
								gin.H{
									"Name":    "View",
									"Route":   "/View",
									"Kind":    "View",
									"DataKey": "Image",
									"Fields": []interface{}{
										gin.H{"Name": "Name", "Kind": "text"},
										gin.H{"Name": "Kind", "Kind": "select"},
									},
								},
								gin.H{
									"Name":         "Edit",
									"Route":        "/Edit",
									"Kind":         "Form",
									"DataKey":      "Image",
									"SubmitAction": "Update",
									"Icon":         "Update",
									"Fields": []interface{}{
										gin.H{"Name": "Name", "Kind": "text", "Require": true,
											"Updatable": false,
											"Min":       5, "Max": 200, "RegExp": "^[0-9a-zA-Z]+$",
											"RegExpMsg": "Please enter alphanumeric characters."},
										gin.H{"Name": "Kind", "Kind": "select", "Require": true,
											"Updatable": true,
											"Options": []string{
												"Server", "Pdu", "RackSpineRouter",
												"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
											}},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
