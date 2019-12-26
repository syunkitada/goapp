// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceDashboardIndexData, uint8, error)
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

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, user *base_spec.UserAuthority, httpReq *http.Request, rw http.ResponseWriter,
	req *base_model.Request, rep *base_model.Response) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "Login":
			var input base_spec.Login
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.Login(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    data.Token,
				Secure:   true,
				HttpOnly: true,
				Expires:  time.Now().Add(1 * time.Hour), // TODO Configurable
			} // FIXME SameSite
			http.SetCookie(rw, &cookie)

		case "Logout":
			rep.ResultMap[query.Name] = base_model.Result{
				Code: base_const.CodeOk,
			}
			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    "",
				Secure:   true,
				HttpOnly: true,
			}
			http.SetCookie(rw, &cookie)

		case "LoginWithToken":
			var input base_spec.LoginWithToken
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.LoginWithToken(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

		case "UpdateService":
			var input base_spec.UpdateService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.UpdateService(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetServiceIndex(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetServiceDashboardIndex(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap["GetServiceDashboardIndex"] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateCluster":
			var input spec.CreateCluster
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateCluster(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateDatacenter":
			var input spec.CreateDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateDatacenter(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateEventRules":
			var input spec.CreateEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateFloor":
			var input spec.CreateFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateFloor(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateImage":
			var input spec.CreateImage
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateImage(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateNetworkV4":
			var input spec.CreateNetworkV4
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateNetworkV4(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreatePhysicalModel":
			var input spec.CreatePhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreatePhysicalModel(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreatePhysicalResource":
			var input spec.CreatePhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreatePhysicalResource(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateRack":
			var input spec.CreateRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateRack(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateRegion":
			var input spec.CreateRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateRegion(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateRegionService":
			var input spec.CreateRegionService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateRegionService(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteCluster":
			var input spec.DeleteCluster
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteCluster(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteClusters":
			var input spec.DeleteClusters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteClusters(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteDatacenter":
			var input spec.DeleteDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteDatacenter(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteDatacenters":
			var input spec.DeleteDatacenters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteDatacenters(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteEventRules":
			var input spec.DeleteEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteFloor":
			var input spec.DeleteFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteFloor(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteFloors":
			var input spec.DeleteFloors
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteFloors(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteImage":
			var input spec.DeleteImage
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteImage(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteImages":
			var input spec.DeleteImages
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteImages(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteNetworkV4":
			var input spec.DeleteNetworkV4
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteNetworkV4(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteNetworkV4s":
			var input spec.DeleteNetworkV4s
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteNetworkV4s(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalModel":
			var input spec.DeletePhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalModel(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalModels":
			var input spec.DeletePhysicalModels
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalModels(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalResource":
			var input spec.DeletePhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalResource(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeletePhysicalResources":
			var input spec.DeletePhysicalResources
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeletePhysicalResources(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRack":
			var input spec.DeleteRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteRack(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRacks":
			var input spec.DeleteRacks
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteRacks(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegion":
			var input spec.DeleteRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteRegion(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegionService":
			var input spec.DeleteRegionService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteRegionService(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegionServices":
			var input spec.DeleteRegionServices
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteRegionServices(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteRegions":
			var input spec.DeleteRegions
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteRegions(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetCluster":
			var input spec.GetCluster
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetCluster(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetClusters":
			var input spec.GetClusters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetClusters(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetDatacenter":
			var input spec.GetDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetDatacenter(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetDatacenters":
			var input spec.GetDatacenters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetDatacenters(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetEventRule":
			var input spec.GetEventRule
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetEventRule(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetEventRules":
			var input spec.GetEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetEvents":
			var input spec.GetEvents
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetEvents(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetFloor":
			var input spec.GetFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetFloor(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetFloors":
			var input spec.GetFloors
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetFloors(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetImage":
			var input spec.GetImage
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetImage(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetImages":
			var input spec.GetImages
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetImages(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetLogParams":
			var input spec.GetLogParams
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetLogParams(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetLogs":
			var input spec.GetLogs
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetLogs(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNetworkV4":
			var input spec.GetNetworkV4
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNetworkV4(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNetworkV4s":
			var input spec.GetNetworkV4s
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNetworkV4s(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNode":
			var input spec.GetNode
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNode(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNodeServices":
			var input spec.GetNodeServices
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNodeServices(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNodes":
			var input spec.GetNodes
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNodes(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalModel":
			var input spec.GetPhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetPhysicalModel(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalModels":
			var input spec.GetPhysicalModels
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetPhysicalModels(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalResource":
			var input spec.GetPhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetPhysicalResource(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetPhysicalResources":
			var input spec.GetPhysicalResources
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetPhysicalResources(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetRack":
			var input spec.GetRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetRack(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetRacks":
			var input spec.GetRacks
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetRacks(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetRegion":
			var input spec.GetRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetRegion(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetRegionService":
			var input spec.GetRegionService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetRegionService(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetRegionServices":
			var input spec.GetRegionServices
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetRegionServices(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetRegions":
			var input spec.GetRegions
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetRegions(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetStatistics":
			var input spec.GetStatistics
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetStatistics(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetTrace":
			var input spec.GetTrace
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetTrace(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateCluster":
			var input spec.UpdateCluster
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateCluster(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateDatacenter":
			var input spec.UpdateDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateDatacenter(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateEventRules":
			var input spec.UpdateEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateFloor":
			var input spec.UpdateFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateFloor(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateImage":
			var input spec.UpdateImage
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateImage(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateNetworkV4":
			var input spec.UpdateNetworkV4
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateNetworkV4(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdatePhysicalModel":
			var input spec.UpdatePhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdatePhysicalModel(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdatePhysicalResource":
			var input spec.UpdatePhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdatePhysicalResource(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateRack":
			var input spec.UpdateRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateRack(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateRegion":
			var input spec.UpdateRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateRegion(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateRegionService":
			var input spec.UpdateRegionService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateRegionService(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
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
