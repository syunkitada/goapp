package base_db_api

import (
	"encoding/hex"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) GetUserWithValidatePassword(tctx *logger.TraceContext, db *gorm.DB, name string, password string) (user *base_db_model.User, code uint8, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var tmpUser base_db_model.User
	if err = db.Where("name = ?", name).First(&tmpUser).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeClientInvalidAuth
		} else {
			code = base_const.CodeRemoteError
		}
		return
	}

	hashedPassword, err := api.generateHashFromPassword(password)
	if err != nil {
		code = base_const.CodeClientInvalidAuth
		return
	}
	if tmpUser.Password != hashedPassword {
		code = base_const.CodeClientInvalidAuth
		err = error_utils.NewInvalidAuthError(name)
	}
	code = base_const.CodeOk
	user = &tmpUser
	return
}

func (api *Api) GetUserAuthority(tctx *logger.TraceContext, db *gorm.DB, username string) (*base_spec.UserAuthority, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var users []base_db_model.CustomUser
	query := "SELECT u.name, " +
		"r.id as role_id, r.name as role_name, " +
		"p.name as project_name, " +
		"pr.id as project_role_id, pr.name as project_role_name, " +
		"s.id as service_id, s.name as service_name, s.scope as service_scope " +
		"FROM users as u " +
		"INNER JOIN user_roles as ur ON u.id = ur.user_id " +
		"INNER JOIN roles as r ON ur.role_id = r.id " +
		"INNER JOIN projects as p ON r.project_id = p.id " +
		"INNER JOIN project_roles as pr ON p.project_role_id = pr.id " +
		"INNER JOIN project_role_services as prs ON pr.id = prs.project_role_id " +
		"INNER JOIN services as s ON prs.service_id = s.id "
	if err = db.Raw(query+"WHERE u.name = ?", username).Scan(&users).Error; err != nil {
		return nil, err
	}

	fmt.Println("DEBUG user", users)

	serviceMap := map[string]uint{}
	projectServiceMap := map[string]base_spec.ProjectService{}
	for _, user := range users {
		switch user.ServiceScope {
		case "user":
			serviceMap[user.ServiceName] = user.ServiceID
		case "project":
			if projectService, ok := projectServiceMap[user.ProjectName]; ok {
				projectService.ServiceMap[user.ServiceName] = user.ServiceID
			} else {
				projectService := base_spec.ProjectService{
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

	userAuthority := base_spec.UserAuthority{
		ServiceMap:        serviceMap,
		ProjectServiceMap: projectServiceMap,
	}

	return &userAuthority, nil
}

func (api *Api) CreateUser(tctx *logger.TraceContext, db *gorm.DB, name string, password string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = error_utils.NewRecoveredError(tmpErr)
		}
		api.Rollback(tctx, tx, err)
	}()

	var user base_db_model.User
	if err = tx.Where("name = ?", name).First(&user).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		var hashedPassword string
		hashedPassword, err = api.generateHashFromPassword(password)
		if err != nil {
			return err
		}

		user = base_db_model.User{
			Name:     name,
			Password: hashedPassword,
		}
		tx.Create(&user)
		err = tx.Commit().Error
	}
	return err
}

func (api *Api) generateHashFromPassword(password string) (string, error) {
	converted, err := scrypt.Key([]byte(password), []byte(api.secrets[0]), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(converted[:]), nil
}
