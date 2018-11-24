package resource_model

type ResourceSpec struct {
	Kind    string
	Name    string
	Cluster string
	Spec    Spec
}

type Spec struct {
	Kind string
}

type ComputeSpec struct {
	ResourceSpec
	Spec ComputeSpecData
}

type ComputeSpecData struct {
	Spec
	Kind       string
	Vcpus      int
	MemorySize int
	DiskSize   int
	Image      string
}

type ContainerSpec struct {
	ResourceSpec
	Spec ContainerSpecData
}

type ContainerSpecData struct {
	Spec
	Kind       string
	Vcpus      int
	MemorySize int
	DiskSize   int
	Image      string
}

type ImageSpec struct {
	ResourceSpec
	Spec ImageSpecData
}

type ImageSpecData struct {
	Spec
	Kind string
	Url  string
}

type VolumeSpec struct {
	ResourceSpec
	Spec VolumeSpecData
}

type VolumeSpecData struct {
	Spec
	Kind string
}

type LoadbalancerSpec struct {
	ResourceSpec
	Spec LoadbalancerSpecData
}

type LoadbalancerSpecData struct {
	Spec
	Kind string
}
