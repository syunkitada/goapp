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

				base_model.QueryModel{Req: GetFloor{}, Rep: GetFloorData{}},
				base_model.QueryModel{Req: GetFloors{}, Rep: GetFloorsData{}},
				base_model.QueryModel{Req: CreateFloor{}, Rep: CreateFloorData{}},
				base_model.QueryModel{Req: UpdateFloor{}, Rep: UpdateFloorData{}},
				base_model.QueryModel{Req: DeleteFloor{}, Rep: DeleteFloorData{}},

				base_model.QueryModel{Req: GetRack{}, Rep: GetRackData{}},
				base_model.QueryModel{Req: GetRacks{}, Rep: GetRacksData{}},
				base_model.QueryModel{Req: CreateRack{}, Rep: CreateRackData{}},
				base_model.QueryModel{Req: UpdateRack{}, Rep: UpdateRackData{}},
				base_model.QueryModel{Req: DeleteRack{}, Rep: DeleteRackData{}},

				base_model.QueryModel{Req: GetPhysicalModel{}, Rep: GetPhysicalModelData{}},
				base_model.QueryModel{Req: GetPhysicalModels{}, Rep: GetPhysicalModelsData{}},
				base_model.QueryModel{Req: CreatePhysicalModel{}, Rep: CreatePhysicalModelData{}},
				base_model.QueryModel{Req: UpdatePhysicalModel{}, Rep: UpdatePhysicalModelData{}},
				base_model.QueryModel{Req: DeletePhysicalModel{}, Rep: DeletePhysicalModelData{}},

				base_model.QueryModel{Req: GetPhysicalResource{}, Rep: GetPhysicalResourceData{}},
				base_model.QueryModel{Req: GetPhysicalResources{}, Rep: GetPhysicalResourcesData{}},
				base_model.QueryModel{Req: CreatePhysicalResource{}, Rep: CreatePhysicalResourceData{}},
				base_model.QueryModel{Req: UpdatePhysicalResource{}, Rep: UpdatePhysicalResourceData{}},
				base_model.QueryModel{Req: DeletePhysicalResource{}, Rep: DeletePhysicalResourceData{}},
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
