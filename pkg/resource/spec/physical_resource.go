package spec

type PhysicalResource struct {
	Kind          string `validate:"required"`
	Name          string `validate:"required"`
	Datacenter    string `validate:"required"`
	Cluster       string
	Rack          string
	PhysicalModel string
	RackPosition  uint8
	Spec          string
}

type GetPhysicalResource struct {
	Datacenter string `validate:"required"`
	Name       string `validate:"required"`
}

type GetPhysicalResourceData struct {
	PhysicalResource PhysicalResource
}

type GetPhysicalResources struct{}

type GetPhysicalResourcesData struct {
	PhysicalResources []PhysicalResource
}

type CreatePhysicalResource struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreatePhysicalResourceData struct{}

type UpdatePhysicalResource struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdatePhysicalResourceData struct{}

type DeletePhysicalResource struct {
	Datacenter string `validate:"required"`
	Name       string `validate:"required"`
}

type DeletePhysicalResourceData struct {
	PhysicalResource PhysicalResource
}
