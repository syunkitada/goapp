package db_model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name          string
	ProjectRole   ProjectRole `gorm:"foreignkey:ProjectRoleID;association_foreignkey:Refer;"`
	ProjectRoleID uint
}
