package spec

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

type DeleteRegionData struct {
	Region Region
}
