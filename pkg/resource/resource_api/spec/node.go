package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

type NodeSpec struct {
	NumaNodes []NumaNodeSpec
	Storages  []StorageSpec
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

type GetNodes struct {
	Cluster string
}

type GetNodesData struct {
	Nodes []base_spec.Node
}

type SyncNode struct {
	Node base_spec.Node
}

type SyncNodeData struct {
	Task NodeTask
}

type NodeTask struct {
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

type ReportNodeTask struct {
	ComputeAssignmentReports []AssignmentReport
}

type ReportNodeTaskData struct {
}

type ReportResource struct {
	Warning  string
	Warnings int
	Error    string
	Errors   int
	Logs     []ResourceLog
	Metrics  []ResourceMetric
	Alerts   []ResourceAlert
}

type ReportResourceData struct {
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
