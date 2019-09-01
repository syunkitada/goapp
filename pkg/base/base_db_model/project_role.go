package base_db_model

import "github.com/jinzhu/gorm"

type ProjectRole struct {
	gorm.Model
	Name     string    `gorm:"size:63;"`
	Services []Service `gorm:"many2many:project_role_services"`
}
