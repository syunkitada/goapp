package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const RegionKind = "Region"

type Region struct {
	gorm.Model
	Name string `gorm:"not null;size:50;unique_index;"`
	Kind string `gorm:"not null;size:25;"`
}

type RegionSpec struct {
	Name string `validate:"required"`
	Kind string `validate:"required"`
}

var RegionCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_region": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: RegionKind,
		Help:    "helptext",
	},
	"update_region": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: RegionKind,
		Help:    "helptext",
	},
	"get_regions": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     RegionKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Region"},
	},
	"get_region": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: RegionKind,
		Help:    "helptext",
	},
	"delete_region": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: RegionKind,
		Help:    "helptext",
	},
}
