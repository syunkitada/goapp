package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

type Meta struct{}

var Spec = base_spec_model.Spec{
	Meta: Meta{},
	Name: "ResourceApi",
	Kind: base_const.KindApi,
	Apis: []base_spec_model.Api{
		base_spec_model.Api{
			Name:            "ResourcePhysical",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				base_spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},

				base_spec_model.QueryModel{Req: GetDatacenter{}, Rep: GetDatacenterData{}},
				base_spec_model.QueryModel{Req: GetDatacenters{}, Rep: GetDatacentersData{}},

				base_spec_model.QueryModel{Req: GetFloor{}, Rep: GetFloorData{}},
				base_spec_model.QueryModel{Req: GetFloors{}, Rep: GetFloorsData{}},

				base_spec_model.QueryModel{Req: GetPhysicalModel{}, Rep: GetPhysicalModelData{}},
				base_spec_model.QueryModel{Req: GetPhysicalModels{}, Rep: GetPhysicalModelsData{}},

				base_spec_model.QueryModel{Req: GetPhysicalResource{}, Rep: GetPhysicalResourceData{}},
				base_spec_model.QueryModel{Req: GetPhysicalResources{}, Rep: GetPhysicalResourcesData{}},
			},
		},
		base_spec_model.Api{
			Name:            "ResourcePhysicalAdmin",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				base_spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				base_spec_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				base_spec_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				base_spec_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
				base_spec_model.QueryModel{Req: DeleteRegions{}, Rep: DeleteRegionsData{}},

				base_spec_model.QueryModel{Req: GetDatacenter{}, Rep: GetDatacenterData{}},
				base_spec_model.QueryModel{Req: GetDatacenters{}, Rep: GetDatacentersData{}},
				base_spec_model.QueryModel{Req: CreateDatacenter{}, Rep: CreateDatacenterData{}},
				base_spec_model.QueryModel{Req: UpdateDatacenter{}, Rep: UpdateDatacenterData{}},
				base_spec_model.QueryModel{Req: DeleteDatacenter{}, Rep: DeleteDatacenterData{}},
				base_spec_model.QueryModel{Req: DeleteDatacenters{}, Rep: DeleteDatacentersData{}},

				base_spec_model.QueryModel{Req: GetFloor{}, Rep: GetFloorData{}},
				base_spec_model.QueryModel{Req: GetFloors{}, Rep: GetFloorsData{}},
				base_spec_model.QueryModel{Req: CreateFloor{}, Rep: CreateFloorData{}},
				base_spec_model.QueryModel{Req: UpdateFloor{}, Rep: UpdateFloorData{}},
				base_spec_model.QueryModel{Req: DeleteFloor{}, Rep: DeleteFloorData{}},
				base_spec_model.QueryModel{Req: DeleteFloors{}, Rep: DeleteFloorsData{}},

				base_spec_model.QueryModel{Req: GetRack{}, Rep: GetRackData{}},
				base_spec_model.QueryModel{Req: GetRacks{}, Rep: GetRacksData{}},
				base_spec_model.QueryModel{Req: CreateRack{}, Rep: CreateRackData{}},
				base_spec_model.QueryModel{Req: UpdateRack{}, Rep: UpdateRackData{}},
				base_spec_model.QueryModel{Req: DeleteRack{}, Rep: DeleteRackData{}},
				base_spec_model.QueryModel{Req: DeleteRacks{}, Rep: DeleteRacksData{}},

				base_spec_model.QueryModel{Req: GetPhysicalModel{}, Rep: GetPhysicalModelData{}},
				base_spec_model.QueryModel{Req: GetPhysicalModels{}, Rep: GetPhysicalModelsData{}},
				base_spec_model.QueryModel{Req: CreatePhysicalModel{}, Rep: CreatePhysicalModelData{}},
				base_spec_model.QueryModel{Req: UpdatePhysicalModel{}, Rep: UpdatePhysicalModelData{}},
				base_spec_model.QueryModel{Req: DeletePhysicalModel{}, Rep: DeletePhysicalModelData{}},
				base_spec_model.QueryModel{Req: DeletePhysicalModels{}, Rep: DeletePhysicalModelsData{}},

				base_spec_model.QueryModel{Req: GetPhysicalResource{}, Rep: GetPhysicalResourceData{}},
				base_spec_model.QueryModel{Req: GetPhysicalResources{}, Rep: GetPhysicalResourcesData{}},
				base_spec_model.QueryModel{Req: CreatePhysicalResource{}, Rep: CreatePhysicalResourceData{}},
				base_spec_model.QueryModel{Req: UpdatePhysicalResource{}, Rep: UpdatePhysicalResourceData{}},
				base_spec_model.QueryModel{Req: DeletePhysicalResource{}, Rep: DeletePhysicalResourceData{}},
				base_spec_model.QueryModel{Req: DeletePhysicalResources{}, Rep: DeletePhysicalResourcesData{}},
			},
		},
		base_spec_model.Api{
			Name:            "ResourceVirtualAdmin",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				base_spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				base_spec_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				base_spec_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				base_spec_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
				base_spec_model.QueryModel{Req: DeleteRegions{}, Rep: DeleteRegionsData{}},

				base_spec_model.QueryModel{Req: GetCluster{}, Rep: GetClusterData{}},
				base_spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},
				base_spec_model.QueryModel{Req: CreateCluster{}, Rep: CreateClusterData{}},
				base_spec_model.QueryModel{Req: UpdateCluster{}, Rep: UpdateClusterData{}},
				base_spec_model.QueryModel{Req: DeleteCluster{}, Rep: DeleteClusterData{}},
				base_spec_model.QueryModel{Req: DeleteClusters{}, Rep: DeleteClustersData{}},

				base_spec_model.QueryModel{Req: GetNodes{}, Rep: GetNodesData{}},
				base_spec_model.QueryModel{Req: GetNodeServices{}, Rep: GetNodeServicesData{}},

				base_spec_model.QueryModel{Req: GetNetworkV4{}, Rep: GetNetworkV4Data{}},
				base_spec_model.QueryModel{Req: GetNetworkV4s{}, Rep: GetNetworkV4sData{}},
				base_spec_model.QueryModel{Req: CreateNetworkV4{}, Rep: CreateNetworkV4Data{}},
				base_spec_model.QueryModel{Req: UpdateNetworkV4{}, Rep: UpdateNetworkV4Data{}},
				base_spec_model.QueryModel{Req: DeleteNetworkV4{}, Rep: DeleteNetworkV4Data{}},
				base_spec_model.QueryModel{Req: DeleteNetworkV4s{}, Rep: DeleteNetworkV4sData{}},

				base_spec_model.QueryModel{Req: GetImage{}, Rep: GetImageData{}},
				base_spec_model.QueryModel{Req: GetImages{}, Rep: GetImagesData{}},
				base_spec_model.QueryModel{Req: CreateImage{}, Rep: CreateImageData{}},
				base_spec_model.QueryModel{Req: UpdateImage{}, Rep: UpdateImageData{}},
				base_spec_model.QueryModel{Req: DeleteImage{}, Rep: DeleteImageData{}},
				base_spec_model.QueryModel{Req: DeleteImages{}, Rep: DeleteImagesData{}},

				base_spec_model.QueryModel{Req: GetRegionService{}, Rep: GetRegionServiceData{}},
				base_spec_model.QueryModel{Req: GetRegionServices{}, Rep: GetRegionServicesData{}},
				base_spec_model.QueryModel{Req: CreateRegionService{}, Rep: CreateRegionServiceData{}},
				base_spec_model.QueryModel{Req: UpdateRegionService{}, Rep: UpdateRegionServiceData{}},
				base_spec_model.QueryModel{Req: DeleteRegionService{}, Rep: DeleteRegionServiceData{}},
				base_spec_model.QueryModel{Req: DeleteRegionServices{}, Rep: DeleteRegionServicesData{}},

				base_spec_model.QueryModel{Req: GetCompute{}, Rep: GetComputeData{}},
				base_spec_model.QueryModel{Req: GetComputes{}, Rep: GetComputesData{}},
				base_spec_model.QueryModel{Req: GetComputeConsole{}, Rep: GetComputeConsoleData{},
					Ws: true, Kind: "Terminal"},
			},
		},
		base_spec_model.Api{
			Name:            "ResourceVirtual",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: GetRegion{}, Rep: GetRegionData{}},
				base_spec_model.QueryModel{Req: GetRegions{}, Rep: GetRegionsData{}},
				base_spec_model.QueryModel{Req: CreateRegion{}, Rep: CreateRegionData{}},
				base_spec_model.QueryModel{Req: UpdateRegion{}, Rep: UpdateRegionData{}},
				base_spec_model.QueryModel{Req: DeleteRegion{}, Rep: DeleteRegionData{}},
				base_spec_model.QueryModel{Req: DeleteRegions{}, Rep: DeleteRegionsData{}},

				base_spec_model.QueryModel{Req: GetCluster{}, Rep: GetClusterData{}},
				base_spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},
				base_spec_model.QueryModel{Req: CreateCluster{}, Rep: CreateClusterData{}},
				base_spec_model.QueryModel{Req: UpdateCluster{}, Rep: UpdateClusterData{}},
				base_spec_model.QueryModel{Req: DeleteCluster{}, Rep: DeleteClusterData{}},
				base_spec_model.QueryModel{Req: DeleteClusters{}, Rep: DeleteClustersData{}},

				base_spec_model.QueryModel{Req: GetImage{}, Rep: GetImageData{}},
				base_spec_model.QueryModel{Req: GetImages{}, Rep: GetImagesData{}},
				base_spec_model.QueryModel{Req: CreateImage{}, Rep: CreateImageData{}},
				base_spec_model.QueryModel{Req: UpdateImage{}, Rep: UpdateImageData{}},
				base_spec_model.QueryModel{Req: DeleteImage{}, Rep: DeleteImageData{}},
				base_spec_model.QueryModel{Req: DeleteImages{}, Rep: DeleteImagesData{}},

				base_spec_model.QueryModel{Req: GetRegionService{}, Rep: GetRegionServiceData{}},
				base_spec_model.QueryModel{Req: GetRegionServices{}, Rep: GetRegionServicesData{}},
				base_spec_model.QueryModel{Req: CreateRegionService{}, Rep: CreateRegionServiceData{}},
				base_spec_model.QueryModel{Req: UpdateRegionService{}, Rep: UpdateRegionServiceData{}},
				base_spec_model.QueryModel{Req: DeleteRegionService{}, Rep: DeleteRegionServiceData{}},
				base_spec_model.QueryModel{Req: DeleteRegionServices{}, Rep: DeleteRegionServicesData{}},
			},
		},
		base_spec_model.Api{
			Name:            "ResourceMonitor",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: GetClusters{}, Rep: GetClustersData{}},

				base_spec_model.QueryModel{Req: GetNodes{}, Rep: GetNodesData{}},
				base_spec_model.QueryModel{Req: GetNode{}, Rep: GetNodeData{}},
				base_spec_model.QueryModel{Req: GetNodeMetrics{}, Rep: GetNodeMetricsData{}},
				base_spec_model.QueryModel{Req: GetStatistics{}, Rep: GetStatisticsData{}},
				base_spec_model.QueryModel{Req: GetLogParams{}, Rep: GetLogParamsData{}},
				base_spec_model.QueryModel{Req: GetLogs{}, Rep: GetLogsData{}},
				base_spec_model.QueryModel{Req: GetTrace{}, Rep: GetTraceData{}},

				base_spec_model.QueryModel{Req: GetEvents{}, Rep: GetEventsData{}},

				base_spec_model.QueryModel{Req: GetEventRule{}, Rep: GetEventRuleData{}},
				base_spec_model.QueryModel{Req: GetEventRules{}, Rep: GetEventRulesData{}},
				base_spec_model.QueryModel{Req: CreateEventRules{}, Rep: CreateEventRulesData{}},
				base_spec_model.QueryModel{Req: UpdateEventRules{}, Rep: UpdateEventRulesData{}},
				base_spec_model.QueryModel{Req: DeleteEventRules{}, Rep: DeleteEventRulesData{}},
			},
		},
	},
}
