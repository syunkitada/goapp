package genpkg

import "github.com/syunkitada/goapp/pkg/base/base_model/index_model"

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
				"get_physical-resources", "get_racks", "get_floors", "get_physical-models"},
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
