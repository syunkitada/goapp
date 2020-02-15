package db_model

import "github.com/jinzhu/gorm"

type Maintenance struct {
	gorm.Model
	Name string `gorm:"not null;size:255;"`
}
