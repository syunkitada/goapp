package spec

import "github.com/syunkitada/goapp/pkg/base/base_model"

type Meta struct{}

var Spec = base_model.Spec{
	Meta: Meta{},
	Name: "Resource",
	Apis: []base_model.Api{
		base_model.Api{
			Name:            "ResourcePhysical",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				base_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				base_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				base_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				base_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
			},
		},
		base_model.Api{
			Name:            "ResourceVirtual",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},
			},
		},
	},
}
