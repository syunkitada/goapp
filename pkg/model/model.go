package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Password     string
	Roles        []Role        `gorm:"many2many:user_roles;"`
	ProjectRoles []ProjectRole `gorm:"many2many:user_project_roles;"`
}

type Role struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_roles"`
}

type Project struct {
	gorm.Model
	Name string
}

type ProjectRole struct {
	gorm.Model
	Name      string
	Project   Project `gorm:"foreignkey:ProjectID;association_foreignkey:Refer"`
	ProjectID uint
	Users     []User `gorm:"many2many:user_project_roles"`
}
