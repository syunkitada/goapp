package genpkg

import "github.com/syunkitada/goapp/pkg/base/base_model/index_model"

var DatacentersTable = index_model.Table{
	Name:    "Datacenters",
	Kind:    "Table",
	Route:   "",
	DataKey: "Datacenters",
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name:      "Name",
			IsSearch:  true,
			Link:      "Datacenters/:0/Resources/Resources",
			LinkParam: "Datacenter",
			LinkSync:  true,
			LinkGetQueries: []string{
				"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
		},
		index_model.TableColumn{Name: "Region", IsSearch: true},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
	SelectActions: []index_model.Action{
		index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Datacenter",
			SelectKey: "Name",
		},
	},
	Actions: []index_model.Action{
		index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "Datacenter",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Options: []string{
						"Private", "Share",
					}},
			},
		},
	},
}

var PhysicalResourcesTable = index_model.Table{
	Name:    "Resources",
	Route:   "PhysicalResources",
	Kind:    "Table",
	DataKey: "PhysicalResources",
	Actions: []index_model.Action{
		index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			Query:    "CreatePhysicalResource",
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
			Query:     "DeletePhysicalResources",
			DataKind:  "PhysicalResource",
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
			Link:           "Datacenters/:Datacenter/Resources/Resources/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       true,
			LinkGetQueries: []string{"GetPhysicalResource"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "Action", Kind: "Action"},
	},
}

var PhysicalResourcesDetail = index_model.Tabs{
	Name:            "Resources",
	Kind:            "RouteTabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Resources",
	Route:           "/Datacenters/:Datacenter/Resources/Resources/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	GetQueries: []string{
		"GetPhysicalResource",
		"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
	ExpectedDataKeys: []string{"PhysicalResource"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "PhysicalResource",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "PhysicalResource",
			SubmitAction: "UpdatePhysicalResource",
			Icon:         "Update",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Updatable: false,
					Min:       5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Updatable: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
}

var RacksTable = index_model.Table{
	Name:    "Racks",
	Route:   "/Racks",
	Kind:    "Table",
	DataKey: "Racks",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Rack",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Datacenters/:Datacenter/Resources/Racks/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetRack"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var FloorsTable = index_model.Table{
	Name:    "Floors",
	Route:   "/Floors",
	Kind:    "Table",
	DataKey: "Floors",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Floor",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Datacenters/:Datacenter/Resources/Floors/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetFloor"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var PhysicalModelsTable = index_model.Table{
	Name:    "Models",
	Route:   "/Models",
	Kind:    "Table",
	DataKey: "PhysicalModels",
	Actions: []index_model.Action{
		index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "PhysicalModel",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
	SelectActions: []index_model.Action{
		index_model.Action{
			Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "PhysicalModel",
			SelectKey: "Name",
		},
	},
	ColumnActions: []index_model.Action{
		index_model.Action{Name: "Detail", Icon: "Detail"},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name:           "Name",
			IsSearch:       true,
			Link:           "Datacenters/:Datacenter/Resources/Models/Detail/:0/View",
			LinkParam:      "resource",
			LinkSync:       false,
			LinkGetQueries: []string{"GetPhysicalModel"}},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "desc"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "Action", Kind: "Action"},
	},
}

var PhysicalModelsDetail = index_model.Tabs{
	Name:            "Models",
	Kind:            "RouteTabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Models",
	Route:           "/Datacenters/:Datacenter/Resources/Models/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	GetQueries: []string{
		"GetPhysicalModel",
		"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
	ExpectedDataKeys: []string{"PhysicalModel"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "PhysicalModel",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "PhysicalModel",
			SubmitAction: "UpdatePhysicalModel",
			Icon:         "Update",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Updatable: false,
					Min:       5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Updatable: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
}
