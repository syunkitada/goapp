package resource_model

type ResourceSpec struct {
	Kind    string `validate:"required"`
	Name    string `validate:"required"`
	Cluster string `validate:"required"`
	Spec    interface{}
}

type DatacenterSpec struct {
	ResourceSpec
	Spec DatacenterSpecData `validate:"required"`
}

type DatacenterSpecData struct {
	Kind         string `validate:"required"`
	Region       string `validate:"required"`
	DomainSuffix string `validate:"required"`
	Spec         interface{}
}

type ClusterSpec struct {
	ResourceSpec
	Spec ClusterSpecData `validate:"required"`
}

type ClusterSpecData struct {
	Kind         string `validate:"required"`
	Datacenter   string `validate:"required"`
	DomainSuffix string `validate:"required"`
	Spec         interface{}
}

type FloorSpec struct {
	ResourceSpec
	Spec FloorSpecData `validate:"required"`
}

type FloorSpecData struct {
	Kind       string `validate:"required"`
	Datacenter string `validate:"required"`
	Zone       string `validate:"required"`
	Floor      uint8  `validate:"required"`
	Spec       interface{}
}

type RackSpec struct {
	ResourceSpec
	Spec RackSpecData `validate:"required"`
}

type RackSpecData struct {
	Kind       string `validate:"required"`
	Datacenter string `validate:"required"`
	Floor      string `validate:"required"`
	Spec       interface{}
}

type PhysicalModelSpec struct {
	ResourceSpec
	Spec PhysicalModelSpecData `validate:"required"`
}

type PhysicalModelSpecData struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Unit        uint8  `validate:"required"`
	Description string
	Spec        interface{}
}

type PhysicalResourceSpec struct {
	ResourceSpec
	Spec PhysicalResourceSpecData `validate:"required"`
}

type PhysicalResourceSpecData struct {
	Kind         string `validate:"required"`
	Name         string `validate:"required"`
	Datacenter   string `validate:"required"`
	Cluster      string
	Rack         string
	Model        string
	RackPosition uint8
	NetLinks     []string
	PowerLinks   []string
	Spec         string
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
	Domain     string `validate:"required"`
	Kind       string `validate:"required"`
	Vcpus      int    `validate:"required"`
	MemorySize int    `validate:"required"`
	DiskSize   int    `validate:"required"`
	Image      string `validate:"required"`
	Network    string `validate:"required"`
	Networks   []string
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
