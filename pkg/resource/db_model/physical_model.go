package db_model

import "github.com/jinzhu/gorm"

type PhysicalModel struct {
	gorm.Model
	Kind        string `gorm:"not null;size:25;"`
	Name        string `gorm:"not null;size:200;index;"`
	Description string `gorm:"not null;size:200;"`
	Unit        uint8  `gorm:"not null;"`
	Spec        string `gorm:"not null;size:5000;"`
}
