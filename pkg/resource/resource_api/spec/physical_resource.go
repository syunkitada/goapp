package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

type PhysicalResource struct {
	Kind          string `validate:"required"` // Server, Pdu, L2Switch, L3Switch, RootSwitch
	Name          string `validate:"required"` // Datacenter内でユニーク
	Datacenter    string `validate:"required"`
	Cluster       string // 仮想リソースを扱う場合はClusterに紐図かせる
	Rack          string
	PhysicalModel string
	RackPosition  uint8
	UpdatedAt     time.Time
	CreatedAt     time.Time
	Spec          string
}

type GetPhysicalResource struct {
	Datacenter string `validate:"required"`
	Name       string `validate:"required"`
}

type GetPhysicalResourceData struct {
	PhysicalResource PhysicalResource
}

type GetPhysicalResources struct {
	Datacenter string `validate:"required"`
}

type GetPhysicalResourcesData struct {
	PhysicalResources []PhysicalResource
}

type CreatePhysicalResource struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreatePhysicalResourceData struct{}

type UpdatePhysicalResource struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdatePhysicalResourceData struct{}

type DeletePhysicalResource struct {
	Datacenter string `validate:"required"`
	Name       string `validate:"required"`
}

type DeletePhysicalResourceData struct{}

type DeletePhysicalResources struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeletePhysicalResourcesData struct{}

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
			LinkKey:      "Name",
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
