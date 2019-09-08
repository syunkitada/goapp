package spec

type Region struct {
	Name string
	Kind string
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
	Spec string `validate:"required" flagKine:"file"`
}

type CreateRegionData struct {
	Region Region
}

type UpdateRegion struct {
	Spec string `validate:"required" flagKine:"file"`
}

type UpdateRegionData struct {
	Region Region
}

type DeleteRegion struct {
	Name string `validate:"required"`
}

type DeleteRegionData struct {
	Region Region
}
