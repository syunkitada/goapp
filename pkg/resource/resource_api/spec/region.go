package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

type Region struct {
	Name      string `validate:"required"`
	Kind      string `validate:"required"`
	UpdatedAt time.Time
	CreatedAt time.Time
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

var RegionsTable = base_index_model.Table{
	Name:    "Regions",
	Kind:    "Table",
	DataKey: "Regions",
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name:       "Name",
			IsSearch:   true,
			Align:      "left",
			LinkPath:   []string{"Regions", "RegionResources", "Clusters"},
			LinkKeyMap: map[string]string{"Region": "Name"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time", Sort: "asc"},
	},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Region",
			SelectKey: "Name",
		},
	},
	Actions: []base_index_model.Action{
		base_index_model.Action{
			Name: "Create", Icon: "Create", Kind: "Form",
			DataKind: "Region",
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
