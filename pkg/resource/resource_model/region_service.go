package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const RegionServiceKind = "RegionService"

const (
	SchedulePolicyAffinity     = "Affinity"
	SchedulePolicyAntiAffinity = "AntiAffinity"
)

type RegionService struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:63;"` // Vip Domain
	Project      string `gorm:"not null;size:63;"`
	Kind         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:100000;"`
}

type RegionServiceSpec struct {
	Name    string `validate:"required"`
	Region  string `validate:"required"`
	Kind    string `validate:"required"`
	Cluster string
	Compute ComputeSpec
}

type NetworkPolicySpec struct {
	Version        int
	Interfaces     int
	AssignPolicy   string
	StaticNetworks []string
}

type PortSpec struct {
	Version int
	Subnet  string
	Gateway string
	Ip      string
	Mac     string
}

type SchedulePolicySpec struct {
	Replicas                    int `validate:"required"`
	ClusterFilters              []string
	ClusterLabelFilters         []string
	NodeFilters                 []string
	NodeLabelFilters            []string
	NodeLabelSoftUntiAffinities []string
	NodeLabelSoftAffinities     []string
	NodeLabelHardUntiAffinities []string
	NodeLabelHardAffinities     []string
}

var RegionServiceCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_region-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: RegionServiceKind,
		Help:    "helptext",
	},
	"update_region-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: RegionServiceKind,
		Help:    "helptext",
	},
	"get_region-services": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     RegionServiceKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Region", "Status", "StatusReason"},
	},
	"get_region-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: RegionServiceKind,
		Help:    "helptext",
	},
	"delete_region-service": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: RegionServiceKind,
		Help:    "helptext",
	},
}
