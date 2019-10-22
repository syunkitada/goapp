package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const FloorKind = "Floor"

type Floor struct {
	gorm.Model
	Datacenter string `gorm:"not null;size:50;index;"`
	Name       string `gorm:"not null;size:50;index;"` // Datacenter内でユニーク
	Kind       string `gorm:"not null;size:25;"`
	Zone       string `gorm:"not null;size:50;"`
	Floor      uint8  `gorm:"not null;"`
	Spec       string `gorm:"not null;size:1000;"`
}

type FloorSpec struct {
	Kind       string `validate:"required"`
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
	Zone       string `validate:"required"`
	Floor      uint8  `validate:"required"`
	Spec       interface{}
}

var FloorCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_floor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
	"update_floor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
	"get_floors": index_model.Cmd{
		Arg:     index_model.ArgOptional,
		ArgType: index_model.ArgTypeString,
		ArgKind: FloorKind,
		FlagMap: map[string]index_model.Flag{
			"datacenter": index_model.Flag{
				Flag:     index_model.ArgRequired,
				FlagType: index_model.ArgTypeString,
				Help:     "datacenter",
			},
		},
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Datacenter", "Zone", "Floor"},
	},
	"get_floor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
	"delete_floor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: FloorKind,
		Help:    "helptext",
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
			Link:           "Datacenters/:datacenter/Resources/Floors/Detail/:0/View",
			LinkParam:      "resource",
			LinkSync:       false,
			LinkGetQueries: []string{"get_floor"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
