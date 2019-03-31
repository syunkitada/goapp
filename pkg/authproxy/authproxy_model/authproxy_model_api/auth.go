package authproxy_model_api

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (modelApi *AuthproxyModelApi) CreateUser(tctx *logger.TraceContext, name string, password string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var user authproxy_model.User

	if err = db.Where("name = ?", name).First(&user).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		var hashedPassword string
		hashedPassword, err = modelApi.GenerateHashFromPassword(name, password)
		if err != nil {
			return err
		}

		user = authproxy_model.User{
			Name:     name,
			Password: hashedPassword,
		}
		db.Create(&user)

		return nil
	}

	return nil
}

func (modelApi *AuthproxyModelApi) CreateRole(tctx *logger.TraceContext, name string, projectName string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var role authproxy_model.Role
	var project authproxy_model.Project

	if err := db.First(&project, "name = ?", projectName).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	if err := db.Where("name = ?", name).First(&role).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		role = authproxy_model.Role{
			Name:      name,
			ProjectID: project.ID,
		}
		db.Create(&role)

		return nil
	}

	return nil
}

func (modelApi *AuthproxyModelApi) AssignRole(tctx *logger.TraceContext, userName string, roleName string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var user authproxy_model.User
	var role authproxy_model.Role

	db.Where("name = ?", roleName).First(&role)

	db.Preload("Roles").First(&user, "name = ?", userName)
	db.Model(&user).Association("Roles").Append(&role)
	return nil
}

func (modelApi *AuthproxyModelApi) CreateProject(tctx *logger.TraceContext, name string, projectRoleName string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var project authproxy_model.Project
	var projectRole authproxy_model.ProjectRole

	if err := db.First(&projectRole, "name = ?", projectRoleName).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	if err := db.Where("name = ?", name).First(&project).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		project = authproxy_model.Project{
			Name:          name,
			ProjectRoleID: projectRole.ID,
		}
		db.Create(&project)

		return nil
	}

	return nil
}

func (modelApi *AuthproxyModelApi) CreateProjectRole(tctx *logger.TraceContext, name string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var projectRole authproxy_model.ProjectRole

	if err := db.Where("name = ?", name).First(&projectRole).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		projectRole = authproxy_model.ProjectRole{
			Name: name,
		}
		db.Create(&projectRole)

		return nil
	}

	return nil

}

func (modelApi *AuthproxyModelApi) AssignProjectRole(tctx *logger.TraceContext, projectName string, projectRoleName string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var project authproxy_model.Project
	var projectRole authproxy_model.ProjectRole

	db.Where("name = ?", projectRoleName).First(&projectRole)

	db.Preload("ProjectRoles").First(&project, "name = ?", projectName)
	db.Model(&project).Association("ProjectRoles").Append(&projectRole)
	return nil
}

func (modelApi *AuthproxyModelApi) CreateService(tctx *logger.TraceContext, name string, scope string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var service authproxy_model.Service

	if err := db.Where("name = ?", name).First(&service).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		service = authproxy_model.Service{
			Name:  name,
			Scope: scope,
		}
		db.Create(&service)

		return nil
	}

	return nil
}

func (modelApi *AuthproxyModelApi) AssignService(tctx *logger.TraceContext, projectRoleName string, serviceName string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var projectRole authproxy_model.ProjectRole
	var service authproxy_model.Service

	db.Where("name = ?", serviceName).First(&service)

	db.Preload("Services").First(&projectRole, "name = ?", projectRoleName)
	db.Model(&projectRole).Association("Services").Append(&service)

	return nil
}

func (modelApi *AuthproxyModelApi) AssignAction(tctx *logger.TraceContext, serviceName string, projectRoleName string, roleName string, actionName string) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var action authproxy_model.Action
	var service authproxy_model.Service
	var projectRole authproxy_model.ProjectRole
	var role authproxy_model.Role
	var roleID uint

	db.Where("name = ?", serviceName).First(&service)
	db.Where("name = ?", projectRoleName).First(&projectRole)
	if roleName != "" {
		if err := db.Where("name = ?", roleName).First(&role).Error; err != nil {
			return err
		}
		roleID = role.ID
	} else {
		roleID = 0
	}

	if err := db.Where("name = ? and service_id = ? and project_role_id = ?",
		actionName, service.ID, projectRole.ID).First(&action).Error; err != nil {

		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		action = authproxy_model.Action{
			Name:          actionName,
			ServiceID:     service.ID,
			ProjectRoleID: projectRole.ID,
			RoleID:        roleID,
		}
		db.Create(&action)
	}

	return nil
}

func (modelApi *AuthproxyModelApi) GetAuthUser(tctx *logger.TraceContext, authRequest *authproxy_model.AuthRequest) (*authproxy_model.User, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return nil, err
	}
	defer func() { err = db.Close() }()

	var users []authproxy_model.User
	if err := db.Where("name = ?", authRequest.Username).Find(&users).Error; err != nil {
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

func (modelApi *AuthproxyModelApi) GenerateHashFromPassword(username string, password string) (string, error) {
	converted, err := scrypt.Key([]byte(password), []byte(modelApi.conf.Admin.Secret+username), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(converted[:]), nil
}

func (modelApi *AuthproxyModelApi) GetUserAuthority(tctx *logger.TraceContext, username string, actionRequest *authproxy_model.ActionRequest) (*authproxy_model.UserAuthority, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return nil, err
	}
	defer func() { err = db.Close() }()

	var users []authproxy_model.CustomUser
	if err := db.Raw(sqlSelectUser+"WHERE u.name = ?", username).Scan(&users).Error; err != nil {
		return nil, err
	}

	serviceMap := map[string]uint{}
	projectServiceMap := map[string]authproxy_model.ProjectService{}
	for _, user := range users {
		switch user.ServiceScope {
		case "user":
			serviceMap[user.ServiceName] = user.ServiceID
		case "project":
			if projectService, ok := projectServiceMap[user.ProjectName]; ok {
				projectService.ServiceMap[user.ServiceName] = user.ServiceID
			} else {
				projectService := authproxy_model.ProjectService{
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

	userAuthority := authproxy_model.UserAuthority{
		ServiceMap:        serviceMap,
		ProjectServiceMap: projectServiceMap,
	}

	if actionRequest != nil && actionRequest.ProjectName != "" && actionRequest.ServiceName != "" {
		projectService, projectServiceOk := projectServiceMap[actionRequest.ProjectName]
		if !projectServiceOk {
			return nil, fmt.Errorf("NotFound %v in projectServiceMap", actionRequest.ProjectName)
		}

		serviceID, serviceOk := projectService.ServiceMap[actionRequest.ServiceName]
		if !serviceOk {
			return nil, fmt.Errorf("NotFound %v in projectService.ServiceMap", actionRequest.ServiceName)
		}

		var actions []authproxy_model.Action
		if err := db.Where("service_id = ? and project_role_id = ?", serviceID, projectService.ProjectRoleID).Find(&actions).Error; err != nil {
			return nil, err
		}

		for _, query := range actionRequest.Queries {
			for _, action := range actions {
				if query.Kind == action.Name {
					if action.RoleID != 0 && action.RoleID != projectService.RoleID {
						return nil, fmt.Errorf("action.RoleID(%v) != projectService.RoleID(%v)",
							action.RoleID, projectService.RoleID)
					}
				}
			}
		}

		userAuthority.ActionProjectService = projectService
	}

	return &userAuthority, nil
}
