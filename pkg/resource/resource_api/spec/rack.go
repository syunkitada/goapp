package spec

import "github.com/syunkitada/goapp/pkg/base/base_index_model"

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

var RacksTable = base_index_model.Table{
	Name:    "Racks",
	Route:   "/Racks",
	Kind:    "Table",
	DataKey: "Racks",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Rack",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:            "Datacenters/:Datacenter/Resources/Racks/Detail/:0/View",
			LinkKeyMap:      map[string]string{"Name": "Name"},
			LinkSync:        false,
			LinkDataQueries: []string{"GetRack"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
