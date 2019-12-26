// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Client struct {
	*base_client.Client
}

func NewClient(conf *base_config.ClientConfig) *Client {
	client := Client{
		Client: base_client.NewClient(conf),
	}
	return &client
}

type CreateClusterResponse struct {
	base_model.Response
	ResultMap CreateClusterResultMap
}

type CreateClusterResultMap struct {
	CreateCluster CreateClusterResult
}

type CreateClusterResult struct {
	Code  uint8
	Error string
	Data  spec.CreateClusterData
}
type CreateDatacenterResponse struct {
	base_model.Response
	ResultMap CreateDatacenterResultMap
}

type CreateDatacenterResultMap struct {
	CreateDatacenter CreateDatacenterResult
}

type CreateDatacenterResult struct {
	Code  uint8
	Error string
	Data  spec.CreateDatacenterData
}
type CreateEventRulesResponse struct {
	base_model.Response
	ResultMap CreateEventRulesResultMap
}

type CreateEventRulesResultMap struct {
	CreateEventRules CreateEventRulesResult
}

type CreateEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.CreateEventRulesData
}
type CreateFloorResponse struct {
	base_model.Response
	ResultMap CreateFloorResultMap
}

type CreateFloorResultMap struct {
	CreateFloor CreateFloorResult
}

type CreateFloorResult struct {
	Code  uint8
	Error string
	Data  spec.CreateFloorData
}
type CreateImageResponse struct {
	base_model.Response
	ResultMap CreateImageResultMap
}

type CreateImageResultMap struct {
	CreateImage CreateImageResult
}

type CreateImageResult struct {
	Code  uint8
	Error string
	Data  spec.CreateImageData
}
type CreateNetworkV4Response struct {
	base_model.Response
	ResultMap CreateNetworkV4ResultMap
}

type CreateNetworkV4ResultMap struct {
	CreateNetworkV4 CreateNetworkV4Result
}

type CreateNetworkV4Result struct {
	Code  uint8
	Error string
	Data  spec.CreateNetworkV4Data
}
type CreatePhysicalModelResponse struct {
	base_model.Response
	ResultMap CreatePhysicalModelResultMap
}

type CreatePhysicalModelResultMap struct {
	CreatePhysicalModel CreatePhysicalModelResult
}

type CreatePhysicalModelResult struct {
	Code  uint8
	Error string
	Data  spec.CreatePhysicalModelData
}
type CreatePhysicalResourceResponse struct {
	base_model.Response
	ResultMap CreatePhysicalResourceResultMap
}

type CreatePhysicalResourceResultMap struct {
	CreatePhysicalResource CreatePhysicalResourceResult
}

type CreatePhysicalResourceResult struct {
	Code  uint8
	Error string
	Data  spec.CreatePhysicalResourceData
}
type CreateRackResponse struct {
	base_model.Response
	ResultMap CreateRackResultMap
}

type CreateRackResultMap struct {
	CreateRack CreateRackResult
}

type CreateRackResult struct {
	Code  uint8
	Error string
	Data  spec.CreateRackData
}
type CreateRegionResponse struct {
	base_model.Response
	ResultMap CreateRegionResultMap
}

type CreateRegionResultMap struct {
	CreateRegion CreateRegionResult
}

type CreateRegionResult struct {
	Code  uint8
	Error string
	Data  spec.CreateRegionData
}
type CreateRegionServiceResponse struct {
	base_model.Response
	ResultMap CreateRegionServiceResultMap
}

type CreateRegionServiceResultMap struct {
	CreateRegionService CreateRegionServiceResult
}

type CreateRegionServiceResult struct {
	Code  uint8
	Error string
	Data  spec.CreateRegionServiceData
}
type DeleteClusterResponse struct {
	base_model.Response
	ResultMap DeleteClusterResultMap
}

type DeleteClusterResultMap struct {
	DeleteCluster DeleteClusterResult
}

type DeleteClusterResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteClusterData
}
type DeleteClustersResponse struct {
	base_model.Response
	ResultMap DeleteClustersResultMap
}

type DeleteClustersResultMap struct {
	DeleteClusters DeleteClustersResult
}

type DeleteClustersResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteClustersData
}
type DeleteDatacenterResponse struct {
	base_model.Response
	ResultMap DeleteDatacenterResultMap
}

type DeleteDatacenterResultMap struct {
	DeleteDatacenter DeleteDatacenterResult
}

type DeleteDatacenterResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteDatacenterData
}
type DeleteDatacentersResponse struct {
	base_model.Response
	ResultMap DeleteDatacentersResultMap
}

type DeleteDatacentersResultMap struct {
	DeleteDatacenters DeleteDatacentersResult
}

type DeleteDatacentersResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteDatacentersData
}
type DeleteEventRulesResponse struct {
	base_model.Response
	ResultMap DeleteEventRulesResultMap
}

type DeleteEventRulesResultMap struct {
	DeleteEventRules DeleteEventRulesResult
}

type DeleteEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteEventRulesData
}
type DeleteFloorResponse struct {
	base_model.Response
	ResultMap DeleteFloorResultMap
}

type DeleteFloorResultMap struct {
	DeleteFloor DeleteFloorResult
}

type DeleteFloorResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteFloorData
}
type DeleteFloorsResponse struct {
	base_model.Response
	ResultMap DeleteFloorsResultMap
}

type DeleteFloorsResultMap struct {
	DeleteFloors DeleteFloorsResult
}

type DeleteFloorsResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteFloorsData
}
type DeleteImageResponse struct {
	base_model.Response
	ResultMap DeleteImageResultMap
}

type DeleteImageResultMap struct {
	DeleteImage DeleteImageResult
}

type DeleteImageResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteImageData
}
type DeleteImagesResponse struct {
	base_model.Response
	ResultMap DeleteImagesResultMap
}

type DeleteImagesResultMap struct {
	DeleteImages DeleteImagesResult
}

type DeleteImagesResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteImagesData
}
type DeleteNetworkV4Response struct {
	base_model.Response
	ResultMap DeleteNetworkV4ResultMap
}

type DeleteNetworkV4ResultMap struct {
	DeleteNetworkV4 DeleteNetworkV4Result
}

type DeleteNetworkV4Result struct {
	Code  uint8
	Error string
	Data  spec.DeleteNetworkV4Data
}
type DeleteNetworkV4sResponse struct {
	base_model.Response
	ResultMap DeleteNetworkV4sResultMap
}

type DeleteNetworkV4sResultMap struct {
	DeleteNetworkV4s DeleteNetworkV4sResult
}

type DeleteNetworkV4sResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteNetworkV4sData
}
type DeletePhysicalModelResponse struct {
	base_model.Response
	ResultMap DeletePhysicalModelResultMap
}

type DeletePhysicalModelResultMap struct {
	DeletePhysicalModel DeletePhysicalModelResult
}

type DeletePhysicalModelResult struct {
	Code  uint8
	Error string
	Data  spec.DeletePhysicalModelData
}
type DeletePhysicalModelsResponse struct {
	base_model.Response
	ResultMap DeletePhysicalModelsResultMap
}

type DeletePhysicalModelsResultMap struct {
	DeletePhysicalModels DeletePhysicalModelsResult
}

type DeletePhysicalModelsResult struct {
	Code  uint8
	Error string
	Data  spec.DeletePhysicalModelsData
}
type DeletePhysicalResourceResponse struct {
	base_model.Response
	ResultMap DeletePhysicalResourceResultMap
}

type DeletePhysicalResourceResultMap struct {
	DeletePhysicalResource DeletePhysicalResourceResult
}

type DeletePhysicalResourceResult struct {
	Code  uint8
	Error string
	Data  spec.DeletePhysicalResourceData
}
type DeletePhysicalResourcesResponse struct {
	base_model.Response
	ResultMap DeletePhysicalResourcesResultMap
}

type DeletePhysicalResourcesResultMap struct {
	DeletePhysicalResources DeletePhysicalResourcesResult
}

type DeletePhysicalResourcesResult struct {
	Code  uint8
	Error string
	Data  spec.DeletePhysicalResourcesData
}
type DeleteRackResponse struct {
	base_model.Response
	ResultMap DeleteRackResultMap
}

type DeleteRackResultMap struct {
	DeleteRack DeleteRackResult
}

type DeleteRackResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteRackData
}
type DeleteRacksResponse struct {
	base_model.Response
	ResultMap DeleteRacksResultMap
}

type DeleteRacksResultMap struct {
	DeleteRacks DeleteRacksResult
}

type DeleteRacksResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteRacksData
}
type DeleteRegionResponse struct {
	base_model.Response
	ResultMap DeleteRegionResultMap
}

type DeleteRegionResultMap struct {
	DeleteRegion DeleteRegionResult
}

type DeleteRegionResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteRegionData
}
type DeleteRegionServiceResponse struct {
	base_model.Response
	ResultMap DeleteRegionServiceResultMap
}

type DeleteRegionServiceResultMap struct {
	DeleteRegionService DeleteRegionServiceResult
}

type DeleteRegionServiceResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteRegionServiceData
}
type DeleteRegionServicesResponse struct {
	base_model.Response
	ResultMap DeleteRegionServicesResultMap
}

type DeleteRegionServicesResultMap struct {
	DeleteRegionServices DeleteRegionServicesResult
}

type DeleteRegionServicesResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteRegionServicesData
}
type DeleteRegionsResponse struct {
	base_model.Response
	ResultMap DeleteRegionsResultMap
}

type DeleteRegionsResultMap struct {
	DeleteRegions DeleteRegionsResult
}

type DeleteRegionsResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteRegionsData
}
type GetClusterResponse struct {
	base_model.Response
	ResultMap GetClusterResultMap
}

type GetClusterResultMap struct {
	GetCluster GetClusterResult
}

type GetClusterResult struct {
	Code  uint8
	Error string
	Data  spec.GetClusterData
}
type GetClustersResponse struct {
	base_model.Response
	ResultMap GetClustersResultMap
}

type GetClustersResultMap struct {
	GetClusters GetClustersResult
}

type GetClustersResult struct {
	Code  uint8
	Error string
	Data  spec.GetClustersData
}
type GetDatacenterResponse struct {
	base_model.Response
	ResultMap GetDatacenterResultMap
}

type GetDatacenterResultMap struct {
	GetDatacenter GetDatacenterResult
}

type GetDatacenterResult struct {
	Code  uint8
	Error string
	Data  spec.GetDatacenterData
}
type GetDatacentersResponse struct {
	base_model.Response
	ResultMap GetDatacentersResultMap
}

type GetDatacentersResultMap struct {
	GetDatacenters GetDatacentersResult
}

type GetDatacentersResult struct {
	Code  uint8
	Error string
	Data  spec.GetDatacentersData
}
type GetEventRuleResponse struct {
	base_model.Response
	ResultMap GetEventRuleResultMap
}

type GetEventRuleResultMap struct {
	GetEventRule GetEventRuleResult
}

type GetEventRuleResult struct {
	Code  uint8
	Error string
	Data  spec.GetEventRuleData
}
type GetEventRulesResponse struct {
	base_model.Response
	ResultMap GetEventRulesResultMap
}

type GetEventRulesResultMap struct {
	GetEventRules GetEventRulesResult
}

type GetEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.GetEventRulesData
}
type GetEventsResponse struct {
	base_model.Response
	ResultMap GetEventsResultMap
}

type GetEventsResultMap struct {
	GetEvents GetEventsResult
}

type GetEventsResult struct {
	Code  uint8
	Error string
	Data  spec.GetEventsData
}
type GetFloorResponse struct {
	base_model.Response
	ResultMap GetFloorResultMap
}

type GetFloorResultMap struct {
	GetFloor GetFloorResult
}

type GetFloorResult struct {
	Code  uint8
	Error string
	Data  spec.GetFloorData
}
type GetFloorsResponse struct {
	base_model.Response
	ResultMap GetFloorsResultMap
}

type GetFloorsResultMap struct {
	GetFloors GetFloorsResult
}

type GetFloorsResult struct {
	Code  uint8
	Error string
	Data  spec.GetFloorsData
}
type GetImageResponse struct {
	base_model.Response
	ResultMap GetImageResultMap
}

type GetImageResultMap struct {
	GetImage GetImageResult
}

type GetImageResult struct {
	Code  uint8
	Error string
	Data  spec.GetImageData
}
type GetImagesResponse struct {
	base_model.Response
	ResultMap GetImagesResultMap
}

type GetImagesResultMap struct {
	GetImages GetImagesResult
}

type GetImagesResult struct {
	Code  uint8
	Error string
	Data  spec.GetImagesData
}
type GetLogParamsResponse struct {
	base_model.Response
	ResultMap GetLogParamsResultMap
}

type GetLogParamsResultMap struct {
	GetLogParams GetLogParamsResult
}

type GetLogParamsResult struct {
	Code  uint8
	Error string
	Data  spec.GetLogParamsData
}
type GetLogsResponse struct {
	base_model.Response
	ResultMap GetLogsResultMap
}

type GetLogsResultMap struct {
	GetLogs GetLogsResult
}

type GetLogsResult struct {
	Code  uint8
	Error string
	Data  spec.GetLogsData
}
type GetNetworkV4Response struct {
	base_model.Response
	ResultMap GetNetworkV4ResultMap
}

type GetNetworkV4ResultMap struct {
	GetNetworkV4 GetNetworkV4Result
}

type GetNetworkV4Result struct {
	Code  uint8
	Error string
	Data  spec.GetNetworkV4Data
}
type GetNetworkV4sResponse struct {
	base_model.Response
	ResultMap GetNetworkV4sResultMap
}

type GetNetworkV4sResultMap struct {
	GetNetworkV4s GetNetworkV4sResult
}

type GetNetworkV4sResult struct {
	Code  uint8
	Error string
	Data  spec.GetNetworkV4sData
}
type GetNodeResponse struct {
	base_model.Response
	ResultMap GetNodeResultMap
}

type GetNodeResultMap struct {
	GetNode GetNodeResult
}

type GetNodeResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodeData
}
type GetNodeServicesResponse struct {
	base_model.Response
	ResultMap GetNodeServicesResultMap
}

type GetNodeServicesResultMap struct {
	GetNodeServices GetNodeServicesResult
}

type GetNodeServicesResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodeServicesData
}
type GetNodesResponse struct {
	base_model.Response
	ResultMap GetNodesResultMap
}

type GetNodesResultMap struct {
	GetNodes GetNodesResult
}

type GetNodesResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodesData
}
type GetPhysicalModelResponse struct {
	base_model.Response
	ResultMap GetPhysicalModelResultMap
}

type GetPhysicalModelResultMap struct {
	GetPhysicalModel GetPhysicalModelResult
}

type GetPhysicalModelResult struct {
	Code  uint8
	Error string
	Data  spec.GetPhysicalModelData
}
type GetPhysicalModelsResponse struct {
	base_model.Response
	ResultMap GetPhysicalModelsResultMap
}

type GetPhysicalModelsResultMap struct {
	GetPhysicalModels GetPhysicalModelsResult
}

type GetPhysicalModelsResult struct {
	Code  uint8
	Error string
	Data  spec.GetPhysicalModelsData
}
type GetPhysicalResourceResponse struct {
	base_model.Response
	ResultMap GetPhysicalResourceResultMap
}

type GetPhysicalResourceResultMap struct {
	GetPhysicalResource GetPhysicalResourceResult
}

type GetPhysicalResourceResult struct {
	Code  uint8
	Error string
	Data  spec.GetPhysicalResourceData
}
type GetPhysicalResourcesResponse struct {
	base_model.Response
	ResultMap GetPhysicalResourcesResultMap
}

type GetPhysicalResourcesResultMap struct {
	GetPhysicalResources GetPhysicalResourcesResult
}

type GetPhysicalResourcesResult struct {
	Code  uint8
	Error string
	Data  spec.GetPhysicalResourcesData
}
type GetRackResponse struct {
	base_model.Response
	ResultMap GetRackResultMap
}

type GetRackResultMap struct {
	GetRack GetRackResult
}

type GetRackResult struct {
	Code  uint8
	Error string
	Data  spec.GetRackData
}
type GetRacksResponse struct {
	base_model.Response
	ResultMap GetRacksResultMap
}

type GetRacksResultMap struct {
	GetRacks GetRacksResult
}

type GetRacksResult struct {
	Code  uint8
	Error string
	Data  spec.GetRacksData
}
type GetRegionResponse struct {
	base_model.Response
	ResultMap GetRegionResultMap
}

type GetRegionResultMap struct {
	GetRegion GetRegionResult
}

type GetRegionResult struct {
	Code  uint8
	Error string
	Data  spec.GetRegionData
}
type GetRegionServiceResponse struct {
	base_model.Response
	ResultMap GetRegionServiceResultMap
}

type GetRegionServiceResultMap struct {
	GetRegionService GetRegionServiceResult
}

type GetRegionServiceResult struct {
	Code  uint8
	Error string
	Data  spec.GetRegionServiceData
}
type GetRegionServicesResponse struct {
	base_model.Response
	ResultMap GetRegionServicesResultMap
}

type GetRegionServicesResultMap struct {
	GetRegionServices GetRegionServicesResult
}

type GetRegionServicesResult struct {
	Code  uint8
	Error string
	Data  spec.GetRegionServicesData
}
type GetRegionsResponse struct {
	base_model.Response
	ResultMap GetRegionsResultMap
}

type GetRegionsResultMap struct {
	GetRegions GetRegionsResult
}

type GetRegionsResult struct {
	Code  uint8
	Error string
	Data  spec.GetRegionsData
}
type GetStatisticsResponse struct {
	base_model.Response
	ResultMap GetStatisticsResultMap
}

type GetStatisticsResultMap struct {
	GetStatistics GetStatisticsResult
}

type GetStatisticsResult struct {
	Code  uint8
	Error string
	Data  spec.GetStatisticsData
}
type GetTraceResponse struct {
	base_model.Response
	ResultMap GetTraceResultMap
}

type GetTraceResultMap struct {
	GetTrace GetTraceResult
}

type GetTraceResult struct {
	Code  uint8
	Error string
	Data  spec.GetTraceData
}
type UpdateClusterResponse struct {
	base_model.Response
	ResultMap UpdateClusterResultMap
}

type UpdateClusterResultMap struct {
	UpdateCluster UpdateClusterResult
}

type UpdateClusterResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateClusterData
}
type UpdateDatacenterResponse struct {
	base_model.Response
	ResultMap UpdateDatacenterResultMap
}

type UpdateDatacenterResultMap struct {
	UpdateDatacenter UpdateDatacenterResult
}

type UpdateDatacenterResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateDatacenterData
}
type UpdateEventRulesResponse struct {
	base_model.Response
	ResultMap UpdateEventRulesResultMap
}

type UpdateEventRulesResultMap struct {
	UpdateEventRules UpdateEventRulesResult
}

type UpdateEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateEventRulesData
}
type UpdateFloorResponse struct {
	base_model.Response
	ResultMap UpdateFloorResultMap
}

type UpdateFloorResultMap struct {
	UpdateFloor UpdateFloorResult
}

type UpdateFloorResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateFloorData
}
type UpdateImageResponse struct {
	base_model.Response
	ResultMap UpdateImageResultMap
}

type UpdateImageResultMap struct {
	UpdateImage UpdateImageResult
}

type UpdateImageResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateImageData
}
type UpdateNetworkV4Response struct {
	base_model.Response
	ResultMap UpdateNetworkV4ResultMap
}

type UpdateNetworkV4ResultMap struct {
	UpdateNetworkV4 UpdateNetworkV4Result
}

type UpdateNetworkV4Result struct {
	Code  uint8
	Error string
	Data  spec.UpdateNetworkV4Data
}
type UpdatePhysicalModelResponse struct {
	base_model.Response
	ResultMap UpdatePhysicalModelResultMap
}

type UpdatePhysicalModelResultMap struct {
	UpdatePhysicalModel UpdatePhysicalModelResult
}

type UpdatePhysicalModelResult struct {
	Code  uint8
	Error string
	Data  spec.UpdatePhysicalModelData
}
type UpdatePhysicalResourceResponse struct {
	base_model.Response
	ResultMap UpdatePhysicalResourceResultMap
}

type UpdatePhysicalResourceResultMap struct {
	UpdatePhysicalResource UpdatePhysicalResourceResult
}

type UpdatePhysicalResourceResult struct {
	Code  uint8
	Error string
	Data  spec.UpdatePhysicalResourceData
}
type UpdateRackResponse struct {
	base_model.Response
	ResultMap UpdateRackResultMap
}

type UpdateRackResultMap struct {
	UpdateRack UpdateRackResult
}

type UpdateRackResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateRackData
}
type UpdateRegionResponse struct {
	base_model.Response
	ResultMap UpdateRegionResultMap
}

type UpdateRegionResultMap struct {
	UpdateRegion UpdateRegionResult
}

type UpdateRegionResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateRegionData
}
type UpdateRegionServiceResponse struct {
	base_model.Response
	ResultMap UpdateRegionServiceResultMap
}

type UpdateRegionServiceResultMap struct {
	UpdateRegionService UpdateRegionServiceResult
}

type UpdateRegionServiceResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateRegionServiceData
}

func (client *Client) ResourcePhysicalGetRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionData, err error) {
	var res GetRegionResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionsData, err error) {
	var res GetRegionsResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetDatacenter(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetDatacenterData, err error) {
	var res GetDatacenterResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetDatacenter
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetDatacenters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetDatacentersData, err error) {
	var res GetDatacentersResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetDatacenters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetFloor(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetFloorData, err error) {
	var res GetFloorResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetFloor
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetFloors(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetFloorsData, err error) {
	var res GetFloorsResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetFloors
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetPhysicalModel(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalModelData, err error) {
	var res GetPhysicalModelResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalModel
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetPhysicalModels(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalModelsData, err error) {
	var res GetPhysicalModelsResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalModels
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetPhysicalResource(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalResourceData, err error) {
	var res GetPhysicalResourceResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalResource
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalGetPhysicalResources(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalResourcesData, err error) {
	var res GetPhysicalResourcesResponse
	err = client.Request(tctx, "ResourcePhysical", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalResources
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionData, err error) {
	var res GetRegionResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionsData, err error) {
	var res GetRegionsResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminCreateRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateRegionData, err error) {
	var res CreateRegionResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminUpdateRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateRegionData, err error) {
	var res UpdateRegionResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionData, err error) {
	var res DeleteRegionResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionsData, err error) {
	var res DeleteRegionsResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetDatacenter(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetDatacenterData, err error) {
	var res GetDatacenterResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetDatacenter
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetDatacenters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetDatacentersData, err error) {
	var res GetDatacentersResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetDatacenters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminCreateDatacenter(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateDatacenterData, err error) {
	var res CreateDatacenterResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateDatacenter
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminUpdateDatacenter(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateDatacenterData, err error) {
	var res UpdateDatacenterResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateDatacenter
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteDatacenter(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteDatacenterData, err error) {
	var res DeleteDatacenterResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteDatacenter
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteDatacenters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteDatacentersData, err error) {
	var res DeleteDatacentersResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteDatacenters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetFloor(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetFloorData, err error) {
	var res GetFloorResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetFloor
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetFloors(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetFloorsData, err error) {
	var res GetFloorsResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetFloors
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminCreateFloor(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateFloorData, err error) {
	var res CreateFloorResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateFloor
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminUpdateFloor(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateFloorData, err error) {
	var res UpdateFloorResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateFloor
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteFloor(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteFloorData, err error) {
	var res DeleteFloorResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteFloor
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteFloors(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteFloorsData, err error) {
	var res DeleteFloorsResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteFloors
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetRack(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRackData, err error) {
	var res GetRackResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRack
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetRacks(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRacksData, err error) {
	var res GetRacksResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRacks
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminCreateRack(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateRackData, err error) {
	var res CreateRackResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateRack
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminUpdateRack(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateRackData, err error) {
	var res UpdateRackResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateRack
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteRack(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRackData, err error) {
	var res DeleteRackResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRack
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeleteRacks(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRacksData, err error) {
	var res DeleteRacksResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRacks
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetPhysicalModel(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalModelData, err error) {
	var res GetPhysicalModelResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalModel
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetPhysicalModels(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalModelsData, err error) {
	var res GetPhysicalModelsResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalModels
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminCreatePhysicalModel(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreatePhysicalModelData, err error) {
	var res CreatePhysicalModelResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreatePhysicalModel
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminUpdatePhysicalModel(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdatePhysicalModelData, err error) {
	var res UpdatePhysicalModelResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdatePhysicalModel
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeletePhysicalModel(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeletePhysicalModelData, err error) {
	var res DeletePhysicalModelResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeletePhysicalModel
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeletePhysicalModels(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeletePhysicalModelsData, err error) {
	var res DeletePhysicalModelsResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeletePhysicalModels
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetPhysicalResource(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalResourceData, err error) {
	var res GetPhysicalResourceResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalResource
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminGetPhysicalResources(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetPhysicalResourcesData, err error) {
	var res GetPhysicalResourcesResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetPhysicalResources
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminCreatePhysicalResource(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreatePhysicalResourceData, err error) {
	var res CreatePhysicalResourceResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreatePhysicalResource
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminUpdatePhysicalResource(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdatePhysicalResourceData, err error) {
	var res UpdatePhysicalResourceResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdatePhysicalResource
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeletePhysicalResource(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeletePhysicalResourceData, err error) {
	var res DeletePhysicalResourceResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeletePhysicalResource
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourcePhysicalAdminDeletePhysicalResources(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeletePhysicalResourcesData, err error) {
	var res DeletePhysicalResourcesResponse
	err = client.Request(tctx, "ResourcePhysicalAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeletePhysicalResources
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionData, err error) {
	var res GetRegionResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionsData, err error) {
	var res GetRegionsResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateRegionData, err error) {
	var res CreateRegionResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateRegionData, err error) {
	var res UpdateRegionResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionData, err error) {
	var res DeleteRegionResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionsData, err error) {
	var res DeleteRegionsResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetClusterData, err error) {
	var res GetClusterResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetClusters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetClustersData, err error) {
	var res GetClustersResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetClusters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateClusterData, err error) {
	var res CreateClusterResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateClusterData, err error) {
	var res UpdateClusterResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteClusterData, err error) {
	var res DeleteClusterResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteClusters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteClustersData, err error) {
	var res DeleteClustersResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteClusters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetNodes(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodesData, err error) {
	var res GetNodesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNodes
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetNodeServices(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodeServicesData, err error) {
	var res GetNodeServicesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNodeServices
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetNetworkV4(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNetworkV4Data, err error) {
	var res GetNetworkV4Response
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNetworkV4
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetNetworkV4s(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNetworkV4sData, err error) {
	var res GetNetworkV4sResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNetworkV4s
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateNetworkV4(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateNetworkV4Data, err error) {
	var res CreateNetworkV4Response
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateNetworkV4
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateNetworkV4(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateNetworkV4Data, err error) {
	var res UpdateNetworkV4Response
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateNetworkV4
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteNetworkV4(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteNetworkV4Data, err error) {
	var res DeleteNetworkV4Response
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteNetworkV4
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteNetworkV4s(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteNetworkV4sData, err error) {
	var res DeleteNetworkV4sResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteNetworkV4s
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetImageData, err error) {
	var res GetImageResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetImages(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetImagesData, err error) {
	var res GetImagesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetImages
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateImageData, err error) {
	var res CreateImageResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateImageData, err error) {
	var res UpdateImageResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteImageData, err error) {
	var res DeleteImageResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteImages(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteImagesData, err error) {
	var res DeleteImagesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteImages
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionServiceData, err error) {
	var res GetRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetRegionServices(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionServicesData, err error) {
	var res GetRegionServicesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegionServices
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateRegionServiceData, err error) {
	var res CreateRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateRegionServiceData, err error) {
	var res UpdateRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionServiceData, err error) {
	var res DeleteRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteRegionServices(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionServicesData, err error) {
	var res DeleteRegionServicesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegionServices
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionData, err error) {
	var res GetRegionResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionsData, err error) {
	var res GetRegionsResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualCreateRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateRegionData, err error) {
	var res CreateRegionResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualUpdateRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateRegionData, err error) {
	var res UpdateRegionResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteRegion(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionData, err error) {
	var res DeleteRegionResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegion
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteRegions(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionsData, err error) {
	var res DeleteRegionsResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegions
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetClusterData, err error) {
	var res GetClusterResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetClusters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetClustersData, err error) {
	var res GetClustersResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetClusters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualCreateCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateClusterData, err error) {
	var res CreateClusterResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualUpdateCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateClusterData, err error) {
	var res UpdateClusterResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteClusterData, err error) {
	var res DeleteClusterResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteCluster
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteClusters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteClustersData, err error) {
	var res DeleteClustersResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteClusters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetImageData, err error) {
	var res GetImageResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetImages(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetImagesData, err error) {
	var res GetImagesResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetImages
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualCreateImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateImageData, err error) {
	var res CreateImageResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualUpdateImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateImageData, err error) {
	var res UpdateImageResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteImage(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteImageData, err error) {
	var res DeleteImageResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteImage
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteImages(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteImagesData, err error) {
	var res DeleteImagesResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteImages
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionServiceData, err error) {
	var res GetRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualGetRegionServices(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetRegionServicesData, err error) {
	var res GetRegionServicesResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetRegionServices
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualCreateRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateRegionServiceData, err error) {
	var res CreateRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualUpdateRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateRegionServiceData, err error) {
	var res UpdateRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteRegionService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionServiceData, err error) {
	var res DeleteRegionServiceResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegionService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualDeleteRegionServices(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteRegionServicesData, err error) {
	var res DeleteRegionServicesResponse
	err = client.Request(tctx, "ResourceVirtual", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteRegionServices
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetClusters(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetClustersData, err error) {
	var res GetClustersResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetClusters
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetNodes(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodesData, err error) {
	var res GetNodesResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNodes
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetNode(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodeData, err error) {
	var res GetNodeResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNode
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetStatistics(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetStatisticsData, err error) {
	var res GetStatisticsResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetStatistics
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetLogParams(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetLogParamsData, err error) {
	var res GetLogParamsResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetLogParams
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetLogs(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetLogsData, err error) {
	var res GetLogsResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetLogs
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetTrace(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetTraceData, err error) {
	var res GetTraceResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetTrace
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetEvents(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetEventsData, err error) {
	var res GetEventsResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetEvents
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetEventRule(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetEventRuleData, err error) {
	var res GetEventRuleResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetEventRule
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorGetEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetEventRulesData, err error) {
	var res GetEventRulesResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorCreateEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateEventRulesData, err error) {
	var res CreateEventRulesResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorUpdateEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateEventRulesData, err error) {
	var res UpdateEventRulesResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceMonitorDeleteEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteEventRulesData, err error) {
	var res DeleteEventRulesResponse
	err = client.Request(tctx, "ResourceMonitor", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
