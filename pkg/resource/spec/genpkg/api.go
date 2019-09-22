// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceDashboardIndex) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	GetRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegion) (*spec.GetRegionData, uint8, error)
	GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (*spec.GetRegionsData, uint8, error)
	CreateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRegion) (*spec.CreateRegionData, uint8, error)
	UpdateRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateRegion) (*spec.UpdateRegionData, uint8, error)
	DeleteRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRegion) (*spec.DeleteRegionData, uint8, error)
	GetDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetDatacenter) (*spec.GetDatacenterData, uint8, error)
	GetDatacenters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetDatacenters) (*spec.GetDatacentersData, uint8, error)
	CreateDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateDatacenter) (*spec.CreateDatacenterData, uint8, error)
	UpdateDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateDatacenter) (*spec.UpdateDatacenterData, uint8, error)
	DeleteDatacenter(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteDatacenter) (*spec.DeleteDatacenterData, uint8, error)
	GetFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetFloor) (*spec.GetFloorData, uint8, error)
	GetFloors(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetFloors) (*spec.GetFloorsData, uint8, error)
	CreateFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateFloor) (*spec.CreateFloorData, uint8, error)
	UpdateFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateFloor) (*spec.UpdateFloorData, uint8, error)
	DeleteFloor(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteFloor) (*spec.DeleteFloorData, uint8, error)
	GetRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRack) (*spec.GetRackData, uint8, error)
	GetRacks(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRacks) (*spec.GetRacksData, uint8, error)
	CreateRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreateRack) (*spec.CreateRackData, uint8, error)
	UpdateRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateRack) (*spec.UpdateRackData, uint8, error)
	DeleteRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRack) (*spec.DeleteRackData, uint8, error)
	GetPhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalModel) (*spec.GetPhysicalModelData, uint8, error)
	GetPhysicalModels(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalModels) (*spec.GetPhysicalModelsData, uint8, error)
	CreatePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreatePhysicalModel) (*spec.CreatePhysicalModelData, uint8, error)
	UpdatePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdatePhysicalModel) (*spec.UpdatePhysicalModelData, uint8, error)
	DeletePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalModel) (*spec.DeletePhysicalModelData, uint8, error)
	GetPhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResource) (*spec.GetPhysicalResourceData, uint8, error)
	GetPhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResources) (*spec.GetPhysicalResourcesData, uint8, error)
	CreatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.CreatePhysicalResource) (*spec.CreatePhysicalResourceData, uint8, error)
	UpdatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdatePhysicalResource) (*spec.UpdatePhysicalResourceData, uint8, error)
	DeletePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalResource) (*spec.DeletePhysicalResourceData, uint8, error)
	GetClusters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetClusters) (*spec.GetClustersData, uint8, error)
}

type QueryHandler struct {
	baseConf *base_config.Config
	appConf  *base_config.AppConfig
	dbApi    base_db_api.IApi
	resolver QueryResolver
}

func NewQueryHandler(baseConf *base_config.Config, appConf *base_config.AppConfig, dbApi base_db_api.IApi, resolver QueryResolver) *QueryHandler {
	return &QueryHandler{
		baseConf: baseConf,
		appConf:  appConf,
		dbApi:    dbApi,
		resolver: resolver,
	}
}

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, db *gorm.DB, user *base_spec.UserAuthority, httpReq *http.Request, rw http.ResponseWriter,
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

			data, code, err := handler.resolver.Login(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["Login"] = data
			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    data.Token,
				Secure:   true,
				HttpOnly: true,
				Expires:  time.Now().Add(1 * time.Hour), // TODO Configurable
			} // FIXME SameSite
			http.SetCookie(rw, &cookie)
		case "Logout":
			rep.Code = base_const.CodeOk
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

			data, code, err := handler.resolver.LoginWithToken(tctx, db, &input, user)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["Login"] = data
		case "UpdateService":
			var input base_spec.UpdateService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdateService(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdateService"] = data
		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetServiceIndex(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetServiceIndex"] = data
		case "GetServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetServiceDashboardIndex(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetServiceDashboardIndex"] = data
		case "GetRegion":
			var input spec.GetRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetRegion(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetRegion"] = data
		case "GetRegions":
			var input spec.GetRegions
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetRegions(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetRegions"] = data
		case "CreateRegion":
			var input spec.CreateRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.CreateRegion(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["CreateRegion"] = data
		case "UpdateRegion":
			var input spec.UpdateRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdateRegion(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdateRegion"] = data
		case "DeleteRegion":
			var input spec.DeleteRegion
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.DeleteRegion(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["DeleteRegion"] = data
		case "GetDatacenter":
			var input spec.GetDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetDatacenter(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetDatacenter"] = data
		case "GetDatacenters":
			var input spec.GetDatacenters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetDatacenters(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetDatacenters"] = data
		case "CreateDatacenter":
			var input spec.CreateDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.CreateDatacenter(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["CreateDatacenter"] = data
		case "UpdateDatacenter":
			var input spec.UpdateDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdateDatacenter(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdateDatacenter"] = data
		case "DeleteDatacenter":
			var input spec.DeleteDatacenter
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.DeleteDatacenter(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["DeleteDatacenter"] = data
		case "GetFloor":
			var input spec.GetFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetFloor(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetFloor"] = data
		case "GetFloors":
			var input spec.GetFloors
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetFloors(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetFloors"] = data
		case "CreateFloor":
			var input spec.CreateFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.CreateFloor(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["CreateFloor"] = data
		case "UpdateFloor":
			var input spec.UpdateFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdateFloor(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdateFloor"] = data
		case "DeleteFloor":
			var input spec.DeleteFloor
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.DeleteFloor(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["DeleteFloor"] = data
		case "GetRack":
			var input spec.GetRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetRack(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetRack"] = data
		case "GetRacks":
			var input spec.GetRacks
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetRacks(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetRacks"] = data
		case "CreateRack":
			var input spec.CreateRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.CreateRack(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["CreateRack"] = data
		case "UpdateRack":
			var input spec.UpdateRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdateRack(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdateRack"] = data
		case "DeleteRack":
			var input spec.DeleteRack
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.DeleteRack(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["DeleteRack"] = data
		case "GetPhysicalModel":
			var input spec.GetPhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetPhysicalModel(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetPhysicalModel"] = data
		case "GetPhysicalModels":
			var input spec.GetPhysicalModels
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetPhysicalModels(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetPhysicalModels"] = data
		case "CreatePhysicalModel":
			var input spec.CreatePhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.CreatePhysicalModel(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["CreatePhysicalModel"] = data
		case "UpdatePhysicalModel":
			var input spec.UpdatePhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdatePhysicalModel(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdatePhysicalModel"] = data
		case "DeletePhysicalModel":
			var input spec.DeletePhysicalModel
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.DeletePhysicalModel(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["DeletePhysicalModel"] = data
		case "GetPhysicalResource":
			var input spec.GetPhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetPhysicalResource(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetPhysicalResource"] = data
		case "GetPhysicalResources":
			var input spec.GetPhysicalResources
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetPhysicalResources(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetPhysicalResources"] = data
		case "CreatePhysicalResource":
			var input spec.CreatePhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.CreatePhysicalResource(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["CreatePhysicalResource"] = data
		case "UpdatePhysicalResource":
			var input spec.UpdatePhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.UpdatePhysicalResource(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdatePhysicalResource"] = data
		case "DeletePhysicalResource":
			var input spec.DeletePhysicalResource
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.DeletePhysicalResource(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["DeletePhysicalResource"] = data
		case "GetClusters":
			var input spec.GetClusters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, err := handler.resolver.GetClusters(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetClusters"] = data
		}
	}
	return nil
}
