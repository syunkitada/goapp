package db_model

import "github.com/jinzhu/gorm"

type Alert struct {
	gorm.Model
	Name string `gorm:"not null;size:50;"`
}
