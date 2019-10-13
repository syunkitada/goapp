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
