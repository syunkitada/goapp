package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type Region struct {
	Name string `validate:"required"`
	Kind string `validate:"required"`
}

type GetRegion struct {
	Name string `validate:"required"`
}

type GetRegionData struct {
	Region Region
}

type GetRegions struct{}

type GetRegionsData struct {
	Regions []Region
}

type CreateRegion struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateRegionData struct{}

type UpdateRegion struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateRegionData struct{}

type DeleteRegion struct {
	Name string `validate:"required"`
}

type DeleteRegionData struct{}

type DeleteRegions struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteRegionsData struct{}

var RegionsTable = index_model.Table{
	Name:    "Regions",
	Kind:    "Table",
	Route:   "",
	DataKey: "Regions",
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name:      "Name",
			IsSearch:  true,
			Link:      "Regions/:0/Resources/Resources",
			LinkParam: "Region",
			LinkSync:  true,
			LinkGetQueries: []string{
				"GetRegionServices", "GetImages"},
		},
		index_model.TableColumn{Name: "Region", IsSearch: true},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
	SelectActions: []index_model.Action{
		index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Region",
			SelectKey: "Name",
		},
	},
	Actions: []index_model.Action{
		index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "Region",
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
