package spec

import (
	"time"

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
	Metric map[string]float64
}

type ResourceAlert struct {
	Name    string
	Time    string
	Level   string
	Handler string
	Msg     string
	Tag     map[string]string
}
