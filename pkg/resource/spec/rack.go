package spec

type Rack struct {
	Kind       string `validate:"required"`
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
	Floor      string `validate:"required"`
	Unit       uint8
}

type GetRack struct {
	Name string `validate:"required"`
}

type GetRackData struct {
	Rack Rack
}

type GetRacks struct{}

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
	Name string `validate:"required"`
}

type DeleteRackData struct {
	Rack Rack
}
