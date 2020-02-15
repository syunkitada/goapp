package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
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

type GetNodeMetrics struct {
	Cluster   string `validate:"required"`
	Name      string `validate:"required"`
	FromTime  string
	UntilTime *time.Time
}

type GetNodeMetricsData struct {
	NodeMetrics Node
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

var NodesTable = base_index_model.Table{
	Name:        "Nodes",
	Route:       "/Nodes",
	Kind:        "Table",
	DataQueries: []string{"GetNodes"},
	DataKey:     "Nodes",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Clusters/:Cluster/Resources/Nodes/Detail/:0/View",
			LinkKey:      "Name",
			LinkSync:       false,
			LinkDataQueries: []string{"GetNode"},
		},
		base_index_model.TableColumn{
			Name: "ActiveServices", Kind: "Popover", Icon: "Success", Color: "Success", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "ActiveServicesData",
				Columns: NodeServicesTableColumns,
			},
		},
		base_index_model.TableColumn{
			Name: "CriticalServices", Kind: "Popover", Icon: "Critical", Color: "Critical", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "CriticalServicesData",
				Columns: NodeServicesTableColumns,
			},
		},
		base_index_model.TableColumn{
			Name: "DisabledServices", Kind: "Popover", Icon: "Silenced", Color: "Silenced", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "DisabledServicesData",
				Columns: NodeServicesTableColumns,
			},
		},
		base_index_model.TableColumn{
			Name: "SuccessEvents", Kind: "Popover", Icon: "Success", Color: "Success", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "SuccessEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		base_index_model.TableColumn{
			Name: "CriticalEvents", Kind: "Popover", Icon: "Critical", Color: "Critical", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "CriticalEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		base_index_model.TableColumn{
			Name: "WarningEvents", Kind: "Popover", Icon: "Warning", Color: "Warning", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "WarningEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		base_index_model.TableColumn{
			Name: "SilencedEvents", Kind: "Popover", Icon: "Silenced", Color: "Silenced", InactiveColor: "Default",
			View: base_index_model.Table{
				Kind:    "Table",
				DataKey: "SilencedEventsData",
				Columns: NodeEventsTableColumns,
			},
		},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
	},
}

var NodeServicesTableColumns = []base_index_model.TableColumn{
	base_index_model.TableColumn{Name: "Kind"},
	base_index_model.TableColumn{Name: "Role"},
	base_index_model.TableColumn{Name: "Status"},
	base_index_model.TableColumn{Name: "StatusReason"},
	base_index_model.TableColumn{Name: "State"},
	base_index_model.TableColumn{Name: "StateReason"},
}

var NodeEventsTableColumns = []base_index_model.TableColumn{
	base_index_model.TableColumn{Name: "Check"},
	base_index_model.TableColumn{Name: "Node"},
	base_index_model.TableColumn{
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
	base_index_model.TableColumn{Name: "Msg"},
	base_index_model.TableColumn{Name: "Time", Kind: "Time"},
}

var NodesDetail = base_index_model.Tabs{
	Name:             "Nodes",
	Kind:             "RouteTabs",
	RouteParamKey:    "Kind",
	RouteParamValue:  "Nodes",
	Route:            "/Clusters/:Cluster/Resources/Nodes/Detail/:Name/:Subkind",
	TabParam:         "Subkind",
	ExpectedDataKeys: []string{"Node"},
	IsSync:           true,
	Tabs: []interface{}{
		base_index_model.View{
			Name:        "View",
			Route:       "/View",
			Kind:        "View",
			DataQueries: []string{"GetNode"},
			DataKey:     "Node",
			PanelsGroups: []interface{}{
				map[string]interface{}{
					"Name": "Node Data",
					"Kind": "Cards",
					"Cards": []interface{}{
						map[string]interface{}{
							"Name": "Node",
							"Kind": "Fields",
							"Fields": []base_index_model.Field{
								base_index_model.Field{Name: "Name", Kind: "text"},
							},
						},
					},
				},
				map[string]interface{}{
					"Name": "Node Services",
					"Kind": "Cards",
					"Cards": []interface{}{
						base_index_model.Table{
							Name:           "ActiveServices",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "ActiveServicesData",
							Columns:        NodeServicesTableColumns,
						},
						base_index_model.Table{
							Name:           "CriticalServices",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "CriticalServicesData",
							Columns:        NodeServicesTableColumns,
						},
						base_index_model.Table{
							Name:           "DisableServices",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "DisabledServicesData",
							Columns:        NodeServicesTableColumns,
						},
					},
				},
				map[string]interface{}{
					"Name": "Node Events",
					"Kind": "Cards",
					"Cards": []interface{}{
						base_index_model.Table{
							Name:           "SuccessEvents",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "SuccessEventsData",
							Columns:        NodeEventsTableColumns,
						},
						base_index_model.Table{
							Name:           "CriticalEvents",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "CriticalEventsData",
							Columns:        NodeEventsTableColumns,
						},
						base_index_model.Table{
							Name:           "WarningEvents",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "WarningEventsData",
							Columns:        NodeEventsTableColumns,
						},
						base_index_model.Table{
							Name:           "SilencedEvents",
							Kind:           "Table",
							DisableToolbar: true,
							DisablePaging:  true,
							DataKey:        "SilencedEventsData",
							Columns:        NodeEventsTableColumns,
						},
					},
				},
			},
		},
		base_index_model.View{
			Name:        "Metrics",
			Route:       "/Metrics",
			Kind:        "View",
			DataQueries: []string{"GetNodeMetrics"},
			DataKey:     "NodeMetrics",
			PanelsGroups: []interface{}{
				map[string]interface{}{
					"Name":        "Inputs",
					"Kind":        "SearchForm",
					"DataQueries": []string{"GetNodeMetrics"},
					"Inputs": []interface{}{
						base_index_model.TableInputField{
							Name:     "FromTime",
							Type:     "Select",
							Data:     []string{"-6h", "-1d", "-3d"},
							Default:  "-6h",
							Multiple: false,
						},
						base_index_model.TableInputField{
							Name: "UntilTime",
							Type: "DateTime",
						},
					},
				},
				map[string]string{
					"Name":    "Metrics",
					"Kind":    "MetricsGroups",
					"DataKey": "MetricsGroups",
				},
			},
		},
	},
}
