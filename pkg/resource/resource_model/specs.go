package resource_model

type ComputeLibvirtSpec struct {
	Vcpus      int
	MemorySize int
	DiskSize   int
	Image      string
}
