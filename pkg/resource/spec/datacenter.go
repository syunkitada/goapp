package spec

type Datacenter struct {
	Kind         string `validate:"required"`
	Name         string `validate:"required"`
	Description  string
	Region       string `validate:"required"`
	DomainSuffix string `validate:"required"`
}

type GetDatacenter struct {
	Name string `validate:"required"`
}

type GetDatacenterData struct {
	Datacenter Datacenter
}

type GetDatacenters struct{}

type GetDatacentersData struct {
	Datacenters []Datacenter
}

type CreateDatacenter struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateDatacenterData struct{}

type UpdateDatacenter struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateDatacenterData struct{}

type DeleteDatacenter struct {
	Name string `validate:"required"`
}

type DeleteDatacenterData struct {
	Datacenter Datacenter
}
