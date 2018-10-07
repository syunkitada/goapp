package ctl_admin

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/authproxy/model"
	"os/exec"
)

func (adminCtl *AdminCtl) MigrateDatabase() error {
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
	db.AutoMigrate(&model.Action{})

	if err := adminCtl.ModelApi.CreateUser(adminCtl.Conf.Admin.Username, adminCtl.Conf.Admin.Password); err != nil {
		glog.Error(err)
		return err
	}

	if err := adminCtl.ModelApi.CreateProjectRole("admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := adminCtl.ModelApi.CreateProjectRole("tenant"); err != nil {
		glog.Error(err)
		return err
	}

	if err := adminCtl.ModelApi.CreateProject("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := adminCtl.ModelApi.CreateRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := adminCtl.ModelApi.AssignRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	userTenantServices := []string{"Wiki", "Chat", "Ticket"}
	userAdminServices := []string{"Datacenter"}
	projectTenantServices := []string{"Resource"}
	actionMap := map[string][]string{}
	actionMap["Resource"] = []string{"GetState"}

	for _, userTenantService := range userTenantServices {
		if err := adminCtl.ModelApi.CreateService(userTenantService, "user"); err != nil {
			glog.Error(err)
			return err
		}

		if err := adminCtl.ModelApi.AssignService("tenant", userTenantService); err != nil {
			glog.Error(err)
			return err
		}

		if err := adminCtl.ModelApi.AssignService("admin", userTenantService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for _, userAdminService := range userAdminServices {
		if err := adminCtl.ModelApi.CreateService(userAdminService, "user"); err != nil {
			glog.Error(err)
			return err
		}

		if err := adminCtl.ModelApi.AssignService("admin", userAdminService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for _, projectTenantService := range projectTenantServices {
		if err := adminCtl.ModelApi.CreateService(projectTenantService, "project"); err != nil {
			glog.Error(err)
			return err
		}

		if err := adminCtl.ModelApi.AssignService("tenant", projectTenantService); err != nil {
			glog.Error(err)
			return err
		}

		if err := adminCtl.ModelApi.AssignService("admin", projectTenantService); err != nil {
			glog.Error(err)
			return err
		}
	}

	for serviceName, actions := range actionMap {
		for _, actionName := range actions {
			if err := adminCtl.ModelApi.AssignAction(serviceName, "admin", "", actionName); err != nil {
				glog.Error(err)
				return err
			}

			if err := adminCtl.ModelApi.AssignAction(serviceName, "tenant", "", actionName); err != nil {
				glog.Error(err)
				return err
			}
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
