package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
)

type Meta struct{}

var Spec = spec_model.Spec{
	Meta: Meta{},
	Name: "ResourceApi",
	Kind: base_const.KindApi,
	Apis: []spec_model.Api{
		spec_model.Api{
			Name:            "ResourcePhysical",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},

				spec_model.QueryModel{Req: GetDatacenter{}, Rep: GetDatacenterData{}},
				spec_model.QueryModel{Req: GetDatacenters{}, Rep: GetDatacentersData{}},

				spec_model.QueryModel{Req: GetFloor{}, Rep: GetFloorData{}},
				spec_model.QueryModel{Req: GetFloors{}, Rep: GetFloorsData{}},

				spec_model.QueryModel{Req: GetPhysicalModel{}, Rep: GetPhysicalModelData{}},
				spec_model.QueryModel{Req: GetPhysicalModels{}, Rep: GetPhysicalModelsData{}},

				spec_model.QueryModel{Req: GetPhysicalResource{}, Rep: GetPhysicalResourceData{}},
				spec_model.QueryModel{Req: GetPhysicalResources{}, Rep: GetPhysicalResourcesData{}},
			},
		},
		spec_model.Api{
			Name:            "ResourcePhysicalAdmin",
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
		},
		spec_model.Api{
			Name:            "ResourceVirtualAdmin",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				spec_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				spec_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				spec_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
				spec_model.QueryModel{Req: DeleteRegions{}, Rep: DeleteRegionsData{}},

				spec_model.QueryModel{Req: GetCluster{}, Rep: GetClusterData{}},
				spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},
				spec_model.QueryModel{Req: CreateCluster{}, Rep: CreateClusterData{}},
				spec_model.QueryModel{Req: UpdateCluster{}, Rep: UpdateClusterData{}},
				spec_model.QueryModel{Req: DeleteCluster{}, Rep: DeleteClusterData{}},
				spec_model.QueryModel{Req: DeleteClusters{}, Rep: DeleteClustersData{}},

				spec_model.QueryModel{Req: GetNodes{}, Rep: GetNodesData{}},
				spec_model.QueryModel{Req: GetNodeServices{}, Rep: GetNodeServicesData{}},

				spec_model.QueryModel{Req: GetNetworkV4{}, Rep: GetNetworkV4Data{}},
				spec_model.QueryModel{Req: GetNetworkV4s{}, Rep: GetNetworkV4sData{}},
				spec_model.QueryModel{Req: CreateNetworkV4{}, Rep: CreateNetworkV4Data{}},
				spec_model.QueryModel{Req: UpdateNetworkV4{}, Rep: UpdateNetworkV4Data{}},
				spec_model.QueryModel{Req: DeleteNetworkV4{}, Rep: DeleteNetworkV4Data{}},
				spec_model.QueryModel{Req: DeleteNetworkV4s{}, Rep: DeleteNetworkV4sData{}},

				spec_model.QueryModel{Req: GetImage{}, Rep: GetImageData{}},
				spec_model.QueryModel{Req: GetImages{}, Rep: GetImagesData{}},
				spec_model.QueryModel{Req: CreateImage{}, Rep: CreateImageData{}},
				spec_model.QueryModel{Req: UpdateImage{}, Rep: UpdateImageData{}},
				spec_model.QueryModel{Req: DeleteImage{}, Rep: DeleteImageData{}},
				spec_model.QueryModel{Req: DeleteImages{}, Rep: DeleteImagesData{}},

				spec_model.QueryModel{Req: GetRegionService{}, Rep: GetRegionServiceData{}},
				spec_model.QueryModel{Req: GetRegionServices{}, Rep: GetRegionServicesData{}},
				spec_model.QueryModel{Req: CreateRegionService{}, Rep: CreateRegionServiceData{}},
				spec_model.QueryModel{Req: UpdateRegionService{}, Rep: UpdateRegionServiceData{}},
				spec_model.QueryModel{Req: DeleteRegionService{}, Rep: DeleteRegionServiceData{}},
				spec_model.QueryModel{Req: DeleteRegionServices{}, Rep: DeleteRegionServicesData{}},
			},
		},
		spec_model.Api{
			Name:            "ResourceVirtual",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				spec_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				spec_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				spec_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
				spec_model.QueryModel{Req: DeleteRegions{}, Rep: DeleteRegionsData{}},

				spec_model.QueryModel{Req: GetCluster{}, Rep: GetClusterData{}},
				spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},
				spec_model.QueryModel{Req: CreateCluster{}, Rep: CreateClusterData{}},
				spec_model.QueryModel{Req: UpdateCluster{}, Rep: UpdateClusterData{}},
				spec_model.QueryModel{Req: DeleteCluster{}, Rep: DeleteClusterData{}},
				spec_model.QueryModel{Req: DeleteClusters{}, Rep: DeleteClustersData{}},

				spec_model.QueryModel{Req: GetImage{}, Rep: GetImageData{}},
				spec_model.QueryModel{Req: GetImages{}, Rep: GetImagesData{}},
				spec_model.QueryModel{Req: CreateImage{}, Rep: CreateImageData{}},
				spec_model.QueryModel{Req: UpdateImage{}, Rep: UpdateImageData{}},
				spec_model.QueryModel{Req: DeleteImage{}, Rep: DeleteImageData{}},
				spec_model.QueryModel{Req: DeleteImages{}, Rep: DeleteImagesData{}},

				spec_model.QueryModel{Req: GetRegionService{}, Rep: GetRegionServiceData{}},
				spec_model.QueryModel{Req: GetRegionServices{}, Rep: GetRegionServicesData{}},
				spec_model.QueryModel{Req: CreateRegionService{}, Rep: CreateRegionServiceData{}},
				spec_model.QueryModel{Req: UpdateRegionService{}, Rep: UpdateRegionServiceData{}},
				spec_model.QueryModel{Req: DeleteRegionService{}, Rep: DeleteRegionServiceData{}},
				spec_model.QueryModel{Req: DeleteRegionServices{}, Rep: DeleteRegionServicesData{}},
			},
		},
		spec_model.Api{
			Name:            "ResourceMonitor",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},

				spec_model.QueryModel{Req: GetNodes{}, Rep: GetNodesData{}},
				spec_model.QueryModel{Req: GetNode{}, Rep: GetNodeData{}},
				spec_model.QueryModel{Req: GetAlerts{}, Rep: GetAlertsData{}},
				spec_model.QueryModel{Req: GetAlertRules{}, Rep: GetAlertRulesData{}},
				spec_model.QueryModel{Req: GetStatistics{}, Rep: GetStatisticsData{}},
				spec_model.QueryModel{Req: GetLogs{}, Rep: GetLogsData{}},
				spec_model.QueryModel{Req: GetTrace{}, Rep: GetTraceData{}},
			},
		},
	},
}
