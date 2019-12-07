package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

type NodeServiceSpec struct {
	NumaNodeServices []NumaNodeServiceSpec
	Storages         []StorageSpec
}

type NumaNodeServiceSpec struct {
	Id              int
	AvailableCpus   int
	UsedCpus        int
	AvailableMemory int
	UsedMemory      int
}

type StorageSpec struct {
	Kind                      string
	Path                      string
	AvailableGb               int
	UsedGb                    int
	AvailableNumaNodeServices []int
}

type GetNodeServices struct {
	Cluster string
}

type GetNodeServicesData struct {
	NodeServices []base_spec.NodeService
}

type SyncNodeService struct {
	NodeService base_spec.NodeService
}

type SyncNodeServiceData struct {
	Task NodeServiceTask
}

type NodeServiceTask struct {
	ComputeAssignments []ComputeAssignmentEx
}

type ComputeAssignmentEx struct {
	ID        uint
	Status    string
	Spec      RegionServiceComputeSpec
	UpdatedAt time.Time
}
type AssignmentReports struct {
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

type ReportNodeServiceTask struct {
	ComputeAssignmentReports []AssignmentReport
}

type ReportNodeServiceTaskData struct {
}

type ReportNode struct {
	Project   string
	Name      string
	State     string
	Warning   string
	Warnings  int
	Error     string
	Errors    int
	Timestate time.Time
	Logs      []ResourceLog
	Metrics   []ResourceMetric
	Alerts    []ResourceAlert
}

type ReportNodeData struct {
}

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

type ResourceAlert struct {
	Name    string
	Time    string
	Level   string
	Handler string
	Msg     string
	Tag     map[string]string
}

type GetAlerts struct {
	Cluster string `validate:"required"`
}

type GetAlertsData struct {
	Alerts []ResourceAlert
}

type GetAlertRules struct {
	Cluster string `validate:"required"`
}

type AlertRule struct {
	Node  string
	Name  string
	Kind  string
	Until time.Time
}

type GetAlertRulesData struct {
	AlertRules []AlertRule
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
	Cluster string `validate:"required"`
}

type GetLogsData struct {
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
	Selectors: []index_model.TableSelector{
		index_model.TableSelector{
			Name:    "App",
			DataKey: "LogApps",
		},
		index_model.TableSelector{
			Name:    "Node",
			DataKey: "LogNodes",
		},
	},
	InputFields: []index_model.TableInputField{
		index_model.TableInputField{
			Name: "TraceId",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
		},
		index_model.TableColumn{Name: "App"},
		index_model.TableColumn{Name: "Host"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var AlertsTable = index_model.Table{
	Name:    "Alerts",
	Route:   "/Alerts",
	Kind:    "Table",
	DataKey: "Alerts",
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

var AlertRulesTable = index_model.Table{
	Name:    "AlertRules",
	Route:   "/AlertRules",
	Kind:    "Table",
	DataKey: "AlertRules",
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
