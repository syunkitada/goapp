package db_model

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:255;"` // Name is unique in Region
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Description  string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}
