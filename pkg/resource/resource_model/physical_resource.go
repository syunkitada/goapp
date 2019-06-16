package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const PhysicalResourceKind = "PhysicalResource"

type PhysicalResource struct {
	gorm.Model
	Datacenter    string `gorm:"not null;size:50;index;"`
	Rack          string `gorm:"not null;size:50;index;"`
	Cluster       string `gorm:"not null;size:50;index;"`  // 仮想リソースを扱う場合はClusterに紐図かせる
	Name          string `gorm:"not null;size:200;index;"` // Datacenter内でユニーク
	Kind          string `gorm:"not null;size:25;"`        // Server, Pdu, L2Switch, L3Switch, RootSwitch
	PhysicalModel string `gorm:"not null;size:200;"`
	RackPosition  uint8  `gorm:"not null;"`
	Status        string `gorm:"not null;size:25;"`
	StatusReason  string `gorm:"not null;size:50;"`
	PowerLinks    string `gorm:"not null;size:5000;"`
	NetLinks      string `gorm:"not null;size:5000;"`
	Spec          string `gorm:"not null;size:5000;"`
}

type PhysicalResourceSpec struct {
	Kind         string `validate:"required"`
	Name         string `validate:"required"`
	Datacenter   string `validate:"required"`
	Cluster      string
	Rack         string
	Model        string
	RackPosition uint8
	NetLinks     []string
	PowerLinks   []string
	Spec         string
}

var PhysicalResourceCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"CreatePhysicalResource": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: PhysicalResourceKind,
		Help:    "helptext",
	},
	"UpdatePhysicalResource": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: PhysicalResourceKind,
		Help:    "helptext",
	},
	"GetPhysicalResources": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     PhysicalResourceKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Datacenter"},
		FlagMap: map[string]index_model.Flag{
			"datacenter": index_model.Flag{
				Flag:     index_model.ArgRequired,
				FlagType: index_model.ArgTypeString,
				Help:     "datacenter",
			},
		},
	},
	"GetPhysicalResource": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: PhysicalResourceKind,
		Help:    "helptext",
	},
	"DeletePhysicalResource": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: PhysicalResourceKind,
		Help:    "helptext",
	},
}
