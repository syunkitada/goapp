package spec

type PhysicalModel struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Unit        uint8
	Description string
	Spec        interface{}
}

type GetPhysicalModel struct {
	Name string `validate:"required"`
}

type GetPhysicalModelData struct {
	PhysicalModel PhysicalModel
}

type GetPhysicalModels struct{}

type GetPhysicalModelsData struct {
	PhysicalModels []PhysicalModel
}

type CreatePhysicalModel struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreatePhysicalModelData struct{}

type UpdatePhysicalModel struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdatePhysicalModelData struct{}

type DeletePhysicalModel struct {
	Name string `validate:"required"`
}

type DeletePhysicalModelData struct {
	PhysicalModel PhysicalModel
}
