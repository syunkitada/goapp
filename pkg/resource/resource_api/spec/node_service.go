package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

type NodeServiceSpec struct {
	NumaNodes []NumaNodeSpec
	Storages  []StorageSpec
}

type NumaNodeSpec struct {
	Id            int
	Cpus          []NumaNodeCpuSpec
	TotalMemory   int
	UsedMemory    int
	Total1GMemory int
	Used1GMemory  int
}

type NumaNodeCpuSpec struct {
	Reserved   bool
	Used       bool
	PhysicalId int // numa
	CoreId     int // core
	Processor  int // thread
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
	Events    []ResourceEvent
}

type ReportNodeData struct {
}
