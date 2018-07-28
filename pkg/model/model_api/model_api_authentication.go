package model_api

import (
	"github.com/syunkitada/goapp/pkg/model"
	"github.com/syunkitada/goapp/pkg/util"

	"errors"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/golang/glog"
	"github.com/jinzhu/gorm"
)

func CreateUser(name string, password string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var user model.User

	if err := db.Debug().Where("name = ?", name).First(&user).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		hashedPassword, hashedErr := util.GenerateHashFromPassword(name, password)
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

func CreateRole(name string, projectName string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
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

func AssignRole(userName string, roleName string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
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

func CreateProject(name string, projectRoleName string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
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

func CreateProjectRole(name string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
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

func AssignProjectRole(projectName string, projectRoleName string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
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

func IssueToken(authRequest *model.AuthRequest) (string, error) {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return "", err
	}

	var users []model.CustomUser
	if err := db.Debug().Raw(sqlSelectUser+" WHERE u.name LIKE ?", authRequest.Username).Scan(&users).Error; err != nil {
		return "", err
	}

	if len(users) != 1 {
		return "", errors.New("Invalid User")
	}

	hashedPassword, hashedErr := util.GenerateHashFromPassword(authRequest.Username, authRequest.Password)
	if hashedErr != nil {
		return "", hashedErr
	}

	user := users[0]
	if user.Password != hashedPassword {
		return "", errors.New("Invalid Password")
	}

	return util.GenerateToken(&user)
}
