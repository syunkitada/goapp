package spec

import "github.com/syunkitada/goapp/pkg/base/base_index_model"

type PhysicalModel struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Unit        uint8
	Description string
	Spec        interface{}
}

type GetPhysicalModel struct {
	Name string `validate:"required"`
}

type GetPhysicalModelData struct {
	PhysicalModel PhysicalModel
}

type GetPhysicalModels struct{}

type GetPhysicalModelsData struct {
	PhysicalModels []PhysicalModel
}

type CreatePhysicalModel struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreatePhysicalModelData struct{}

type UpdatePhysicalModel struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdatePhysicalModelData struct{}

type DeletePhysicalModel struct {
	Name string `validate:"required"`
}

type DeletePhysicalModelData struct{}

type DeletePhysicalModels struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeletePhysicalModelsData struct{}

var PhysicalModelsTable = base_index_model.Table{
	Name:    "Models",
	Route:   "/Models",
	Kind:    "Table",
	DataKey: "PhysicalModels",
	Actions: []base_index_model.Action{
		base_index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "PhysicalModel",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text", Required: true,
					Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				base_index_model.Field{Name: "Kind", Kind: "select", Required: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "PhysicalModel",
			SelectKey: "Name",
		},
	},
	ColumnActions: []base_index_model.Action{
		base_index_model.Action{Name: "Detail", Icon: "Detail"},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name:           "Name",
			IsSearch:       true,
			Link:           "Datacenters/:Datacenter/Resources/Models/Detail/:0/View",
			LinkKey:      "resource",
			LinkSync:       false,
			LinkDataQueries: []string{"GetPhysicalModel"}},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "desc"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "Action", Kind: "Action"},
	},
}

var PhysicalModelsDetail = base_index_model.Tabs{
	Name:            "Models",
	Kind:            "RouteTabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Models",
	Route:           "/Datacenters/:Datacenter/Resources/Models/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	DataQueries: []string{
		"GetPhysicalModel",
		"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
	ExpectedDataKeys: []string{"PhysicalModel"},
	IsSync:           true,
	Tabs: []interface{}{
		base_index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "PhysicalModel",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text"},
				base_index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		base_index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "PhysicalModel",
			SubmitAction: "UpdatePhysicalModel",
			Icon:         "Update",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text", Required: true,
					Updatable: false,
					Min:       5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				base_index_model.Field{Name: "Kind", Kind: "select", Required: true,
					Updatable: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
}
