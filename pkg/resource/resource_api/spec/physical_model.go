package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

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
			LinkKey:      "resource",
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
