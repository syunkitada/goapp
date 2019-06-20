package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const DatacenterKind = "Datacenter"

type Datacenter struct {
	gorm.Model
	Name         string `gorm:"not null;size:50;unique_index;"`
	Kind         string `gorm:"not null;size:25;"`
	Description  string `gorm:"not null;size:200;"`
	Region       string `gorm:"not null;size:50;unique_index;"`
	DomainSuffix string `gorm:"not null;size:255;unique;"`
	Spec         string `gorm:"not null;size:1000;"`
}

type DatacenterSpec struct {
	Kind         string `validate:"required"`
	Name         string `validate:"required"`
	Description  string
	Region       string `validate:"required"`
	DomainSuffix string `validate:"required"`
	Spec         interface{}
}

var DatacenterCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"CreateDatacenter": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: DatacenterKind,
		Help:    "helptext",
	},
	"UpdateDatacenter": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: DatacenterKind,
		Help:    "helptext",
	},
	"GetDatacenters": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     DatacenterKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Region", "DomainSuffix"},
	},
	"GetDatacenter": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: DatacenterKind,
		Help:    "helptext",
	},
	"DeleteDatacenter": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: DatacenterKind,
		Help:    "helptext",
	},
}

var DatacentersTable = index_model.Table{
	Name:    "Datacenters",
	Kind:    "Table",
	Route:   "",
	Subname: "datacenter",
	DataKey: "Datacenters",
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name:      "Name",
			IsSearch:  true,
			Link:      "Datacenters/:0/Resources/Resources",
			LinkParam: "datacenter",
			LinkSync:  true,
			LinkGetQueries: []string{
				"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
		},
		index_model.TableColumn{Name: "Region", IsSearch: true},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
	SelectActions: []index_model.Action{
		index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Datacenter",
			SelectKey: "Name",
		},
	},
	Actions: []index_model.Action{
		index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "Datacenter",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Options: []string{
						"Private", "Share",
					}},
			},
		},
	},
}
