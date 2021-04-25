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
	Cluster      string `validate:"required"`
	Name         string `validate:"required"`
	Target       string
	TimeDuration string
	UntilTime    string
}

type GetNodeMetricsData struct {
	NodeMetrics Node
}

var NodesTable = map[string]interface{}{
	"Name":        "Nodes",
	"DataQueries": []string{"GetNodes"},
	"Kind":        "Pane",
	"Views": []interface{}{
		base_index_model.Table{
			Kind:    "Table",
			DataKey: "Nodes",
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
					LinkPath:   []string{"Clusters", "Resources", "Nodes", "Node", "View"},
					LinkKeyMap: map[string]string{"Name": "Name"},
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
		},
	},
	"Children": []interface{}{
		NodesDetail,
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
	Name:            "Node",
	Kind:            "Tabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Nodes",
	TabParam:        "Subkind",
	Children: []interface{}{
		base_index_model.View{
			Name:        "View",
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
			Kind:        "View",
			DataQueries: []string{"GetNodeMetrics"},
			DataKey:     "NodeMetrics",
			PanelsGroups: []interface{}{
				map[string]interface{}{
					"Name":        "Display Time",
					"Kind":        "SearchForm",
					"DataQueries": []string{"GetNodeMetrics"},
					"Inputs": []interface{}{
						base_index_model.TableInputField{
							Name:     "TimeDuration",
							Type:     "Select",
							Data:     []string{"-3h", "-6h", "-1d", "-3d", "-7d"},
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
		base_index_model.View{
			Name:        "Proc Metrics",
			Kind:        "View",
			DataQueries: []string{"GetNodeMetrics"},
			DataKey:     "NodeMetrics",
			ViewParams: map[string]interface{}{
				"Target": "Proc",
			},
			PanelsGroups: []interface{}{
				map[string]interface{}{
					"Name":        "Display Time",
					"Kind":        "SearchForm",
					"DataQueries": []string{"GetNodeMetrics"},
					"Inputs": []interface{}{
						base_index_model.TableInputField{
							Name:     "TimeDuration",
							Type:     "Select",
							Data:     []string{"-3h", "-6h", "-1d", "-3d", "-7d"},
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
