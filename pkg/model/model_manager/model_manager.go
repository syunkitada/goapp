package model_manager

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/model"
	"github.com/syunkitada/goapp/pkg/model/model_api"
	"os/exec"
)

var (
	Conf = &config.Conf
)

func MigrateDatabase() error {
	db, err := gorm.Open("mysql", Conf.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}

	glog.Info("Connected DB")

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	db.AutoMigrate(&model.ProjectRole{})
	db.AutoMigrate(&model.Project{})

	if err := model_api.CreateUser(Conf.Admin.Username, Conf.Admin.Password); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateRole("admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.AssignRole("admin", "admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateProject("admin"); err != nil {
		glog.Error(err)
		return err
	}

	if err := model_api.CreateProjectRole("admin", "admin", "admin"); err != nil {
		glog.Error(err)
		return err
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
