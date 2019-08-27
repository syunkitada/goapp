package db_model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type CustomUser struct {
	Name            string
	RoleID          uint
	RoleName        string
	ProjectName     string
	ProjectRoleID   uint
	ProjectRoleName string
	ServiceID       uint
	ServiceName     string
	ServiceScope    string
}
