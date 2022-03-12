package base_db_api

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) GetUserWithValidatePassword(tctx *logger.TraceContext, name string, password string) (user *base_db_model.User, code uint8, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var tmpUser base_db_model.User
	if err = api.DB.Where("name = ?", name).First(&tmpUser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = error_utils.NewNotFoundError(name)
			code = base_const.CodeClientInvalidAuth
		} else {
			code = base_const.CodeRemoteError
		}
		return
	}

	var hashedPassword string
	if hashedPassword, err = api.generateHashFromPassword(password); err != nil {
		code = base_const.CodeClientInvalidAuth
		return
	}
	if tmpUser.Password != hashedPassword {
		code = base_const.CodeClientInvalidAuth
		err = error_utils.NewInvalidAuthError("Invalid password")
		return
	}
	code = base_const.CodeOk
	user = &tmpUser
	return
}

func (api *Api) GetUserAuthority(tctx *logger.TraceContext, username string) (userAuthority *base_spec.UserAuthority, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var users []base_db_model.CustomUser
	query := "SELECT u.name, " +
		"r.id as role_id, r.name as role_name, " +
		"p.name as project_name, " +
		"pr.id as project_role_id, pr.name as project_role_name, " +
		"s.id as service_id, s.name as service_name, s.icon as service_icon, s.scope as service_scope " +
		"FROM users as u " +
		"INNER JOIN user_roles as ur ON u.id = ur.user_id " +
		"INNER JOIN roles as r ON ur.role_id = r.id " +
		"INNER JOIN projects as p ON r.project_id = p.id " +
		"INNER JOIN project_roles as pr ON p.project_role_id = pr.id " +
		"INNER JOIN project_role_services as prs ON pr.id = prs.project_role_id " +
		"INNER JOIN services as s ON prs.service_id = s.id "
	if err = api.DB.Raw(query+"WHERE u.name = ?", username).Scan(&users).Error; err != nil {
		return
	}

	serviceMap := map[string]base_spec.ServiceData{}
	projectServiceMap := map[string]base_spec.ProjectService{}
	for _, user := range users {
		if user.ServiceName == "Auth" {
			continue
		}
		serviceData := base_spec.ServiceData{
			Id:   user.ServiceID,
			Icon: user.ServiceIcon,
		}
		switch user.ServiceScope {
		case "user":
			serviceMap[user.ServiceName] = serviceData
		case "project":
			if projectService, ok := projectServiceMap[user.ProjectName]; ok {
				projectService.ServiceMap[user.ServiceName] = serviceData
			} else {
				projectService := base_spec.ProjectService{
					RoleID:          user.RoleID,
					RoleName:        user.RoleName,
					ProjectName:     user.ProjectName,
					ProjectRoleID:   user.ProjectRoleID,
					ProjectRoleName: user.ProjectRoleName,
					ServiceMap:      map[string]base_spec.ServiceData{},
				}
				projectService.ServiceMap[user.ServiceName] = serviceData
				projectServiceMap[user.ProjectName] = projectService
			}
		}
	}

	userAuthority = &base_spec.UserAuthority{
		Name:              username,
		ServiceMap:        serviceMap,
		ProjectServiceMap: projectServiceMap,
	}
	return
}

func (api *Api) CreateUser(tctx *logger.TraceContext, name string, password string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var user base_db_model.User
		if err = tx.Where("name = ?", name).First(&user).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}

			var hashedPassword string
			hashedPassword, err = api.generateHashFromPassword(password)
			if err != nil {
				return
			}

			user = base_db_model.User{
				Name:     name,
				Password: hashedPassword,
			}
			err = tx.Create(&user).Error
		}
		return
	})
	return
}

func (api *Api) UpdateUserPassword(tctx *logger.TraceContext, name string, password string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var hashedPassword string
		hashedPassword, err = api.generateHashFromPassword(password)
		if err != nil {
			return
		}

		err = tx.Table("users").Where("name = ?", name).Update("password", hashedPassword).Error
		fmt.Println("DEBUG password", hashedPassword, err)
		return
	})
	return
}

func (api *Api) generateHashFromPassword(password string) (hash string, err error) {
	var converted []byte
	if converted, err = scrypt.Key([]byte(password), []byte(api.secrets[0]), 16384, 8, 1, 32); err != nil {
		return
	}
	hash = hex.EncodeToString(converted[:])
	return
}

func (api *Api) IssueToken(userName string) (token string, err error) {
	claims := base_protocol.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    userName,
		},
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with key
	token, err = newToken.SignedString([]byte(api.secrets[0]))
	return
}

func (api *Api) LoginWithToken(tctx *logger.TraceContext, token string) (data *base_spec.UserAuthority, err error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			return nil, msg
		}
		return []byte(api.secrets[0]), nil
	})

	if err != nil {
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		err = error_utils.NewInvalidAuthError("Token is ng")
		return
	}
	if !parsedToken.Valid {
		err = error_utils.NewInvalidAuthError("Invalid Token")
		return
	}

	issuer := claims["iss"].(string)
	data, err = api.GetUserAuthority(tctx, issuer)
	return
}

func (api *Api) GetUsers(tctx *logger.TraceContext, projectName string) (users []base_db_model.CustomUser, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	query := api.DB.Table("users as u").
		Select("u.name, r.id as role_id, r.name as role_name, p.name as project_name").
		Joins("INNER JOIN user_roles as ur ON u.id = ur.user_id").
		Joins("INNER JOIN roles as r ON ur.role_id = r.id").
		Joins("INNER JOIN projects as p ON r.project_id = p.id")
	if projectName != "" {
		query.Where("p.name = ?", projectName)
	}
	err = query.Scan(&users).Error
	return
}
