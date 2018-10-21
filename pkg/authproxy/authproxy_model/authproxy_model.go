package authproxy_model

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
	Name     string
	Services []Service `gorm:"many2many:project_role_services"`
}

type Service struct {
	gorm.Model
	Name  string
	Scope string
}

type Action struct {
	Name          string
	Service       Service `gorm:"foreignkey:ServiceID;association_foreignkey:Refer;"`
	ServiceID     uint
	Role          Role `gorm:"foreignkey:RoleID;association_foreignkey:Refer;"`
	RoleID        uint
	ProjectRole   ProjectRole `gorm:"foreignkey:ProjectRoleID;association_foreignkey:Refer;"`
	ProjectRoleID uint
}

type CustomUser struct {
	gorm.Model
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

type CustomProject struct {
	Name            string
	RoleName        string
	ProjectRoleName string
}

type UserAuthority struct {
	ServiceMap           map[string]uint
	ProjectServiceMap    map[string]ProjectService
	ActionProjectService ProjectService
}

type ProjectService struct {
	RoleID          uint
	RoleName        string
	ProjectName     string
	ProjectRoleID   uint
	ProjectRoleName string
	ServiceMap      map[string]uint
}
