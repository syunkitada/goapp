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
	"CreateFloor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
	"UpdateFloor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
	"GetFloors": index_model.Cmd{
		Arg:     index_model.ArgOptional,
		ArgType: index_model.ArgTypeString,
		ArgKind: FloorKind,
		FlagMap: map[string]string{
			"datacenter": index_model.ArgRequired,
		},
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Datacenter", "Zone", "Floor"},
	},
	"GetFloor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
	"DeleteFloor": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: FloorKind,
		Help:    "helptext",
	},
}
