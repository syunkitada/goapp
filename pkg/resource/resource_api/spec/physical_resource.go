package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
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

var PhysicalResourcesTable = base_index_model.Table{
	Name:    "Resources",
	Route:   "PhysicalResources",
	Kind:    "Table",
	DataKey: "PhysicalResources",
	Actions: []base_index_model.Action{
		base_index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			Query:    "CreatePhysicalResource",
			DataKind: "PhysicalResource",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text",
					Required: true, Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$"},
				base_index_model.Field{Name: "Kind", Kind: "select", Required: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
				base_index_model.Field{Name: "Rack", Kind: "select", Required: true,
					DataKey: "Racks"},
				base_index_model.Field{Name: "Model", Kind: "select", Required: true,
					DataKey: "PhysicalModels"},
			},
		},
	},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			Query:     "DeletePhysicalResources",
			DataKind:  "PhysicalResource",
			SelectKey: "Name",
		},
	},
	ColumnActions: []base_index_model.Action{
		base_index_model.Action{Name: "Detail", Icon: "Detail"},
		base_index_model.Action{Name: "Update", Icon: "Update"},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:            "Datacenters/:Datacenter/Resources/Resources/Detail/:0/View",
			LinkKeyMap:      map[string]string{"Name": "Name"},
			LinkSync:        true,
			LinkDataQueries: []string{"GetPhysicalResource"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "Action", Kind: "Action"},
	},
}

var PhysicalResourcesDetail = base_index_model.Tabs{
	Name:            "Resources",
	Kind:            "RouteTabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Resources",
	Route:           "/Datacenters/:Datacenter/Resources/Resources/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	DataQueries: []string{
		"GetPhysicalResource",
		"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
	ExpectedDataKeys: []string{"PhysicalResource"},
	IsSync:           true,
	Tabs: []interface{}{
		base_index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "PhysicalResource",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text"},
				base_index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		base_index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "PhysicalResource",
			SubmitAction: "UpdatePhysicalResource",
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
