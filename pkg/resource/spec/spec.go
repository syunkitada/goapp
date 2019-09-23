package spec

import "github.com/syunkitada/goapp/pkg/base/base_model/spec_model"

type Meta struct{}

var Spec = spec_model.Spec{
	Meta: Meta{},
	Name: "Resource",
	Apis: []spec_model.Api{
		spec_model.Api{
			Name:            "ResourcePhysical",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				spec_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				spec_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				spec_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
				spec_model.QueryModel{Req: DeleteRegions{}, Rep: DeleteRegionsData{}},

				spec_model.QueryModel{Req: GetDatacenter{}, Rep: GetDatacenterData{}},
				spec_model.QueryModel{Req: GetDatacenters{}, Rep: GetDatacentersData{}},
				spec_model.QueryModel{Req: CreateDatacenter{}, Rep: CreateDatacenterData{}},
				spec_model.QueryModel{Req: UpdateDatacenter{}, Rep: UpdateDatacenterData{}},
				spec_model.QueryModel{Req: DeleteDatacenter{}, Rep: DeleteDatacenterData{}},
				spec_model.QueryModel{Req: DeleteDatacenters{}, Rep: DeleteDatacentersData{}},

				spec_model.QueryModel{Req: GetFloor{}, Rep: GetFloorData{}},
				spec_model.QueryModel{Req: GetFloors{}, Rep: GetFloorsData{}},
				spec_model.QueryModel{Req: CreateFloor{}, Rep: CreateFloorData{}},
				spec_model.QueryModel{Req: UpdateFloor{}, Rep: UpdateFloorData{}},
				spec_model.QueryModel{Req: DeleteFloor{}, Rep: DeleteFloorData{}},
				spec_model.QueryModel{Req: DeleteFloors{}, Rep: DeleteFloorsData{}},

				spec_model.QueryModel{Req: GetRack{}, Rep: GetRackData{}},
				spec_model.QueryModel{Req: GetRacks{}, Rep: GetRacksData{}},
				spec_model.QueryModel{Req: CreateRack{}, Rep: CreateRackData{}},
				spec_model.QueryModel{Req: UpdateRack{}, Rep: UpdateRackData{}},
				spec_model.QueryModel{Req: DeleteRack{}, Rep: DeleteRackData{}},
				spec_model.QueryModel{Req: DeleteRacks{}, Rep: DeleteRacksData{}},

				spec_model.QueryModel{Req: GetPhysicalModel{}, Rep: GetPhysicalModelData{}},
				spec_model.QueryModel{Req: GetPhysicalModels{}, Rep: GetPhysicalModelsData{}},
				spec_model.QueryModel{Req: CreatePhysicalModel{}, Rep: CreatePhysicalModelData{}},
				spec_model.QueryModel{Req: UpdatePhysicalModel{}, Rep: UpdatePhysicalModelData{}},
				spec_model.QueryModel{Req: DeletePhysicalModel{}, Rep: DeletePhysicalModelData{}},
				spec_model.QueryModel{Req: DeletePhysicalModels{}, Rep: DeletePhysicalModelsData{}},

				spec_model.QueryModel{Req: GetPhysicalResource{}, Rep: GetPhysicalResourceData{}},
				spec_model.QueryModel{Req: GetPhysicalResources{}, Rep: GetPhysicalResourcesData{}},
				spec_model.QueryModel{Req: CreatePhysicalResource{}, Rep: CreatePhysicalResourceData{}},
				spec_model.QueryModel{Req: UpdatePhysicalResource{}, Rep: UpdatePhysicalResourceData{}},
				spec_model.QueryModel{Req: DeletePhysicalResource{}, Rep: DeletePhysicalResourceData{}},
				spec_model.QueryModel{Req: DeletePhysicalResources{}, Rep: DeletePhysicalResourcesData{}},
			},
			ViewModels: []interface{}{
				spec_model.Table{
					Name:          "Datacenters",
					Data:          GetDatacenter{},
					Actions:       []interface{}{CreateDatacenter{}},
					SelectActions: []interface{}{DeleteDatacenter{}},
					ColumnLinkMap: map[string]spec_model.Link{
						"Name": spec_model.Link{Target: "Resources/PhysicalResources"},
					},
				},
				spec_model.Tab{
					Name: "Resources",
					Tabs: []interface{}{
						spec_model.Table{
							Name:          "PhysicalResources",
							Data:          GetPhysicalResources{},
							Actions:       []interface{}{CreateDatacenter{}},
							SelectActions: []interface{}{DeleteDatacenter{}},
						},
						spec_model.Table{
							Name:          "Racks",
							Data:          GetRacks{},
							Actions:       []interface{}{CreateRack{}},
							SelectActions: []interface{}{DeleteRack{}},
						},
						spec_model.Table{
							Name:          "Floor",
							Data:          GetFloors{},
							Actions:       []interface{}{CreateFloor{}},
							SelectActions: []interface{}{DeleteFloor{}},
						},
						spec_model.Table{
							Name:          "PhysicalModels",
							Data:          GetPhysicalModels{},
							Actions:       []interface{}{CreatePhysicalModel{}},
							SelectActions: []interface{}{DeletePhysicalModel{}},
						},
					},
				},
			},
		},
		spec_model.Api{
			Name:            "ResourceVirtual",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},
			},
		},
	},
}
