package ctl_admin

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/authproxy/model"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
	"os/exec"
)

func MigrateDatabase() error {
	db, err := gorm.Open("mysql", Conf.AuthproxyDatabase.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	glog.Info("Connected DB")

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.Project{})
	db.AutoMigrate(&model.ProjectRole{})
	db.AutoMigrate(&model.Service{})

	if err := model_api.CreateUser(Conf.Admin.Username, Conf.Admin.Password); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateProjectRole("admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateProjectRole("tenant"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateProject("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.AssignRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	userTenantServices := []string{"Wiki", "Chat", "Ticket"}
	userAdminServices := []string{"Datacenter"}
	projectTenantServices := []string{"Resource"}

	for _, userTenantService := range userTenantServices {
		if err := model_api.CreateService(userTenantService, "user"); err != nil {
			glog.Error(err)
			return err
		}

		if err := model_api.AssignService("tenant", userTenantService); err != nil {
			glog.Error(err)
			return err
		}

		if err := model_api.AssignService("admin", userTenantService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for _, userAdminService := range userAdminServices {
		if err := model_api.CreateService(userAdminService, "user"); err != nil {
			glog.Error(err)
			return err
		}

		if err := model_api.AssignService("admin", userAdminService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for _, projectTenantService := range projectTenantServices {
		if err := model_api.CreateService(projectTenantService, "project"); err != nil {
			glog.Error(err)
			return err
		}

		if err := model_api.AssignService("tenant", projectTenantService); err != nil {
			glog.Error(err)
			return err
		}

		if err := model_api.AssignService("admin", projectTenantService); err != nil {
			glog.Error(err)
			return err
		}
	}

	return nil
}

func RecreateTestDatabase() error {
	if err := DropTestDatabase(); err != nil {
		return err
	}

	if _, err := exec.Command("/usr/lib/mysql", "-u", "root", "-e", "create database testdb").Output(); err != nil {
		return err
	}

	return nil
}

func DropTestDatabase() error {
	if _, err := exec.Command("/usr/lib/mysql", "-u", "root", "-e", "drop database if exists testdb").Output(); err != nil {
		return err
	}

	return nil
}
