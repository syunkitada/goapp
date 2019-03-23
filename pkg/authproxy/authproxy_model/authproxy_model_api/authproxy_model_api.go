package authproxy_model_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type AuthproxyModelApi struct {
	conf *config.Config
}

func NewAuthproxyModelApi(conf *config.Config) *AuthproxyModelApi {
	modelApi := AuthproxyModelApi{
		conf: conf,
	}

	return &modelApi
}

func (modelApi *AuthproxyModelApi) Bootstrap(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	db.AutoMigrate(&authproxy_model.User{})
	db.AutoMigrate(&authproxy_model.Role{})
	db.AutoMigrate(&authproxy_model.Project{})
	db.AutoMigrate(&authproxy_model.ProjectRole{})
	db.AutoMigrate(&authproxy_model.Service{})
	db.AutoMigrate(&authproxy_model.Action{})

	if err = modelApi.CreateUser(tctx, modelApi.conf.Admin.Username, modelApi.conf.Admin.Password); err != nil {
		return err
	}

	if err = modelApi.CreateProjectRole(tctx, "admin"); err != nil {
		return err
	}

	if err = modelApi.CreateProjectRole(tctx, "tenant"); err != nil {
		return err
	}

	if err = modelApi.CreateProject(tctx, "admin", "admin"); err != nil {
		return err
	}

	if err = modelApi.CreateRole(tctx, "admin", "admin"); err != nil {
		return err
	}

	if err = modelApi.AssignRole(tctx, "admin", "admin"); err != nil {
		return err
	}

	userTenantServices := []string{"Wiki", "Chat", "Ticket", "Home"}
	userAdminServices := []string{"Datacenter", "Home"}
	projectTenantServices := []string{"Resource.Physical", "Resource.Virtual", "Monitor", "Home"}
	actionMap := map[string][]string{}
	actionMap["Resource.Physical"] = []string{"UserQuery"}
	actionMap["Resource.Virtual"] = []string{"UserQuery"}
	actionMap["Resource.Physical"] = []string{
		"UserQuery",
		"GetPhysicalIndex", "GetVirtualIndex",
		"CreatePhysicalResource", "UpdatePhysicalResource",
		"CreateVirtualResource", "UpdateVirtualResource",
		"GetState", "GetCluster", "GetNode",
		"GetNetwork", "CreateNetwork", "UpdateNetwork", "DeleteNetwork",
		"GetNetworkV4", "CreateNetworkV4", "UpdateNetworkV4", "DeleteNetworkV4",
		"GetNetworkV6", "CreateNetworkV6", "UpdateNetworkV6", "DeleteNetworkV6",
		"GetCompute", "CreateCompute", "UpdateCompute", "DeleteCompute",
		"GetContainer", "CreateContainer", "UpdateContainer", "DeleteContainer",
		"GetImage", "CreateImage", "UpdateImage", "DeleteImage",
		"GetVolume", "CreateVolume", "UpdateVolume", "DeleteVolume",
		"GetLoadbalancer", "CreateLoadbalancer", "UpdateLoadbalancer", "DeleteLoadbalancer",
	}
	actionMap["Monitor"] = []string{
		"GetState", "GetUserState", "GetIndexState", "GetNode", "GetIndex", "GetHost", "GetLog", "GetMetric",
		"GetIgnoreAlert", "CreateIgnoreAlert", "UpdateIgnoreAlert", "DeleteIgnoreAlert",
	}
	actionMap["Home"] = []string{"UserQuery"}

	for _, userTenantService := range userTenantServices {
		if err = modelApi.CreateService(tctx, userTenantService, "user"); err != nil {
			return err
		}

		if err = modelApi.AssignService(tctx, "tenant", userTenantService); err != nil {
			return err
		}

		if err = modelApi.AssignService(tctx, "admin", userTenantService); err != nil {
			return err
		}
	}

	for _, userAdminService := range userAdminServices {
		if err = modelApi.CreateService(tctx, userAdminService, "user"); err != nil {
			return err
		}

		if err = modelApi.AssignService(tctx, "admin", userAdminService); err != nil {
			return err
		}
	}

	for _, projectTenantService := range projectTenantServices {
		if err = modelApi.CreateService(tctx, projectTenantService, "project"); err != nil {
			return err
		}

		if err = modelApi.AssignService(tctx, "tenant", projectTenantService); err != nil {
			return err
		}

		if err = modelApi.AssignService(tctx, "admin", projectTenantService); err != nil {
			return err
		}
	}

	for serviceName, actions := range actionMap {
		for _, actionName := range actions {
			if err = modelApi.AssignAction(tctx, serviceName, "admin", "", actionName); err != nil {
				return err
			}

			if err = modelApi.AssignAction(tctx, serviceName, "tenant", "", actionName); err != nil {
				return err
			}
		}
	}

	return nil
}

func (modelApi *AuthproxyModelApi) open(tctx *logger.TraceContext) (*gorm.DB, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = gorm.Open("mysql", modelApi.conf.Authproxy.Database.Connection)
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	return db, nil
}
