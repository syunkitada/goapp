package spec

import "github.com/syunkitada/goapp/pkg/base/base_index_model"

type Floor struct {
	Kind       string `validate:"required"`
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
	Zone       string `validate:"required"`
	Floor      uint8  `validate:"required"`
}

type GetFloor struct {
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
}

type GetFloorData struct {
	Floor Floor
}

type GetFloors struct {
	Datacenter string `validate:"required"`
}

type GetFloorsData struct {
	Floors []Floor
}

type CreateFloor struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateFloorData struct{}

type UpdateFloor struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateFloorData struct{}

type DeleteFloor struct {
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
}

type DeleteFloorData struct{}

type DeleteFloors struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteFloorsData struct{}

var FloorsTable = base_index_model.Table{
	Name:    "Floors",
	Route:   "/Floors",
	Kind:    "Table",
	DataKey: "Floors",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Floor",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Datacenters/:Datacenter/Resources/Floors/Detail/:0/View",
			LinkKey:        "Name",
			LinkSync:       false,
			LinkDataQueries: []string{"GetFloor"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
