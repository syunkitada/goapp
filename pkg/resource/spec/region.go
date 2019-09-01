package spec

type GetRegion struct {
	Name string
}

type GetRegionData struct {
	Name string
	Kind string
}

type GetRegions struct{}

type GetRegionsData struct {
	Regions []GetRegionData
}
