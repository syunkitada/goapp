package base_db_api

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type IApi interface {
	MustOpen()
	MustClose()
	GetUserWithValidatePassword(tctx *logger.TraceContext, name string, password string) (user *base_db_model.User, code uint8, err error)
	GetUserAuthority(tctx *logger.TraceContext, username string) (data *base_spec.UserAuthority, err error)
	CreateOrUpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (err error)
	GetServices(tctx *logger.TraceContext, input *base_spec.GetServices) (data *base_spec.GetServicesData, err error)
	CreateOrUpdateNode(tctx *logger.TraceContext, input *base_spec.UpdateNode) (err error)
	SyncNodeRole(tctx *logger.TraceContext, kind string) (nodes []base_db_model.Node, err error)
	LoginWithToken(tctx *logger.TraceContext, token string) (data *base_spec.UserAuthority, err error)
	IssueToken(userName string) (token string, err error)
}

type Api struct {
	baseConf             *base_config.Config
	appConf              *base_config.AppConfig
	DB                   *gorm.DB
	databaseConf         base_config.DatabaseConfig
	nodeDownTimeDuration time.Duration
	secrets              []string
	apiQueryMap          map[string]map[string]spec_model.QueryModel
}

func New(baseConf *base_config.Config, appConf *base_config.AppConfig, apiQueryMap map[string]map[string]spec_model.QueryModel) *Api {
	api := Api{
		baseConf:             baseConf,
		appConf:              appConf,
		nodeDownTimeDuration: time.Duration(appConf.NodeDownTimeDuration) * time.Second,
		databaseConf:         appConf.Database,
		secrets:              appConf.Auth.Secrets,
		apiQueryMap:          apiQueryMap,
	}

	return &api
}

func (api *Api) MustOpen() {
	db, tmpErr := gorm.Open("mysql", api.databaseConf.Connection)
	if tmpErr != nil {
		logger.StdoutFatalf("Failed Open: %v", tmpErr)
	}
	db.LogMode(api.baseConf.EnableDatabaseLog)
	api.DB = db
}

func (api *Api) MustClose() {
	if tmpErr := api.DB.Close(); tmpErr != nil {
		logger.StdoutFatalf("Failed Close: %v", tmpErr)
	}
}

func (api *Api) Transact(tctx *logger.TraceContext, txFunc func(tx *gorm.DB) (err error)) (err error) {
	tx := api.DB.Begin()
	if err = tx.Error; err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed rollback on recover")
			}
			err = error_utils.NewRecoveredError(p)
		} else if err != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				logger.Errorf(tctx, tmpErr, "Failed rollback on err")
			}
		} else {
			if err = tx.Commit().Error; err != nil {
				if tmpErr := tx.Rollback().Error; tmpErr != nil {
					logger.Errorf(tctx, tmpErr, "Failed rollback on commit")
				}
			}
		}
	}()
	err = txFunc(tx)
	return
}
