package db_model

import "github.com/jinzhu/gorm"

type Node struct {
	gorm.Model
	Name     string `gorm:"not null;size:255;"`
	State    string `gorm:"not null;size:25;"`
	Warnings int    `gorm:"not null;"`
	Errors   int    `gorm:"not null;"`
	Labels   string `gorm:"not null;size:500;"`
}
