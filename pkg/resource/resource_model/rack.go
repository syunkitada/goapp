package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const RackKind = "Rack"

type Rack struct {
	gorm.Model
	Datacenter string `gorm:"not null;size:50;index;"`
	Floor      string `gorm:"not null;size:50;index;"`
	Name       string `gorm:"not null;size:200;index;"` // Datacenter内でユニーク
	Kind       string `gorm:"not null;size:25;"`
	Unit       uint8  `gorm:"not null;"`
	Spec       string `gorm:"not null;size:1000;"`
}

type RackSpec struct {
	Kind       string `validate:"required"`
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
	Floor      string `validate:"required"`
	Unit       uint8
	Spec       interface{}
}

var RackCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"CreateRack": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: RackKind,
		Help:    "helptext",
	},
	"UpdateRack": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: RackKind,
		Help:    "helptext",
	},
	"GetRacks": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     RackKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Datacenter", "Floor", "Unit"},
		FlagMap: map[string]index_model.Flag{
			"datacenter": index_model.Flag{
				Flag:     index_model.ArgRequired,
				FlagType: index_model.ArgTypeString,
				Help:     "datacenter",
			},
		},
	},
	"GetRack": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: RackKind,
		Help:    "helptext",
		FlagMap: map[string]index_model.Flag{
			"datacenter": index_model.Flag{
				Flag:     index_model.ArgRequired,
				FlagType: index_model.ArgTypeString,
				Help:     "datacenter",
			},
		},
	},
	"DeleteRack": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: RackKind,
		Help:    "helptext",
	},
}
