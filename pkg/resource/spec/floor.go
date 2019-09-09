package spec

type Floor struct {
	Kind       string `validate:"required"`
	Name       string `validate:"required"`
	Datacenter string `validate:"required"`
	Zone       string `validate:"required"`
	Floor      uint8  `validate:"required"`
}

type GetFloor struct {
	Name string `validate:"required"`
}

type GetFloorData struct {
	Floor Floor
}

type GetFloors struct{}

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
	Name string `validate:"required"`
}

type DeleteFloorData struct {
	Floor Floor
}
