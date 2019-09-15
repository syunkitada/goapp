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
