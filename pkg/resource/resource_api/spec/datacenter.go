package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

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

type DeleteDatacenterData struct{}

type DeleteDatacenters struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteDatacentersData struct{}

var DatacentersTable = base_index_model.Table{
	Name:        "Datacenters",
	Kind:        "Table",
	Route:       "",
	DataQueries: []string{"GetDatacenters"},
	DataKey:     "Datacenters",
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name:       "Name",
			IsSearch:   true,
			Link:       "Datacenters/:0/Resources/Resources",
			LinkKeyMap: map[string]string{"Datacenter": "Datacenter"},
			LinkSync:   true,
			LinkDataQueries: []string{
				"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
		},
		base_index_model.TableColumn{Name: "Region", IsSearch: true},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Datacenter",
			SelectKey: "Name",
		},
	},
	Actions: []base_index_model.Action{
		base_index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "Datacenter",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text", Required: true,
					Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				base_index_model.Field{Name: "Kind", Kind: "select", Required: true,
					Options: []string{
						"Private", "Share",
					}},
			},
		},
	},
}
