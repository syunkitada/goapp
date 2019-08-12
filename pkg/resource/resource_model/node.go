package resource_model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const NodeKind = "Node"

type Node struct {
	gorm.Model
	Name           string `gorm:"not null;size:255;"`
	Kind           string `gorm:"not null;size:25;"`
	Role           string `gorm:"not null;size:25;"`
	Status         string `gorm:"not null;size:25;"`
	StatusReason   string `gorm:"not null;size:50;"`
	State          string `gorm:"not null;size:25;"`
	StateReason    string `gorm:"not null;size:50;"`
	ClusterName    string `gorm:"not null;size:50;"`
	Labels         string `gorm:"not null;size:500;"`
	ResourceLabels string `gorm:"not null;size:500;"`
	Spec           string `gorm:"not null;size:5000;"`
	Weight         int    `gorm:"-"`
}

type NodeSpec struct {
	Name           string `validate:"required"`
	Kind           string `validate:"required"`
	Role           string
	Status         string
	StatusReason   string
	State          string
	StateReason    string
	Labels         []string
	ResourceLabels []string
	NumaNodes      []NumaNodeSpec
	Storages       []StorageSpec
	Spec           interface{}
}

type NumaNodeSpec struct {
	Id              int
	AvailableCpus   int
	UsedCpus        int
	AvailableMemory int
	UsedMemory      int
}

type StorageSpec struct {
	Kind               string
	Path               string
	AvailableGb        int
	UsedGb             int
	AvailableNumaNodes []int
}

type UpdateNodeResponse struct {
	Tctx authproxy_grpc_pb.TraceContext
	Data NodeTask
}

type NodeTask struct {
	ComputeAssignments []ComputeAssignmentEx
}

type AssignmentReportMap struct {
	ComputeAssignmentReports []AssignmentReport
}

type AssignmentReport struct {
	ID           uint
	Status       string
	StatusReason string
	State        string
	StateReason  string
	UpdatedAt    time.Time
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
		Arg:     index_model.ArgOptional,
		ArgType: index_model.ArgTypeString,
		ArgKind: NodeKind,
		FlagMap: map[string]index_model.Flag{
			"c,cluster": index_model.Flag{
				Flag:     index_model.ArgOptional,
				FlagType: index_model.ArgTypeString,
				Help:     "cluster",
			},
		},
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
