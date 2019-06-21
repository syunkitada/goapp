package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

type Compute struct {
	gorm.Model
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Cluster            string           `gorm:"not null;size:50;"`
	Name               string           `gorm:"not null;size:200;"`
	Description        string           `gorm:"not null;size:200;"`
	Kind               string           `gorm:"not null;size:25;"`
	Labels             string           `gorm:"not null;size:255;"`
	Status             string           `gorm:"not null;size:25;"`
	StatusReason       string           `gorm:"not null;size:50;"`
	Spec               string           `gorm:"not null;size:5000;"`
	Domain             string           `gorm:"not null;size:255;"`
	LinkSpec           string           `gorm:"not null;size:2500;"`
}

var ComputesTable = index_model.Table{
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
}

var ComputesDetail = index_model.Tabs{
	Name:            "Computes",
	Kind:            "RouteTabs",
	RouteParamKey:   "kind",
	RouteParamValue: "Computes",
	Route:           "/Clusters/:datacenter/Resources/Computes/Detail/:resource/:subkind",
	TabParam:        "subkind",
	GetQueries: []string{
		"GetCompute",
		"GetComputes", "GetImages"},
	ExpectedDataKeys: []string{"Compute"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "Compute",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "Compute",
			SubmitAction: "Update",
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
