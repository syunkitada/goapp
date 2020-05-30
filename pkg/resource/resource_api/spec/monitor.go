package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
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

var LogsTable = base_index_model.Table{
	Name:        "Logs",
	Route:       "/Logs",
	Kind:        "Table",
	DataQueries: []string{"GetLogParams", "GetLogs"},
	DataKey:     "Logs",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	SearchForm: base_index_model.SearchForm{
		Kind: "SearchForm",
		Inputs: []base_index_model.TableInputField{
			base_index_model.TableInputField{
				Name:     "Apps",
				Type:     "Selector",
				DataKey:  "LogApps",
				Multiple: true,
			},
			base_index_model.TableInputField{
				Name:     "Nodes",
				Type:     "Selector",
				DataKey:  "LogNodes",
				Multiple: true,
			},
			base_index_model.TableInputField{
				Name: "TraceId",
				Type: "Text",
			},
			base_index_model.TableInputField{
				Name:     "LimitLogs",
				Type:     "Selector",
				Data:     []string{"5k", "10k", "20k", "30k", "40k", "50k"},
				Default:  "10k",
				Multiple: false,
			},
			base_index_model.TableInputField{
				Name:     "FromTime",
				Type:     "Selector",
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
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{Name: "App"},
		base_index_model.TableColumn{Name: "Node"},
		base_index_model.TableColumn{Name: "Msg", IsSearch: true},
		base_index_model.TableColumn{Name: "Func"},
		base_index_model.TableColumn{Name: "Level"},
		base_index_model.TableColumn{Name: "TraceId"},
		base_index_model.TableColumn{Name: "Time", Kind: "Time"},
	},
}

var StatisticsTable = base_index_model.Table{
	Name:    "Statistics",
	Route:   "/Statistics",
	Kind:    "Table",
	DataKey: "Statistics",
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
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
