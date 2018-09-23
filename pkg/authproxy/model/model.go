package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	gorm.Model
	Name      string
	Project   Project `gorm:"foreignkey:ProjectID;association_foreignkey:Refer;"`
	ProjectID uint
	Users     []User `gorm:"many2many:user_roles"`
}

type Project struct {
	gorm.Model
	Name          string
	ProjectRole   ProjectRole `gorm:"foreignkey:ProjectRoleID;association_foreignkey:Refer;"`
	ProjectRoleID uint
}

type ProjectRole struct {
	gorm.Model
	Name string
}

type CustomUser struct {
	gorm.Model
	Name            string
	Password        string
	RoleName        string
	ProjectName     string
	ProjectRoleName string
}

type CustomProject struct {
}
