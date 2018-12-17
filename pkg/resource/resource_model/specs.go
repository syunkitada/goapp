package resource_model

type ResourceSpec struct {
	Kind    string `validate:"required"`
	Name    string `validate:"required"`
	Cluster string `validate:"required"`
}

type NetworkV4Spec struct {
	ResourceSpec
	Spec NetworkV4SpecData `validate:"required"`
}

type NetworkV4SpecData struct {
	Kind    string `validate:"required"`
	Subnet  string `validate:"required"`
	StartIp string `validate:"required"`
	EndIp   string `validate:"required"`
	Gateway string `validate:"required"`
}

type ComputeSpec struct {
	ResourceSpec
	Spec ComputeSpecData `validate:"required"`
}

type ComputeSpecData struct {
	Kind       string `validate:"required"`
	Vcpus      int    `validate:"required"`
	MemorySize int    `validate:"required"`
	DiskSize   int    `validate:"required"`
	Image      string `validate:"required"`
	Network    string `validate:"required"`
}

type ContainerSpec struct {
	ResourceSpec
	Spec ContainerSpecData `validate:"required"`
}

type ContainerSpecData struct {
	Kind       string `validate:"required"`
	Vcpus      int    `validate:"required"`
	MemorySize int    `validate:"required"`
	DiskSize   int    `validate:"required"`
	Image      string `validate:"required"`
}

type ImageSpec struct {
	ResourceSpec
	Spec ImageSpecData `validate:"required"`
}

type ImageSpecData struct {
	Kind string `validate:"required"`
	Url  string `validate:"required"`
}

type VolumeSpec struct {
	ResourceSpec
	Spec VolumeSpecData `validate:"required"`
}

type VolumeSpecData struct {
	Kind string `validate:"required"`
}

type LoadbalancerSpec struct {
	ResourceSpec
	Spec LoadbalancerSpecData `validate:"required"`
}

type LoadbalancerSpecData struct {
	Kind string `validate:"required"`
}
