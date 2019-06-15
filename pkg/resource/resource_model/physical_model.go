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
