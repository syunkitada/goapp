package base_db_api

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) Bootstrap(tctx *logger.TraceContext, isRecreate bool) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	if err = exec_utils.CreateDatabase(tctx, api.baseConf, api.databaseConf.Connection, isRecreate); err != nil {
		return err
	}

	api.MustOpen()
	defer api.MustClose()

	if err = api.DB.AutoMigrate(&base_db_model.User{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&base_db_model.Role{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&base_db_model.Project{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&base_db_model.ProjectRole{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&base_db_model.Service{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&base_db_model.Node{}).Error; err != nil {
		return err
	}

	projectRoles := []string{}
	for _, projectRole := range api.appConf.Auth.DefaultProjectRoles {
		if err = api.CreateProjectRole(tctx, projectRole.Name); err != nil {
			return err
		}
		projectRoles = append(projectRoles, projectRole.Name)
		fmt.Printf("Created ProjectRole: %s\n", projectRole.Name)
	}

	for _, project := range api.appConf.Auth.DefaultProjects {
		if err = api.CreateProject(tctx, project.Name, project.ProjectRole); err != nil {
			return err
		}
		fmt.Printf("Created Project: %s\n", project.Name)
	}

	for _, role := range api.appConf.Auth.DefaultRoles {
		if err = api.CreateRole(tctx, role.Name, role.Project); err != nil {
			return err
		}
		fmt.Printf("Created Role: %s\n", role.Name)
	}

	for _, user := range api.appConf.Auth.DefaultUsers {
		if err = api.CreateUser(tctx, user.Name, user.Password); err != nil {
			return err
		}
		for _, roleName := range user.Roles {
			if err = api.AssignRoleToUser(tctx, roleName, user.Name); err != nil {
				return err
			}
		}

		fmt.Printf("Created User: %s\n", user.Name)
	}

	if err = api.CreateOrUpdateService(tctx, &base_spec.UpdateService{
		Name:            "Auth",
		Scope:           "user",
		SyncRootCluster: false,
		ProjectRoles:    projectRoles,
		Endpoints:       []string{},
		QueryMap: map[string]spec_model.QueryModel{
			"Login":          spec_model.QueryModel{},
			"LoginWithToken": spec_model.QueryModel{},
			"Logout":         spec_model.QueryModel{},
			"UpdateService":  spec_model.QueryModel{},
		},
	}); err != nil {
		return err
	}

	for _, service := range api.appConf.Auth.DefaultServices {
		queryMap, ok := api.apiQueryMap[service.Name]
		if !ok {
			fmt.Printf("Invalid service: querymap not found: %s\n", service.Name)
			continue
		}
		queryMap["GetServiceIndex"] = spec_model.QueryModel{}
		queryMap["GetServiceDashboardIndex"] = spec_model.QueryModel{}

		if err = api.CreateOrUpdateService(tctx, &base_spec.UpdateService{
			Name:            service.Name,
			Scope:           service.Scope,
			SyncRootCluster: service.SyncRootCluster,
			ProjectRoles:    service.ProjectRoles,
			Endpoints:       []string{},
			QueryMap:        queryMap,
		}); err != nil {
			return err
		}
	}

	return nil
}
