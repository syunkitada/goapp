package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const ComputeKind = "Compute"

type Compute struct {
	gorm.Model
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Region             string           `gorm:"not null;size:50;"`
	Cluster            string           `gorm:"not null;size:50;"`
	RegionService      string           `gorm:"not null;size:63;"`
	Name               string           `gorm:"not null;size:255;unique_index"`
	Kind               string           `gorm:"not null;size:25;"`
	Labels             string           `gorm:"not null;size:255;"`
	Status             string           `gorm:"not null;size:25;"`
	StatusReason       string           `gorm:"not null;size:50;"`
	Spec               string           `gorm:"not null;size:5000;"`
	Project            string           `gorm:"not null;size:63;"`
	LinkSpec           string           `gorm:"not null;size:2500;"`
	Image              string           `gorm:"not null;size:255;"`
	Vcpus              uint             `gorm:"not null;"`
	Memory             uint             `gorm:"not null;"`
	Disk               uint             `gorm:"not null;"`
}

type ComputeAssignment struct {
	gorm.Model
	Compute      Compute `gorm:"foreignkey:ComputeID;association_foreignkey:Refer;"`
	ComputeID    uint    `gorm:"not null;"`
	Node         Node    `gorm:"foreignkey:NodeID;association_foreignkey:Refer;"`
	NodeID       uint    `gorm:"not null;"`
	Status       string  `gorm:"not null;size:25;"`
	StatusReason string  `gorm:"not null;size:50;"`
}

type ComputeAssignmentWithComputeAndNode struct {
	ComputeSpec  string
	ComputeName  string
	NodeID       uint
	NodeName     string
	Status       string
	StatusReason string
	Cost         int // calcurated by vcpu, memory, disk
}

type ComputeSpec struct {
	Kind           string `validate:"required"`
	Image          string `validate:"required"`
	SchedulePolicy SchedulePolicySpec
	NetworkPolicy  NetworkPolicySpec
	Name           string
	Cluster        string
	Ports          []PortSpec
	Vcpus          uint
	Memory         uint
	Disk           uint
}

type ActionResponse struct {
	Tctx authproxy_grpc_pb.TraceContext
	Data ResponseData
}

type ResponseData struct {
	Computes []Compute
}

var ComputeCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"get_computes": index_model.Cmd{
		Arg:     index_model.ArgOptional,
		ArgType: index_model.ArgTypeString,
		ArgKind: ComputeKind,
		FlagMap: map[string]index_model.Flag{
			"region": index_model.Flag{
				Flag:     index_model.ArgRequired,
				FlagType: index_model.ArgTypeString,
				Help:     "region",
			},
		},
		Help:        "get computes",
		TableHeader: []string{"Region", "Cluster", "Name", "Image", "Vcpus", "Memory", "Disk", "Status"},
	},
	"get_compute": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: ComputeKind,
		Help:    "get compute",
	},
}

var ComputesTable = index_model.Table{
	Name:    "Computes",
	Route:   "Computes",
	Kind:    "Table",
	DataKey: "Computes",
	Actions: []index_model.Action{
		index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "Compute",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text",
					Require: true, Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$"},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
				index_model.Field{Name: "Rack", Kind: "select", Require: true,
					DataKey: "Racks"},
				index_model.Field{Name: "Model", Kind: "select", Require: true,
					DataKey: "PhysicalModels"},
			},
		},
	},
	SelectActions: []index_model.Action{
		index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Compute",
			SelectKey: "Name",
		},
	},
	ColumnActions: []index_model.Action{
		index_model.Action{Name: "Detail", Icon: "Detail"},
		index_model.Action{Name: "Update", Icon: "Update"},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Clusters/:datacenter/Resources/Computes/Detail/:0/View",
			LinkParam:      "resource",
			LinkSync:       false,
			LinkGetQueries: []string{"GetCompute"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "Action", Kind: "Action"},
	},
}

var ComputesDetail = index_model.Tabs{
	Name:            "Computes",
	Kind:            "RouteTabs",
	RouteParamKey:   "kind",
	RouteParamValue: "Computes",
	Route:           "/Clusters/:datacenter/Resources/Computes/Detail/:resource/:subkind",
	TabParam:        "subkind",
	GetQueries: []string{
		"GetCompute",
		"GetComputes", "GetImages"},
	ExpectedDataKeys: []string{"Compute"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "Compute",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "Compute",
			SubmitAction: "update compute",
			Icon:         "Update",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Updatable: false,
					Min:       5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Updatable: true,
					Options: []string{
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
}
