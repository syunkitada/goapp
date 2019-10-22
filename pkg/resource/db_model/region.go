package db_model

import "github.com/jinzhu/gorm"

type Region struct {
	gorm.Model
	Name string `gorm:"not null;size:25;unique_index;"`
	Kind string `gorm:"not null;size:25;"`
}
