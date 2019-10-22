package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type Compute struct {
	Region        string
	Cluster       string
	RegionService string
	Name          string
	Kind          string
	Labels        string
	Status        string
	StatusReason  string
	Project       string
	Spec          interface{}
	LinkSpec      string
	Image         string
	Vcpus         uint
	Memory        uint
	Disk          uint
}

type GetCompute struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type GetComputeData struct {
	Compute Compute
}

type GetComputes struct {
	Region string `validate:"required"`
}

type GetComputesData struct {
	Computes []Compute
}

type CreateCompute struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateComputeData struct{}

type UpdateCompute struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateComputeData struct{}

type DeleteCompute struct {
	Name string `validate:"required"`
}

type DeleteComputeData struct{}

type DeleteComputes struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteComputesData struct{}

var ComputesTable = index_model.Table{
	Name:    "Computes",
	Route:   "/Computes",
	Kind:    "Table",
	DataKey: "Computes",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Compute",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Datacenters/:Datacenter/Resources/Computes/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetCompute"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
