package resource_authproxy

import (
	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

func (resource *Resource) getIndex() interface{} {
	return index_model.Panels{
		Name:      "Root",
		Kind:      "RoutePanels",
		SyncDelay: 20000,
		Panels: []interface{}{
			index_model.Table{
				Name:    "Datacenters",
				Kind:    "Table",
				Route:   "",
				Subname: "datacenter",
				DataKey: "Datacenters",
				Columns: []index_model.TableColumn{
					index_model.TableColumn{
						Name:      "Name",
						IsSearch:  true,
						Link:      "Datacenters/:0/Resources/Resources",
						LinkParam: "datacenter",
						LinkSync:  true,
						LinkGetQueries: []string{
							"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
					},
					index_model.TableColumn{Name: "Region", IsSearch: true},
					index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
					index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
				},
			},
			index_model.Tabs{
				Name:             "Resources",
				Kind:             "RouteTabs",
				Subname:          "kind",
				Route:            "/Datacenters/:datacenter/Resources/:kind",
				TabParam:         "kind",
				GetQueries:       []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
				ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
				IsSync:           true,
				Tabs: []interface{}{
					index_model.Table{
						Name:    "Resources",
						Route:   "PhysicalResources",
						Kind:    "Table",
						DataKey: "PhysicalResources",
						Actions: []index_model.Action{
							index_model.Action{
								Name: "Create", Icon: "Create", Kind: "Form",
								DataKind: "PhysicalResource",
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
								DataKind:  "PhysicalResource",
								SelectKey: "Name",
							},
						},
						ColumnActions: []index_model.Action{
							index_model.Action{Name: "Detail", Icon: "Detail"},
							index_model.Action{Name: "Update", Icon: "Update"},
						},
						Columns: []index_model.TableColumn{
							index_model.TableColumn{Name: "Name", IsSearch: true},
							index_model.TableColumn{Name: "Kind"},
							index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "Action", Kind: "Action"},
						},
					},
					gin.H{
						"Name":    "Racks",
						"Route":   "/Racks",
						"Kind":    "Table",
						"DataKey": "Racks",
						"SelectActions": []interface{}{
							gin.H{"Name": "Delete", "Icon": "Delete",
								"Kind":      "Form",
								"DataKind":  "Rack",
								"SelectKey": "Name",
							},
						},
						"Columns": []interface{}{
							gin.H{"Name": "Name", "IsSearch": true},
							gin.H{"Name": "Kind"},
							gin.H{"Name": "UpdatedAt", "Kind": "Time"},
							gin.H{"Name": "CreatedAt", "Kind": "Time"},
						},
					},
					gin.H{
						"Name":    "Floors",
						"Route":   "/Floors",
						"Kind":    "Table",
						"DataKey": "Floors",
						"SelectActions": []interface{}{
							gin.H{"Name": "Delete", "Icon": "Delete",
								"Kind":      "Form",
								"DataKind":  "Floor",
								"SelectKey": "Name",
							},
						},
						"Columns": []interface{}{
							gin.H{"Name": "Name", "IsSearch": true},
							gin.H{"Name": "Kind"},
							gin.H{"Name": "UpdatedAt", "Kind": "Time"},
							gin.H{"Name": "CreatedAt", "Kind": "Time"},
						},
					},
					gin.H{
						"Name":    "Models",
						"Route":   "/Models",
						"Kind":    "Table",
						"DataKey": "PhysicalModels",
						"Actions": []interface{}{
							gin.H{
								"Name": "Create", "Icon": "Create", "Kind": "Form",
								"DataKind": "PhysicalModel",
								"Fields": []interface{}{
									gin.H{"Name": "Name", "Kind": "text", "Require": true,
										"Min": 5, "Max": 200, "RegExp": "^[0-9a-zA-Z]+$",
										"RegExpMsg": "Please enter alphanumeric characters."},
									gin.H{"Name": "Kind", "Kind": "select", "Require": true,
										"Options": []string{
											"Server", "Pdu", "RackSpineRouter",
											"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
										}},
								},
							},
						},
						"SelectActions": []interface{}{
							gin.H{"Name": "Delete", "Icon": "Delete",
								"Kind":      "Form",
								"DataKind":  "PhysicalModel",
								"SelectKey": "Name",
							},
						},
						"ColumnActions": []interface{}{
							gin.H{"Name": "Detail", "Icon": "Detail"},
							gin.H{
								"Name": "Update", "Icon": "Update", "Kind": "Form",
								"GetQueries": []string{"GetPhysicalModel"},
								"DataKind":   "PhysicalModel",
								"Fields": []interface{}{
									gin.H{"Name": "Name", "Kind": "text", "Require": true,
										"Min": 5, "Max": 200, "RegExp": "^[0-9a-zA-Z]+$",
										"RegExpMsg": "Please enter alphanumeric characters."},
									gin.H{"Name": "Kind", "Kind": "select", "Require": true,
										"Options": []string{
											"Server", "Pdu", "RackSpineRouter",
											"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
										}},
								}},
						},
						"Columns": []interface{}{
							gin.H{"Name": "Name", "IsSearch": true,
								"Link":           "Datacenters/:datacenter/Resources/Models/Detail/:0/View",
								"LinkParam":      "resource",
								"LinkSync":       false,
								"LinkGetQueries": []string{"GetPhysicalModel"}},
							gin.H{"Name": "Kind"},
							gin.H{"Name": "UpdatedAt", "Kind": "Time", "Sort": "desc"},
							gin.H{"Name": "CreatedAt", "Kind": "Time"},
							gin.H{"Name": "Action", "Kind": "Action"},
						},
					},
				}, // Tabs
			},
			gin.H{
				"Name":     "Resource",
				"Subname":  "resource",
				"Route":    "/Datacenters/:datacenter/Resources/:kind/Detail/:resource/:subkind",
				"TabParam": "subkind",
				"Kind":     "RouteTabs",
				"GetQueries": []string{
					"GetPhysicalModel",
					"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
				"ExpectedDataKeys": []string{"PhysicalModel"},
				"IsSync":           true,
				"Tabs": []interface{}{
					gin.H{
						"Name":    "View",
						"Route":   "/View",
						"Kind":    "View",
						"DataKey": "PhysicalModel",
						"Fields": []interface{}{
							gin.H{"Name": "Name", "Kind": "text"},
							gin.H{"Name": "Kind", "Kind": "select"},
						},
					},
					gin.H{
						"Name":         "Edit",
						"Route":        "/Edit",
						"Kind":         "Form",
						"DataKey":      "PhysicalModel",
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
	}
}
