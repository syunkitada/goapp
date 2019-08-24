package genpkg

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type QueryResolver interface {
	IssueToken(tctx *logger.TraceContext, db *gorm.DB, input *spec.IssueToken) (*spec.IssueTokenData, error)
	UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateService) (*spec.UpdateServiceData, error)
	GetAllUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetAllUsers) (*spec.GetAllUsersData, error)
	GetUser(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUser) (*spec.GetUserData, error)
	GetUsers(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetUsers) (*spec.GetUsersData, error)
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

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, req *base_model.Request, rep *base_model.Reply) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "IssueToken":
			var input spec.IssueToken
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, err := handler.resolver.IssueToken(tctx, db, &input)
			rep.Data["IssueToken"] = data
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

			data, err := handler.resolver.UpdateService(tctx, db, &input)
			rep.Data["UpdateService"] = data
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

			data, err := handler.resolver.GetAllUsers(tctx, db, &input)
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

			data, err := handler.resolver.GetUser(tctx, db, &input)
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

			data, err := handler.resolver.GetUsers(tctx, db, &input)
			rep.Data["GetUsers"] = data
			return err
		}
	}
	return nil
}
