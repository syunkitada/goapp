package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

type ResourceLog struct {
	Name string
	Time string
	Log  map[string]string
}

type ResourceMetric struct {
	Name   string
	Time   string
	Tag    map[string]string
	Metric map[string]interface{}
}

type GetEventRules struct {
	Cluster string `validate:"required"`
}

type EventRule struct {
	Node  string
	Name  string
	Kind  string
	Until time.Time
}

type GetEventRulesData struct {
	EventRules []EventRule
}

type GetStatistics struct {
	Cluster string `validate:"required"`
}

type GetStatisticsData struct {
}

type GetLogParams struct {
	Cluster string `validate:"required"`
}

type GetLogParamsData struct {
	LogNodes []string
	LogApps  []string
}

type GetLogs struct {
	Cluster   string `validate:"required"`
	Project   string
	LimitLogs string
	FromTime  string
	UntilTime time.Time
	Apps      []string
	Nodes     []string
	TraceId   string
}

type GetLogsData struct {
	Logs []map[string]interface{}
}

type GetTrace struct {
	Cluster string `validate:"required"`
	TraceId string `validate:"required"`
}

type GetTraceData struct {
}

var LogsTable = index_model.Table{
	Name:    "Logs",
	Route:   "/Logs",
	Kind:    "Table",
	DataKey: "Logs",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	ExInputs: []index_model.TableInputField{
		index_model.TableInputField{
			Name:     "Apps",
			Type:     "Selector",
			DataKey:  "LogApps",
			Multiple: true,
		},
		index_model.TableInputField{
			Name:     "Nodes",
			Type:     "Selector",
			DataKey:  "LogNodes",
			Multiple: true,
		},
		index_model.TableInputField{
			Name: "TraceId",
			Type: "Text",
		},
		index_model.TableInputField{
			Name:     "LimitLogs",
			Type:     "Selector",
			Data:     []string{"5k", "10k", "20k", "30k", "40k", "50k"},
			Default:  "10k",
			Multiple: false,
		},
		index_model.TableInputField{
			Name:     "FromTime",
			Type:     "Selector",
			Data:     []string{"-6h", "-1d", "-3d"},
			Default:  "-6h",
			Multiple: false,
		},
		index_model.TableInputField{
			Name: "UntilTime",
			Type: "DateTime",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{Name: "App"},
		index_model.TableColumn{Name: "Node"},
		index_model.TableColumn{Name: "Msg", IsSearch: true},
		index_model.TableColumn{Name: "Func"},
		index_model.TableColumn{Name: "Level"},
		index_model.TableColumn{Name: "TraceId"},
		index_model.TableColumn{Name: "Time", Kind: "Time"},
	},
}

var EventsTable = index_model.Table{
	Name:    "Events",
	Route:   "/Events",
	Kind:    "Table",
	DataKey: "Events",
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
	},
}

var EventRulesTable = index_model.Table{
	Name:    "EventRules",
	Route:   "/EventRules",
	Kind:    "Table",
	DataKey: "EventRules",
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
		},
		index_model.TableColumn{Name: "Host"},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "Until", Kind: "Time"},
	},
}

var StatisticsTable = index_model.Table{
	Name:    "Statistics",
	Route:   "/Statistics",
	Kind:    "Table",
	DataKey: "Statistics",
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
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
