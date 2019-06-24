package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const NetworkV4Kind = "NetworkV4"

type NetworkV4 struct {
	gorm.Model
	Cluster            string           `gorm:"not null;size:50;"`
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Name               string           `gorm:"not null;size:200;"`
	Description        string           `gorm:"not null;size:255;"`
	Kind               string           `gorm:"not null;size:25;"`
	Labels             string           `gorm:"not null;size:255;"`
	Status             string           `gorm:"not null;size:25;"`
	StatusReason       string           `gorm:"not null;size:50;"`
	Spec               string           `gorm:"not null;size:1000;"`
	Subnet             string           `gorm:"not null;"`
	StartIp            string           `gorm:"not null;"`
	EndIp              string           `gorm:"not null;"`
	Gateway            string           `gorm:"not null;"`
}

type Network struct {
	Id           uint
	Name         string
	AvailableIps []string
}

type NetworkV4Port struct {
	gorm.Model
	NetworkV4   NetworkV4 `gorm:"foreignkey:NetworkV4ID;association_foreignkey:Refer;"`
	NetworkV4ID uint      `gorm:"not null;"`
	Ip          string    `gorm:"not null;"`
	Mac         string    `gorm:"not null;"`
}

type NetworkV4Spec struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Cluster     string `validate:"required"`
	Subnet      string `validate:"required"`
	StartIp     string `validate:"required"`
	EndIp       string `validate:"required"`
	Gateway     string `validate:"required"`
}

var NetworkV4Cmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_network-v4": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: NetworkV4Kind,
		Help:    "create network-v4",
	},
	"update_network-v4": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: NetworkV4Kind,
		Help:    "update network-v4",
	},
	"get_network-v4s": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     NetworkV4Kind,
		Help:        "get network-v4s",
		TableHeader: []string{"Name", "Kind", "Cluster", "Subnet", "StartIp", "EndIp", "Status"},
	},
	"get_network-v4": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: NetworkV4Kind,
		Help:    "get network-v4",
	},
	"delete_network-v4": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: NetworkV4Kind,
		Help:    "delete network-v4",
	},
}
