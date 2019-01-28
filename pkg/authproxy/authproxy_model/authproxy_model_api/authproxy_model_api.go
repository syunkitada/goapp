package authproxy_model_api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/config"
)

type AuthproxyModelApi struct {
	Conf *config.Config
}

func NewAuthproxyModelApi(conf *config.Config) *AuthproxyModelApi {
	modelApi := AuthproxyModelApi{
		Conf: conf,
	}

	return &modelApi
}

func (modelApi *AuthproxyModelApi) Bootstrap() error {
	db, dbErr := gorm.Open("mysql", modelApi.Conf.Authproxy.Database.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}
	db.LogMode(modelApi.Conf.Default.EnableDatabaseLog)

	db.AutoMigrate(&authproxy_model.User{})
	db.AutoMigrate(&authproxy_model.Role{})
	db.AutoMigrate(&authproxy_model.Project{})
	db.AutoMigrate(&authproxy_model.ProjectRole{})
	db.AutoMigrate(&authproxy_model.Service{})
	db.AutoMigrate(&authproxy_model.Action{})

	if err := modelApi.CreateUser(modelApi.Conf.Admin.Username, modelApi.Conf.Admin.Password); err != nil {
		glog.Error(err)
		return err
	}

	if err := modelApi.CreateProjectRole("admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := modelApi.CreateProjectRole("tenant"); err != nil {
		glog.Error(err)
		return err
	}

	if err := modelApi.CreateProject("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := modelApi.CreateRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := modelApi.AssignRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	userTenantServices := []string{"Wiki", "Chat", "Ticket"}
	userAdminServices := []string{"Datacenter"}
	projectTenantServices := []string{"Resource", "Monitor"}
	actionMap := map[string][]string{}
	actionMap["Resource"] = []string{
		"GetState", "GetCluster", "GetNode",
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

	for _, userTenantService := range userTenantServices {
		if err := modelApi.CreateService(userTenantService, "user"); err != nil {
			glog.Error(err)
			return err
		}

		if err := modelApi.AssignService("tenant", userTenantService); err != nil {
			glog.Error(err)
			return err
		}

		if err := modelApi.AssignService("admin", userTenantService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for _, userAdminService := range userAdminServices {
		if err := modelApi.CreateService(userAdminService, "user"); err != nil {
			glog.Error(err)
			return err
		}

		if err := modelApi.AssignService("admin", userAdminService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for _, projectTenantService := range projectTenantServices {
		if err := modelApi.CreateService(projectTenantService, "project"); err != nil {
			glog.Error(err)
			return err
		}

		if err := modelApi.AssignService("tenant", projectTenantService); err != nil {
			glog.Error(err)
			return err
		}

		if err := modelApi.AssignService("admin", projectTenantService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for serviceName, actions := range actionMap {
		for _, actionName := range actions {
			if err := modelApi.AssignAction(serviceName, "admin", "", actionName); err != nil {
				glog.Error(err)
				return err
			}

			if err := modelApi.AssignAction(serviceName, "tenant", "", actionName); err != nil {
				glog.Error(err)
				return err
			}
		}
	}

	return nil
}
