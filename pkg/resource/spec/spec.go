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

				base_model.QueryModel{Req: GetDatacenter{}, Rep: GetDatacenterData{}},
				base_model.QueryModel{Req: GetDatacenters{}, Rep: GetDatacentersData{}},
				base_model.QueryModel{Req: CreateDatacenter{}, Rep: CreateDatacenterData{}},
				base_model.QueryModel{Req: UpdateDatacenter{}, Rep: UpdateDatacenterData{}},
				base_model.QueryModel{Req: DeleteDatacenter{}, Rep: DeleteDatacenterData{}},
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
