package genpkg

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type QueryResolver interface {
	GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (*base_spec.GetServiceIndexData, uint8, error)
	UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	Login(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	GetAllUsers(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetAllUsers) (*base_spec.GetAllUsersData, uint8, error)
	GetUser(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetUser) (*base_spec.GetUserData, uint8, error)
	GetUsers(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetUsers) (*base_spec.GetUsersData, uint8, error)
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

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, req *base_model.Request, rep *base_model.Response) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetServiceIndex(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetServiceIndex"] = data
		case "UpdateService":
			var input base_spec.UpdateService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.UpdateService(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["UpdateService"] = data
		case "Login":
			var input base_spec.Login
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.Login(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["Login"] = data
		case "GetAllUsers":
			var input base_spec.GetAllUsers
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetAllUsers(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetAllUsers"] = data
		case "GetUser":
			var input base_spec.GetUser
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetUser(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetUser"] = data
		case "GetUsers":
			var input base_spec.GetUsers
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetUsers(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetUsers"] = data
		}
	}
	return nil
}
