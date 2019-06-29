package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const NodeKind = "Node"

type Node struct {
	gorm.Model
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Role         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	State        string `gorm:"not null;size:25;"`
	StateReason  string `gorm:"not null;size:50;"`
}

type NodeSpec struct {
	Name         string `validate:"required"`
	Kind         string `validate:"required"`
	Role         string
	Status       string
	StatusReason string
	State        string
	StateReason  string
	Spec         interface{}
}

var NodeCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_node": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: NodeKind,
		Help:    "create node",
	},
	"update_node": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: NodeKind,
		Help:    "update node",
	},
	"get_nodes": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     NodeKind,
		Help:        "get nodes",
		TableHeader: []string{"Cluster", "Name", "Kind", "Role", "Status", "State"},
	},
	"get_node": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: NodeKind,
		Help:    "get node",
	},
	"delete_node": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: NodeKind,
		Help:    "delete node",
	},
}
