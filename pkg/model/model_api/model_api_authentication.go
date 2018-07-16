package model_api

import (
	"github.com/syunkitada/goapp/pkg/model"
	"github.com/syunkitada/goapp/pkg/util"

	_ "github.com/go-sql-driver/mysql"
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

func CreateRole(name string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var role model.Role

	if err := db.Debug().Where("name = ?", name).First(&role).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		role = model.Role{
			Name: name,
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

func CreateProject(name string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var project model.Project

	if err := db.Debug().Where("name = ?", name).First(&project).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		project = model.Project{
			Name: name,
		}
		db.Debug().Create(&project)

		return nil
	}

	return nil
}

func CreateProjectRole(name string, userName string, projectName string) error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	var user model.User
	var project model.Project
	var projectRole model.ProjectRole

	db.Debug().Where("name = ?", projectName).First(&project)

	if err := db.Debug().Where("name = ? and project_id = ?", name, project.ID).First(&projectRole).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		projectRole = model.ProjectRole{
			Name:      name,
			ProjectID: project.ID,
		}
		db.Debug().Create(&projectRole)
	}

	db.Debug().Preload("ProjectRoles").First(&user, "name = ?", userName)
	db.Debug().Model(&user).Association("ProjectRoles").Append(&projectRole)

	return nil
}

func IssueToken(authRequest *model.AuthRequest) (string, error) {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return "", err
	}

	var user model.User

	hashedPassword, hashedErr := util.GenerateHashFromPassword(authRequest.Username, authRequest.Password)
	if hashedErr != nil {
		return "", hashedErr
	}

	if err := db.Debug().Where("name = ? and password = ?", authRequest.Username, hashedPassword).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return "", nil
		}

		return "", err
	}

	return util.GenerateToken(authRequest)
}
