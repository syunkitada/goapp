package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const PhysicalModelKind = "PhysicalModel"

type PhysicalModel struct {
	gorm.Model
	Kind        string `gorm:"not null;size:25;"`
	Name        string `gorm:"not null;size:200;index;"`
	Description string `gorm:"not null;size:200;"`
	Unit        uint8  `gorm:"not null;"`
	Spec        string `gorm:"not null;size:5000;"`
}

type PhysicalModelSpec struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Unit        uint8
	Description string
	Spec        interface{}
}

var PhysicalModelCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"CreatePhysicalModel": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: PhysicalModelKind,
		Help:    "helptext",
	},
	"UpdatePhysicalModel": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: PhysicalModelKind,
		Help:    "helptext",
	},
	"GetPhysicalModels": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     PhysicalModelKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind"},
	},
	"GetPhysicalModel": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: PhysicalModelKind,
		Help:    "helptext",
	},
	"DeletePhysicalModel": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: PhysicalModelKind,
		Help:    "helptext",
	},
}

var ResourceModels = index_model.Table{
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
			Link:           "Datacenters/:datacenter/Resources/Models/Detail/:0/View",
			LinkParam:      "resource",
			LinkSync:       false,
			LinkGetQueries: []string{"GetPhysicalModel"}},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "desc"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "Action", Kind: "Action"},
	},
}
