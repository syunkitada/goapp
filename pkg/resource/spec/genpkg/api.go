// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	CreateCluster(tctx *logger.TraceContext, input *spec.CreateCluster) (*spec.CreateClusterData, uint8, error)
	CreateDatacenter(tctx *logger.TraceContext, input *spec.CreateDatacenter) (*spec.CreateDatacenterData, uint8, error)
	CreateFloor(tctx *logger.TraceContext, input *spec.CreateFloor) (*spec.CreateFloorData, uint8, error)
	CreateImage(tctx *logger.TraceContext, input *spec.CreateImage) (*spec.CreateImageData, uint8, error)
	CreateNetworkV4(tctx *logger.TraceContext, input *spec.CreateNetworkV4) (*spec.CreateNetworkV4Data, uint8, error)
	CreatePhysicalModel(tctx *logger.TraceContext, input *spec.CreatePhysicalModel) (*spec.CreatePhysicalModelData, uint8, error)
	CreatePhysicalResource(tctx *logger.TraceContext, input *spec.CreatePhysicalResource) (*spec.CreatePhysicalResourceData, uint8, error)
	CreateRack(tctx *logger.TraceContext, input *spec.CreateRack) (*spec.CreateRackData, uint8, error)
	CreateRegion(tctx *logger.TraceContext, input *spec.CreateRegion) (*spec.CreateRegionData, uint8, error)
	CreateRegionService(tctx *logger.TraceContext, input *spec.CreateRegionService) (*spec.CreateRegionServiceData, uint8, error)
	DeleteCluster(tctx *logger.TraceContext, input *spec.DeleteCluster) (*spec.DeleteClusterData, uint8, error)
	DeleteClusters(tctx *logger.TraceContext, input *spec.DeleteClusters) (*spec.DeleteClustersData, uint8, error)
	DeleteDatacenter(tctx *logger.TraceContext, input *spec.DeleteDatacenter) (*spec.DeleteDatacenterData, uint8, error)
	DeleteDatacenters(tctx *logger.TraceContext, input *spec.DeleteDatacenters) (*spec.DeleteDatacentersData, uint8, error)
	DeleteFloor(tctx *logger.TraceContext, input *spec.DeleteFloor) (*spec.DeleteFloorData, uint8, error)
	DeleteFloors(tctx *logger.TraceContext, input *spec.DeleteFloors) (*spec.DeleteFloorsData, uint8, error)
	DeleteImage(tctx *logger.TraceContext, input *spec.DeleteImage) (*spec.DeleteImageData, uint8, error)
	DeleteImages(tctx *logger.TraceContext, input *spec.DeleteImages) (*spec.DeleteImagesData, uint8, error)
	DeleteNetworkV4(tctx *logger.TraceContext, input *spec.DeleteNetworkV4) (*spec.DeleteNetworkV4Data, uint8, error)
	DeleteNetworkV4s(tctx *logger.TraceContext, input *spec.DeleteNetworkV4s) (*spec.DeleteNetworkV4sData, uint8, error)
	DeletePhysicalModel(tctx *logger.TraceContext, input *spec.DeletePhysicalModel) (*spec.DeletePhysicalModelData, uint8, error)
	DeletePhysicalModels(tctx *logger.TraceContext, input *spec.DeletePhysicalModels) (*spec.DeletePhysicalModelsData, uint8, error)
	DeletePhysicalResource(tctx *logger.TraceContext, input *spec.DeletePhysicalResource) (*spec.DeletePhysicalResourceData, uint8, error)
	DeletePhysicalResources(tctx *logger.TraceContext, input *spec.DeletePhysicalResources) (*spec.DeletePhysicalResourcesData, uint8, error)
	DeleteRack(tctx *logger.TraceContext, input *spec.DeleteRack) (*spec.DeleteRackData, uint8, error)
	DeleteRacks(tctx *logger.TraceContext, input *spec.DeleteRacks) (*spec.DeleteRacksData, uint8, error)
	DeleteRegion(tctx *logger.TraceContext, input *spec.DeleteRegion) (*spec.DeleteRegionData, uint8, error)
	DeleteRegionService(tctx *logger.TraceContext, input *spec.DeleteRegionService) (*spec.DeleteRegionServiceData, uint8, error)
	DeleteRegionServices(tctx *logger.TraceContext, input *spec.DeleteRegionServices) (*spec.DeleteRegionServicesData, uint8, error)
	DeleteRegions(tctx *logger.TraceContext, input *spec.DeleteRegions) (*spec.DeleteRegionsData, uint8, error)
	GetCluster(tctx *logger.TraceContext, input *spec.GetCluster) (*spec.GetClusterData, uint8, error)
	GetClusters(tctx *logger.TraceContext, input *spec.GetClusters) (*spec.GetClustersData, uint8, error)
	GetDatacenter(tctx *logger.TraceContext, input *spec.GetDatacenter) (*spec.GetDatacenterData, uint8, error)
	GetDatacenters(tctx *logger.TraceContext, input *spec.GetDatacenters) (*spec.GetDatacentersData, uint8, error)
	GetFloor(tctx *logger.TraceContext, input *spec.GetFloor) (*spec.GetFloorData, uint8, error)
	GetFloors(tctx *logger.TraceContext, input *spec.GetFloors) (*spec.GetFloorsData, uint8, error)
	GetImage(tctx *logger.TraceContext, input *spec.GetImage) (*spec.GetImageData, uint8, error)
	GetImages(tctx *logger.TraceContext, input *spec.GetImages) (*spec.GetImagesData, uint8, error)
	GetNetworkV4(tctx *logger.TraceContext, input *spec.GetNetworkV4) (*spec.GetNetworkV4Data, uint8, error)
	GetNetworkV4s(tctx *logger.TraceContext, input *spec.GetNetworkV4s) (*spec.GetNetworkV4sData, uint8, error)
	GetPhysicalModel(tctx *logger.TraceContext, input *spec.GetPhysicalModel) (*spec.GetPhysicalModelData, uint8, error)
	GetPhysicalModels(tctx *logger.TraceContext, input *spec.GetPhysicalModels) (*spec.GetPhysicalModelsData, uint8, error)
	GetPhysicalResource(tctx *logger.TraceContext, input *spec.GetPhysicalResource) (*spec.GetPhysicalResourceData, uint8, error)
	GetPhysicalResources(tctx *logger.TraceContext, input *spec.GetPhysicalResources) (*spec.GetPhysicalResourcesData, uint8, error)
	GetRack(tctx *logger.TraceContext, input *spec.GetRack) (*spec.GetRackData, uint8, error)
	GetRacks(tctx *logger.TraceContext, input *spec.GetRacks) (*spec.GetRacksData, uint8, error)
	GetRegion(tctx *logger.TraceContext, input *spec.GetRegion) (*spec.GetRegionData, uint8, error)
	GetRegionService(tctx *logger.TraceContext, input *spec.GetRegionService) (*spec.GetRegionServiceData, uint8, error)
	GetRegionServices(tctx *logger.TraceContext, input *spec.GetRegionServices) (*spec.GetRegionServicesData, uint8, error)
	GetRegions(tctx *logger.TraceContext, input *spec.GetRegions) (*spec.GetRegionsData, uint8, error)
	UpdateCluster(tctx *logger.TraceContext, input *spec.UpdateCluster) (*spec.UpdateClusterData, uint8, error)
	UpdateDatacenter(tctx *logger.TraceContext, input *spec.UpdateDatacenter) (*spec.UpdateDatacenterData, uint8, error)
	UpdateFloor(tctx *logger.TraceContext, input *spec.UpdateFloor) (*spec.UpdateFloorData, uint8, error)
	UpdateImage(tctx *logger.TraceContext, input *spec.UpdateImage) (*spec.UpdateImageData, uint8, error)
	UpdateNetworkV4(tctx *logger.TraceContext, input *spec.UpdateNetworkV4) (*spec.UpdateNetworkV4Data, uint8, error)
	UpdatePhysicalModel(tctx *logger.TraceContext, input *spec.UpdatePhysicalModel) (*spec.UpdatePhysicalModelData, uint8, error)
	UpdatePhysicalResource(tctx *logger.TraceContext, input *spec.UpdatePhysicalResource) (*spec.UpdatePhysicalResourceData, uint8, error)
	UpdateRack(tctx *logger.TraceContext, input *spec.UpdateRack) (*spec.UpdateRackData, uint8, error)
	UpdateRegion(tctx *logger.TraceContext, input *spec.UpdateRegion) (*spec.UpdateRegionData, uint8, error)
	UpdateRegionService(tctx *logger.TraceContext, input *spec.UpdateRegionService) (*spec.UpdateRegionServiceData, uint8, error)
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

			data, code, tmpErr := handler.resolver.GetServiceIndex(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetServiceDashboardIndex(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateCluster(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateDatacenter(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateFloor(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateImage(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateNetworkV4(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreatePhysicalModel(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreatePhysicalResource(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateRack(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateRegion(tctx, &input)
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

			data, code, tmpErr := handler.resolver.CreateRegionService(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteCluster(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteClusters(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteDatacenter(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteDatacenters(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteFloor(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteFloors(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteImage(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteImages(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteNetworkV4(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteNetworkV4s(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeletePhysicalModel(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeletePhysicalModels(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeletePhysicalResource(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeletePhysicalResources(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteRack(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteRacks(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteRegion(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteRegionService(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteRegionServices(tctx, &input)
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

			data, code, tmpErr := handler.resolver.DeleteRegions(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetCluster(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetClusters(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetDatacenter(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetDatacenters(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetFloor(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetFloors(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetImage(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetImages(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetNetworkV4(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetNetworkV4s(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetPhysicalModel(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetPhysicalModels(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetPhysicalResource(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetPhysicalResources(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetRack(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetRacks(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetRegion(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetRegionService(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetRegionServices(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetRegions(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateCluster(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateDatacenter(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateFloor(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateImage(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateNetworkV4(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdatePhysicalModel(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdatePhysicalResource(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateRack(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateRegion(tctx, &input)
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

			data, code, tmpErr := handler.resolver.UpdateRegionService(tctx, &input)
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
		}
	}
	return nil
}
