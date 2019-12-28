package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

type Node struct {
	Name                 string
	DisabledServices     int
	ActiveServices       int
	CriticalServices     int
	DisabledServicesData []base_spec.NodeService
	ActiveServicesData   []base_spec.NodeService
	CriticalServicesData []base_spec.NodeService
	SuccessEvents        int
	CriticalEvents       int
	WarningEvents        int
	SilencedEvents       int
	SuccessEventsData    []Event
	CriticalEventsData   []Event
	WarningEventsData    []Event
	SilencedEventsData   []Event
	MetricsGroups        []MetricsGroup
	Labels               string
	UpdatedAt            time.Time
}

type GetNodes struct {
	Cluster string `validate:"required"`
}

type GetNodesData struct {
	Nodes []Node
}

type GetNode struct {
	Cluster string `validate:"required"`
	Name    string `validate:"required"`
}

type GetNodeData struct {
	Node Node
}

type MetricsGroup struct {
	Name    string
	Metrics []Metric
}

type Metric struct {
	Name   string
	Keys   []string
	Values []map[string]interface{}
}

var NodesTable = index_model.Table{
	Name:    "Nodes",
	Route:   "/Nodes",
	Kind:    "Table",
	DataKey: "Nodes",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Clusters/:Cluster/Resources/Nodes/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetNode"},
		},
		index_model.TableColumn{
			Name: "ActiveServices", Kind: "Popover", Icon: "Success", Color: "Success", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "ActiveServicesData",
				Columns: NodeServicesTableColumns,
			},
		},
		index_model.TableColumn{
			Name: "CriticalServices", Kind: "Popover", Icon: "Critical", Color: "Critical", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "CriticalServicesData",
				Columns: NodeServicesTableColumns,
			},
		},
		index_model.TableColumn{
			Name: "DisabledServices", Kind: "Popover", Icon: "Silenced", Color: "Silenced", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "DisabledServicesData",
				Columns: NodeServicesTableColumns,
			},
		},
		index_model.TableColumn{
			Name: "SuccessEvents", Kind: "Popover", Icon: "Success", Color: "Success", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "SuccessEvents",
				Columns: NodeEventsTableColumns,
			},
		},
		index_model.TableColumn{
			Name: "CriticalEvents", Kind: "Popover", Icon: "Critical", Color: "Critical", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "CriticalEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		index_model.TableColumn{
			Name: "WarningEvents", Kind: "Popover", Icon: "Warning", Color: "Warning", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "WarningEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		index_model.TableColumn{
			Name: "SilencedEvents", Kind: "Popover", Icon: "Silenced", Color: "Silenced", InactiveColor: "Default",
			View: index_model.Table{
				Kind:    "Table",
				DataKey: "SilencedEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
	},
}

var NodeServicesTableColumns = []index_model.TableColumn{
	index_model.TableColumn{Name: "Kind"},
	index_model.TableColumn{Name: "Role"},
	index_model.TableColumn{Name: "Status"},
	index_model.TableColumn{Name: "StatusReason"},
	index_model.TableColumn{Name: "State"},
	index_model.TableColumn{Name: "StateReason"},
}

var NodeEventsTableColumns = []index_model.TableColumn{
	index_model.TableColumn{
		Name: "Check", IsSearch: true,
	},
	index_model.TableColumn{
		Name: "Node", IsSearch: true,
	},
	index_model.TableColumn{
		Name:           "Level",
		RowColoringMap: map[string]string{"Warning": "Warning", "Critical": "Critical"},
		FilterValues: []map[string]string{
			map[string]string{
				"Icon":  "Warning",
				"Value": "Warning",
			},
			map[string]string{
				"Icon":  "Critical",
				"Value": "Critical",
			},
		},
	},
	index_model.TableColumn{Name: "Msg"},
	index_model.TableColumn{Name: "Time", Kind: "Time"},
}

var NodesDetail = index_model.Tabs{
	Name:            "Nodes",
	Kind:            "RouteTabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Nodes",
	Route:           "/Clusters/:Cluster/Resources/Nodes/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	GetQueries: []string{
		"GetNode",
	},
	ExpectedDataKeys: []string{"Node"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "Node",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
			PanelsGroups: []interface{}{
				map[string]string{
					"Name":     "Metrics",
					"DataKey":  "MetricsGroups",
					"DataType": "MetricsGroups",
				},
			},
		},
	},
}
