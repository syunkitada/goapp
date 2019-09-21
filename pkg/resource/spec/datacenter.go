package spec

import "time"

type Datacenter struct {
	Kind         string `validate:"required"`
	Name         string `validate:"required" view:"isSearch=true"`
	Description  string
	Region       string    `validate:"required" view:"isSearch=true"`
	DomainSuffix string    `validate:"required"`
	UpdatedAt    time.Time `view:"sort=asc"`
	CreatedAt    time.Time
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
