package genpkg

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type QueryResolver interface {
	GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetServiceIndex) (*spec.GetServiceIndexData, uint8, error)
	UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateService) (*spec.UpdateServiceData, uint8, error)
	Login(tctx *logger.TraceContext, db *gorm.DB, input *spec.Login) (*spec.LoginData, uint8, error)
	GetAllUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetAllUsers) (*spec.GetAllUsersData, uint8, error)
	GetUser(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUser) (*spec.GetUserData, uint8, error)
	GetUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUsers) (*spec.GetUsersData, uint8, error)
}

type QueryHandler struct {
	resolver QueryResolver
	dbApi    *db_api.Api
}

func NewQueryHandler(baseConf *base_config.Config, mainConf *config.Config, resolver QueryResolver) *QueryHandler {
	return &QueryHandler{
		resolver: resolver,
		dbApi:    db_api.New(baseConf, mainConf),
	}
}

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, req *base_model.Request, rep *base_model.Response) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "GetServiceIndex":
			var input spec.GetServiceIndex
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
			return err
		case "UpdateService":
			var input spec.UpdateService
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
			return err
		case "Login":
			var input spec.Login
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
			return err
		case "GetAllUsers":
			var input spec.GetAllUsers
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
			return err
		case "GetUser":
			var input spec.GetUser
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
			return err
		case "GetUsers":
			var input spec.GetUsers
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
			return err
		}
	}
	return nil
}
