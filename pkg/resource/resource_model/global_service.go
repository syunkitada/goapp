package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const GlobalServiceKind = "GlobalService"

type GlobalService struct {
	gorm.Model
	Domain string `gorm:"not null;"` // GSLB Domain
}

var GlobalServiceCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_global-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: GlobalServiceKind,
		Help:    "helptext",
	},
	"update_global-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: GlobalServiceKind,
		Help:    "helptext",
	},
	"get_global-services": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     GlobalServiceKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Region", "DomainSuffix"},
	},
	"get_global-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: GlobalServiceKind,
		Help:    "helptext",
	},
	"delete_global-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: GlobalServiceKind,
		Help:    "helptext",
	},
}
