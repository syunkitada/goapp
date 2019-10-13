package base_db_model

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Name          string      `gorm:"not null;size:255;"`
	ProjectRole   ProjectRole `gorm:"foreignkey:ProjectRoleID;association_foreignkey:Refer;"`
	ProjectRoleID uint
}
