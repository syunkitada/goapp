package db_model

import "github.com/jinzhu/gorm"

type IgnoreEvent struct {
	gorm.Model
	Name string `gorm:"not null;size:50;unique_index;"`
}
