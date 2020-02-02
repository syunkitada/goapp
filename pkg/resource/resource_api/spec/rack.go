package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type Rack struct {
	Kind       string `validate:"required"`
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
	Floor      string `validate:"required"`
	Unit       uint8
}

type GetRack struct {
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
}

type GetRackData struct {
	Rack Rack
}

type GetRacks struct {
	Datacenter string `validate:"required"`
}

type GetRacksData struct {
	Racks []Rack
}

type CreateRack struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateRackData struct{}

type UpdateRack struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateRackData struct{}

type DeleteRack struct {
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
}

type DeleteRackData struct{}

type DeleteRacks struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteRacksData struct{}

var RacksTable = index_model.Table{
	Name:    "Racks",
	Route:   "/Racks",
	Kind:    "Table",
	DataKey: "Racks",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Rack",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Datacenters/:Datacenter/Resources/Racks/Detail/:0/View",
			LinkKey:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetRack"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
