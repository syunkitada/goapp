package spec

type Compute struct {
	Region        string
	Cluster       string
	RegionService string
	Name          string
	Kind          string
	Labels        string
	Status        string
	StatusReason  string
	Spec          string
	Project       string
	LinkSpec      string
	Image         string
	Vcpus         uint
	Memory        uint
	Disk          uint
}
