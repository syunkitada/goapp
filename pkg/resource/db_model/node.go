package db_model

import "github.com/jinzhu/gorm"

type Node struct {
	gorm.Model
	Name    string `gorm:"not null;size:255;"`
	Kind    string `gorm:"not null;size:25;"`
	State   string `gorm:"not null;size:25;"`
	Wanings string `gorm:"not null;size:25;"`
	Errors  string `gorm:"not null;size:25;"`
	Labels  string `gorm:"not null;size:500;"`
}
