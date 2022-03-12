// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceIndexData, uint8, error)
	GetProjectServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	GetProjectServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	CreateCluster(tctx *logger.TraceContext, input *spec.CreateCluster, user *base_spec.UserAuthority) (*spec.CreateClusterData, uint8, error)
	CreateDatacenter(tctx *logger.TraceContext, input *spec.CreateDatacenter, user *base_spec.UserAuthority) (*spec.CreateDatacenterData, uint8, error)
	CreateEventRules(tctx *logger.TraceContext, input *spec.CreateEventRules, user *base_spec.UserAuthority) (*spec.CreateEventRulesData, uint8, error)
	CreateFloor(tctx *logger.TraceContext, input *spec.CreateFloor, user *base_spec.UserAuthority) (*spec.CreateFloorData, uint8, error)
	CreateImage(tctx *logger.TraceContext, input *spec.CreateImage, user *base_spec.UserAuthority) (*spec.CreateImageData, uint8, error)
	CreateNetworkV4(tctx *logger.TraceContext, input *spec.CreateNetworkV4, user *base_spec.UserAuthority) (*spec.CreateNetworkV4Data, uint8, error)
	CreatePhysicalModel(tctx *logger.TraceContext, input *spec.CreatePhysicalModel, user *base_spec.UserAuthority) (*spec.CreatePhysicalModelData, uint8, error)
	CreatePhysicalResource(tctx *logger.TraceContext, input *spec.CreatePhysicalResource, user *base_spec.UserAuthority) (*spec.CreatePhysicalResourceData, uint8, error)
	CreateRack(tctx *logger.TraceContext, input *spec.CreateRack, user *base_spec.UserAuthority) (*spec.CreateRackData, uint8, error)
	CreateRegion(tctx *logger.TraceContext, input *spec.CreateRegion, user *base_spec.UserAuthority) (*spec.CreateRegionData, uint8, error)
	CreateRegionService(tctx *logger.TraceContext, input *spec.CreateRegionService, user *base_spec.UserAuthority) (*spec.CreateRegionServiceData, uint8, error)
	DeleteCluster(tctx *logger.TraceContext, input *spec.DeleteCluster, user *base_spec.UserAuthority) (*spec.DeleteClusterData, uint8, error)
	DeleteClusters(tctx *logger.TraceContext, input *spec.DeleteClusters, user *base_spec.UserAuthority) (*spec.DeleteClustersData, uint8, error)
	DeleteDatacenter(tctx *logger.TraceContext, input *spec.DeleteDatacenter, user *base_spec.UserAuthority) (*spec.DeleteDatacenterData, uint8, error)
	DeleteDatacenters(tctx *logger.TraceContext, input *spec.DeleteDatacenters, user *base_spec.UserAuthority) (*spec.DeleteDatacentersData, uint8, error)
	DeleteEventRules(tctx *logger.TraceContext, input *spec.DeleteEventRules, user *base_spec.UserAuthority) (*spec.DeleteEventRulesData, uint8, error)
	DeleteFloor(tctx *logger.TraceContext, input *spec.DeleteFloor, user *base_spec.UserAuthority) (*spec.DeleteFloorData, uint8, error)
	DeleteFloors(tctx *logger.TraceContext, input *spec.DeleteFloors, user *base_spec.UserAuthority) (*spec.DeleteFloorsData, uint8, error)
	DeleteImage(tctx *logger.TraceContext, input *spec.DeleteImage, user *base_spec.UserAuthority) (*spec.DeleteImageData, uint8, error)
	DeleteImages(tctx *logger.TraceContext, input *spec.DeleteImages, user *base_spec.UserAuthority) (*spec.DeleteImagesData, uint8, error)
	DeleteNetworkV4(tctx *logger.TraceContext, input *spec.DeleteNetworkV4, user *base_spec.UserAuthority) (*spec.DeleteNetworkV4Data, uint8, error)
	DeleteNetworkV4s(tctx *logger.TraceContext, input *spec.DeleteNetworkV4s, user *base_spec.UserAuthority) (*spec.DeleteNetworkV4sData, uint8, error)
	DeletePhysicalModel(tctx *logger.TraceContext, input *spec.DeletePhysicalModel, user *base_spec.UserAuthority) (*spec.DeletePhysicalModelData, uint8, error)
	DeletePhysicalModels(tctx *logger.TraceContext, input *spec.DeletePhysicalModels, user *base_spec.UserAuthority) (*spec.DeletePhysicalModelsData, uint8, error)
	DeletePhysicalResource(tctx *logger.TraceContext, input *spec.DeletePhysicalResource, user *base_spec.UserAuthority) (*spec.DeletePhysicalResourceData, uint8, error)
	DeletePhysicalResources(tctx *logger.TraceContext, input *spec.DeletePhysicalResources, user *base_spec.UserAuthority) (*spec.DeletePhysicalResourcesData, uint8, error)
	DeleteRack(tctx *logger.TraceContext, input *spec.DeleteRack, user *base_spec.UserAuthority) (*spec.DeleteRackData, uint8, error)
	DeleteRacks(tctx *logger.TraceContext, input *spec.DeleteRacks, user *base_spec.UserAuthority) (*spec.DeleteRacksData, uint8, error)
	DeleteRegion(tctx *logger.TraceContext, input *spec.DeleteRegion, user *base_spec.UserAuthority) (*spec.DeleteRegionData, uint8, error)
	DeleteRegionService(tctx *logger.TraceContext, input *spec.DeleteRegionService, user *base_spec.UserAuthority) (*spec.DeleteRegionServiceData, uint8, error)
	DeleteRegionServices(tctx *logger.TraceContext, input *spec.DeleteRegionServices, user *base_spec.UserAuthority) (*spec.DeleteRegionServicesData, uint8, error)
	DeleteRegions(tctx *logger.TraceContext, input *spec.DeleteRegions, user *base_spec.UserAuthority) (*spec.DeleteRegionsData, uint8, error)
	GetCluster(tctx *logger.TraceContext, input *spec.GetCluster, user *base_spec.UserAuthority) (*spec.GetClusterData, uint8, error)
	GetClusters(tctx *logger.TraceContext, input *spec.GetClusters, user *base_spec.UserAuthority) (*spec.GetClustersData, uint8, error)
	GetCompute(tctx *logger.TraceContext, input *spec.GetCompute, user *base_spec.UserAuthority) (*spec.GetComputeData, uint8, error)
	GetComputeConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, user *base_spec.UserAuthority, conn *websocket.Conn) (*spec.GetComputeConsoleData, uint8, error)
	GetComputes(tctx *logger.TraceContext, input *spec.GetComputes, user *base_spec.UserAuthority) (*spec.GetComputesData, uint8, error)
	GetDatacenter(tctx *logger.TraceContext, input *spec.GetDatacenter, user *base_spec.UserAuthority) (*spec.GetDatacenterData, uint8, error)
	GetDatacenters(tctx *logger.TraceContext, input *spec.GetDatacenters, user *base_spec.UserAuthority) (*spec.GetDatacentersData, uint8, error)
	GetEventRule(tctx *logger.TraceContext, input *spec.GetEventRule, user *base_spec.UserAuthority) (*spec.GetEventRuleData, uint8, error)
	GetEventRules(tctx *logger.TraceContext, input *spec.GetEventRules, user *base_spec.UserAuthority) (*spec.GetEventRulesData, uint8, error)
	GetEvents(tctx *logger.TraceContext, input *spec.GetEvents, user *base_spec.UserAuthority) (*spec.GetEventsData, uint8, error)
	GetFloor(tctx *logger.TraceContext, input *spec.GetFloor, user *base_spec.UserAuthority) (*spec.GetFloorData, uint8, error)
	GetFloors(tctx *logger.TraceContext, input *spec.GetFloors, user *base_spec.UserAuthority) (*spec.GetFloorsData, uint8, error)
	GetImage(tctx *logger.TraceContext, input *spec.GetImage, user *base_spec.UserAuthority) (*spec.GetImageData, uint8, error)
	GetImages(tctx *logger.TraceContext, input *spec.GetImages, user *base_spec.UserAuthority) (*spec.GetImagesData, uint8, error)
	GetLogParams(tctx *logger.TraceContext, input *spec.GetLogParams, user *base_spec.UserAuthority) (*spec.GetLogParamsData, uint8, error)
	GetLogs(tctx *logger.TraceContext, input *spec.GetLogs, user *base_spec.UserAuthority) (*spec.GetLogsData, uint8, error)
	GetNetworkV4(tctx *logger.TraceContext, input *spec.GetNetworkV4, user *base_spec.UserAuthority) (*spec.GetNetworkV4Data, uint8, error)
	GetNetworkV4s(tctx *logger.TraceContext, input *spec.GetNetworkV4s, user *base_spec.UserAuthority) (*spec.GetNetworkV4sData, uint8, error)
	GetNode(tctx *logger.TraceContext, input *spec.GetNode, user *base_spec.UserAuthority) (*spec.GetNodeData, uint8, error)
	GetNodeMetrics(tctx *logger.TraceContext, input *spec.GetNodeMetrics, user *base_spec.UserAuthority) (*spec.GetNodeMetricsData, uint8, error)
	GetNodeServices(tctx *logger.TraceContext, input *spec.GetNodeServices, user *base_spec.UserAuthority) (*spec.GetNodeServicesData, uint8, error)
	GetNodes(tctx *logger.TraceContext, input *spec.GetNodes, user *base_spec.UserAuthority) (*spec.GetNodesData, uint8, error)
	GetPhysicalModel(tctx *logger.TraceContext, input *spec.GetPhysicalModel, user *base_spec.UserAuthority) (*spec.GetPhysicalModelData, uint8, error)
	GetPhysicalModels(tctx *logger.TraceContext, input *spec.GetPhysicalModels, user *base_spec.UserAuthority) (*spec.GetPhysicalModelsData, uint8, error)
	GetPhysicalResource(tctx *logger.TraceContext, input *spec.GetPhysicalResource, user *base_spec.UserAuthority) (*spec.GetPhysicalResourceData, uint8, error)
	GetPhysicalResources(tctx *logger.TraceContext, input *spec.GetPhysicalResources, user *base_spec.UserAuthority) (*spec.GetPhysicalResourcesData, uint8, error)
	GetRack(tctx *logger.TraceContext, input *spec.GetRack, user *base_spec.UserAuthority) (*spec.GetRackData, uint8, error)
	GetRacks(tctx *logger.TraceContext, input *spec.GetRacks, user *base_spec.UserAuthority) (*spec.GetRacksData, uint8, error)
	GetRegion(tctx *logger.TraceContext, input *spec.GetRegion, user *base_spec.UserAuthority) (*spec.GetRegionData, uint8, error)
	GetRegionService(tctx *logger.TraceContext, input *spec.GetRegionService, user *base_spec.UserAuthority) (*spec.GetRegionServiceData, uint8, error)
	GetRegionServices(tctx *logger.TraceContext, input *spec.GetRegionServices, user *base_spec.UserAuthority) (*spec.GetRegionServicesData, uint8, error)
	GetRegions(tctx *logger.TraceContext, input *spec.GetRegions, user *base_spec.UserAuthority) (*spec.GetRegionsData, uint8, error)
	GetStatistics(tctx *logger.TraceContext, input *spec.GetStatistics, user *base_spec.UserAuthority) (*spec.GetStatisticsData, uint8, error)
	GetTrace(tctx *logger.TraceContext, input *spec.GetTrace, user *base_spec.UserAuthority) (*spec.GetTraceData, uint8, error)
	UpdateCluster(tctx *logger.TraceContext, input *spec.UpdateCluster, user *base_spec.UserAuthority) (*spec.UpdateClusterData, uint8, error)
	UpdateDatacenter(tctx *logger.TraceContext, input *spec.UpdateDatacenter, user *base_spec.UserAuthority) (*spec.UpdateDatacenterData, uint8, error)
	UpdateEventRules(tctx *logger.TraceContext, input *spec.UpdateEventRules, user *base_spec.UserAuthority) (*spec.UpdateEventRulesData, uint8, error)
	UpdateFloor(tctx *logger.TraceContext, input *spec.UpdateFloor, user *base_spec.UserAuthority) (*spec.UpdateFloorData, uint8, error)
	UpdateImage(tctx *logger.TraceContext, input *spec.UpdateImage, user *base_spec.UserAuthority) (*spec.UpdateImageData, uint8, error)
	UpdateNetworkV4(tctx *logger.TraceContext, input *spec.UpdateNetworkV4, user *base_spec.UserAuthority) (*spec.UpdateNetworkV4Data, uint8, error)
	UpdatePhysicalModel(tctx *logger.TraceContext, input *spec.UpdatePhysicalModel, user *base_spec.UserAuthority) (*spec.UpdatePhysicalModelData, uint8, error)
	UpdatePhysicalResource(tctx *logger.TraceContext, input *spec.UpdatePhysicalResource, user *base_spec.UserAuthority) (*spec.UpdatePhysicalResourceData, uint8, error)
	UpdateRack(tctx *logger.TraceContext, input *spec.UpdateRack, user *base_spec.UserAuthority) (*spec.UpdateRackData, uint8, error)
	UpdateRegion(tctx *logger.TraceContext, input *spec.UpdateRegion, user *base_spec.UserAuthority) (*spec.UpdateRegionData, uint8, error)
	UpdateRegionService(tctx *logger.TraceContext, input *spec.UpdateRegionService, user *base_spec.UserAuthority) (*spec.UpdateRegionServiceData, uint8, error)
}

type QueryHandler struct {
	baseConf *base_config.Config
	appConf  *base_config.AppConfig
	resolver QueryResolver
}

func NewQueryHandler(baseConf *base_config.Config, appConf *base_config.AppConfig, resolver QueryResolver) *QueryHandler {
	return &QueryHandler{
		baseConf: baseConf,
		appConf:  appConf,
		resolver: resolver,
	}
}

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, httpReq *http.Request, rw http.ResponseWriter,
	req *base_protocol.Request, rep *base_protocol.Response) (err error) {
	for _, query := range req.Queries {
		switch query.Name {
		case "Login":
			var input base_spec.Login
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.Login(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    data.Token,
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,         // TODO Configurable
				Expires:  time.Now().Add(1 * time.Hour), // TODO Configurable
			}
			http.SetCookie(rw, &cookie)

		case "Logout":
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: base_const.CodeOk,
			}
			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    "",
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
			}
			http.SetCookie(rw, &cookie)

		case "LoginWithToken":
			var input base_spec.LoginWithToken
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.LoginWithToken(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

		case "UpdateService":
			var input base_spec.UpdateService
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.UpdateService(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetServiceIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetProjectServiceIndex":
			var input base_spec.GetServiceIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetProjectServiceIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetServiceDashboardIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap["GetServiceDashboardIndex"] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetProjectServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetProjectServiceDashboardIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap["GetServiceDashboardIndex"] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateCluster":
			var input spec.CreateCluster
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateCluster(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateDatacenter":
			var input spec.CreateDatacenter
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateDatacenter(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateEventRules":
			var input spec.CreateEventRules
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateEventRules(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateFloor":
			var input spec.CreateFloor
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateFloor(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateImage":
			var input spec.CreateImage
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateImage(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateNetworkV4":
			var input spec.CreateNetworkV4
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateNetworkV4(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreatePhysicalModel":
			var input spec.CreatePhysicalModel
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreatePhysicalModel(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreatePhysicalResource":
			var input spec.CreatePhysicalResource
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreatePhysicalResource(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateRack":
			var input spec.CreateRack
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateRack(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateRegion":
			var input spec.CreateRegion
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateRegion(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "CreateRegionService":
			var input spec.CreateRegionService
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.CreateRegionService(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteCluster":
			var input spec.DeleteCluster
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteCluster(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteClusters":
			var input spec.DeleteClusters
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteClusters(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteDatacenter":
			var input spec.DeleteDatacenter
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteDatacenter(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteDatacenters":
			var input spec.DeleteDatacenters
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteDatacenters(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteEventRules":
			var input spec.DeleteEventRules
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteEventRules(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteFloor":
			var input spec.DeleteFloor
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteFloor(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteFloors":
			var input spec.DeleteFloors
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteFloors(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteImage":
			var input spec.DeleteImage
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteImage(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteImages":
			var input spec.DeleteImages
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteImages(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteNetworkV4":
			var input spec.DeleteNetworkV4
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteNetworkV4(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteNetworkV4s":
			var input spec.DeleteNetworkV4s
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteNetworkV4s(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalModel":
			var input spec.DeletePhysicalModel
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalModel(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalModels":
			var input spec.DeletePhysicalModels
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalModels(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalResource":
			var input spec.DeletePhysicalResource
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalResource(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalResources":
			var input spec.DeletePhysicalResources
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalResources(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRack":
			var input spec.DeleteRack
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteRack(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRacks":
			var input spec.DeleteRacks
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteRacks(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegion":
			var input spec.DeleteRegion
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteRegion(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegionService":
			var input spec.DeleteRegionService
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteRegionService(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegionServices":
			var input spec.DeleteRegionServices
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteRegionServices(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegions":
			var input spec.DeleteRegions
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.DeleteRegions(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetCluster":
			var input spec.GetCluster
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetCluster(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetClusters":
			var input spec.GetClusters
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetClusters(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetCompute":
			var input spec.GetCompute
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetCompute(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetComputes":
			var input spec.GetComputes
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetComputes(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetDatacenter":
			var input spec.GetDatacenter
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetDatacenter(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetDatacenters":
			var input spec.GetDatacenters
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetDatacenters(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetEventRule":
			var input spec.GetEventRule
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetEventRule(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetEventRules":
			var input spec.GetEventRules
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetEventRules(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetEvents":
			var input spec.GetEvents
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetEvents(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetFloor":
			var input spec.GetFloor
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetFloor(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetFloors":
			var input spec.GetFloors
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetFloors(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetImage":
			var input spec.GetImage
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetImage(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetImages":
			var input spec.GetImages
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetImages(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetLogParams":
			var input spec.GetLogParams
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetLogParams(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetLogs":
			var input spec.GetLogs
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetLogs(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetNetworkV4":
			var input spec.GetNetworkV4
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetNetworkV4(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetNetworkV4s":
			var input spec.GetNetworkV4s
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetNetworkV4s(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetNode":
			var input spec.GetNode
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetNode(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetNodeMetrics":
			var input spec.GetNodeMetrics
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetNodeMetrics(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetNodeServices":
			var input spec.GetNodeServices
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetNodeServices(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetNodes":
			var input spec.GetNodes
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetNodes(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalModel":
			var input spec.GetPhysicalModel
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetPhysicalModel(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalModels":
			var input spec.GetPhysicalModels
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetPhysicalModels(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalResource":
			var input spec.GetPhysicalResource
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetPhysicalResource(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalResources":
			var input spec.GetPhysicalResources
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetPhysicalResources(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetRack":
			var input spec.GetRack
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetRack(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetRacks":
			var input spec.GetRacks
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetRacks(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetRegion":
			var input spec.GetRegion
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetRegion(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetRegionService":
			var input spec.GetRegionService
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetRegionService(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetRegionServices":
			var input spec.GetRegionServices
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetRegionServices(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetRegions":
			var input spec.GetRegions
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetRegions(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetStatistics":
			var input spec.GetStatistics
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetStatistics(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetTrace":
			var input spec.GetTrace
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetTrace(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateCluster":
			var input spec.UpdateCluster
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateCluster(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateDatacenter":
			var input spec.UpdateDatacenter
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateDatacenter(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateEventRules":
			var input spec.UpdateEventRules
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateEventRules(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateFloor":
			var input spec.UpdateFloor
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateFloor(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateImage":
			var input spec.UpdateImage
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateImage(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateNetworkV4":
			var input spec.UpdateNetworkV4
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateNetworkV4(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdatePhysicalModel":
			var input spec.UpdatePhysicalModel
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdatePhysicalModel(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdatePhysicalResource":
			var input spec.UpdatePhysicalResource
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdatePhysicalResource(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateRack":
			var input spec.UpdateRack
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateRack(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateRegion":
			var input spec.UpdateRegion
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateRegion(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "UpdateRegionService":
			var input spec.UpdateRegionService
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.UpdateRegionService(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return nil
}

func (handler *QueryHandler) ExecWs(tctx *logger.TraceContext, httpReq *http.Request, rw http.ResponseWriter,
	req *base_protocol.Request, rep *base_protocol.Response, conn *websocket.Conn) (err error) {
	for _, query := range req.Queries {
		switch query.Name {
		case "GetComputeConsole":
			var input spec.GetComputeConsole
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}
			data, code, tmpErr := handler.resolver.GetComputeConsole(tctx, &input, req.UserAuthority, conn)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return
}
