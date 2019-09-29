package spec

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
