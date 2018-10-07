package model_api

import (
	"encoding/hex"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"

	"github.com/syunkitada/goapp/pkg/authproxy/model"
)

func (modelApi *ModelApi) CreateUser(name string, password string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var user model.User

	if err := db.Debug().Where("name = ?", name).First(&user).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		hashedPassword, hashedErr := modelApi.GenerateHashFromPassword(name, password)
		if hashedErr != nil {
			return hashedErr
		}

		user = model.User{
			Name:     name,
			Password: hashedPassword,
		}
		db.Debug().Create(&user)

		return nil
	}

	return nil
}

func (modelApi *ModelApi) CreateRole(name string, projectName string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var role model.Role
	var project model.Project

	if err := db.Debug().First(&project, "name = ?", projectName).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	if err := db.Debug().Where("name = ?", name).First(&role).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		role = model.Role{
			Name:      name,
			ProjectID: project.ID,
		}
		db.Debug().Create(&role)

		return nil
	}

	return nil
}

func (modelApi *ModelApi) AssignRole(userName string, roleName string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var user model.User
	var role model.Role

	db.Debug().Where("name = ?", roleName).First(&role)

	db.Debug().Preload("Roles").First(&user, "name = ?", userName)
	db.Debug().Model(&user).Association("Roles").Append(&role)
	return nil
}

func (modelApi *ModelApi) CreateProject(name string, projectRoleName string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var project model.Project
	var projectRole model.ProjectRole

	if err := db.Debug().First(&projectRole, "name = ?", projectRoleName).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	if err := db.Debug().Where("name = ?", name).First(&project).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		project = model.Project{
			Name:          name,
			ProjectRoleID: projectRole.ID,
		}
		db.Debug().Create(&project)

		return nil
	}

	return nil
}

func (modelApi *ModelApi) CreateProjectRole(name string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var projectRole model.ProjectRole

	if err := db.Debug().Where("name = ?", name).First(&projectRole).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		projectRole = model.ProjectRole{
			Name: name,
		}
		db.Debug().Create(&projectRole)

		return nil
	}

	return nil

}

func (modelApi *ModelApi) AssignProjectRole(projectName string, projectRoleName string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var project model.Project
	var projectRole model.ProjectRole

	db.Debug().Where("name = ?", projectRoleName).First(&projectRole)

	db.Debug().Preload("ProjectRoles").First(&project, "name = ?", projectName)
	db.Debug().Model(&project).Association("ProjectRoles").Append(&projectRole)
	return nil
}

func (modelApi *ModelApi) CreateService(name string, scope string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var service model.Service

	if err := db.Debug().Where("name = ?", name).First(&service).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		service = model.Service{
			Name:  name,
			Scope: scope,
		}
		db.Debug().Create(&service)

		return nil
	}

	return nil
}

func (modelApi *ModelApi) AssignService(projectRoleName string, serviceName string) error {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var projectRole model.ProjectRole
	var service model.Service

	db.Debug().Where("name = ?", serviceName).First(&service)

	db.Debug().Preload("Services").First(&projectRole, "name = ?", projectRoleName)
	db.Debug().Model(&projectRole).Association("Services").Append(&service)

	return nil
}

func (modelApi *ModelApi) AssignAction(serviceName string, projectRoleName string, roleName string, actionName string) error {
	db, dbErr := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}

	var action model.Action
	var service model.Service
	var projectRole model.ProjectRole
	var role model.Role
	var roleID uint

	db.Debug().Where("name = ?", serviceName).First(&service)
	db.Debug().Where("name = ?", projectRoleName).First(&projectRole)
	if roleName != "" {
		if err := db.Debug().Where("name = ?", roleName).First(&role).Error; err != nil {
			return err
		}
		roleID = role.ID
	} else {
		roleID = 0
	}

	if err := db.Debug().Where("name = ? and service_id = ? and project_role_id = ?",
		actionName, service.ID, projectRole.ID).First(&action).Error; err != nil {

		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		action = model.Action{
			Name:          actionName,
			ServiceID:     service.ID,
			ProjectRoleID: projectRole.ID,
			RoleID:        roleID,
		}
		db.Debug().Create(&action)
	}

	return nil
}

func (modelApi *ModelApi) GetAuthUser(authRequest *model.AuthRequest) (*model.User, error) {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var users []model.User
	if err := db.Debug().Where("name = ?", authRequest.Username).Find(&users).Error; err != nil {
		return nil, err
	}

	if len(users) != 1 {
		return nil, errors.New("Invalid User")
	}

	hashedPassword, hashedErr := modelApi.GenerateHashFromPassword(authRequest.Username, authRequest.Password)
	if hashedErr != nil {
		return nil, hashedErr
	}

	user := users[0]
	if user.Password != hashedPassword {
		return nil, errors.New("Invalid Password")
	}

	return &user, nil
}

func (modelApi *ModelApi) GenerateHashFromPassword(username string, password string) (string, error) {
	converted, err := scrypt.Key([]byte(password), []byte(modelApi.Conf.Admin.Secret+username), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(converted[:]), nil
}

func (modelApi *ModelApi) GetUserAuthority(username string, actionRequest *model.ActionRequest) (*model.UserAuthority, error) {
	db, err := gorm.Open("mysql", modelApi.Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}

	var users []model.CustomUser
	if err := db.Debug().Raw(sqlSelectUser+"WHERE u.name = ?", username).Scan(&users).Error; err != nil {
		return nil, err
	}

	serviceMap := map[string]uint{}
	projectServiceMap := map[string]model.ProjectService{}
	for _, user := range users {
		switch user.ServiceScope {
		case "user":
			serviceMap[user.ServiceName] = user.ServiceID
		case "project":
			glog.Info(user)
			if projectService, ok := projectServiceMap[user.ProjectName]; ok {
				projectService.ServiceMap[user.ServiceName] = user.ServiceID
			} else {
				projectService := model.ProjectService{
					RoleID:          user.RoleID,
					RoleName:        user.RoleName,
					ProjectName:     user.ProjectName,
					ProjectRoleID:   user.ProjectRoleID,
					ProjectRoleName: user.ProjectRoleName,
					ServiceMap:      map[string]uint{},
				}
				projectService.ServiceMap[user.ServiceName] = user.ServiceID
				projectServiceMap[user.ProjectName] = projectService
			}
		}
	}

	userAuthority := model.UserAuthority{
		ServiceMap:        serviceMap,
		ProjectServiceMap: projectServiceMap,
	}

	glog.Info("DEBUGLALALALALALA")
	if actionRequest != nil && actionRequest.ProjectName != "" && actionRequest.ServiceName != "" && actionRequest.Name != "" {
		projectService, projectServiceOk := projectServiceMap[actionRequest.ProjectName]
		if !projectServiceOk {
			return nil, errors.New(fmt.Sprintf("NotFound %v in projectServiceMap", actionRequest.ProjectName))
		}

		serviceID, serviceOk := projectService.ServiceMap[actionRequest.ServiceName]
		if !serviceOk {
			return nil, errors.New(fmt.Sprintf("NotFound %v in projectService.ServiceMap", actionRequest.ServiceName))
		}

		var action model.Action
		if err := db.Debug().Where("service_id = ? and name = ? and project_role_id = ?", serviceID, actionRequest.Name, projectService.ProjectRoleID).First(&action).Error; err != nil {
			return nil, err
		}

		if action.RoleID != 0 && action.RoleID != projectService.RoleID {
			return nil, errors.New(fmt.Sprintf("action.RoleID(%v) != projectService.RoleID(%v)",
				action.RoleID, projectService.RoleID))
		}

		userAuthority.ActionProjectService = projectService
	}

	return &userAuthority, nil
}
