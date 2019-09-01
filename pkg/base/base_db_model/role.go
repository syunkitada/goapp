package base_db_model

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name      string
	Project   Project `gorm:"foreignkey:ProjectID;association_foreignkey:Refer;"`
	ProjectID uint
	Users     []User `gorm:"many2many:user_roles"`
}
