package db_model

import "github.com/jinzhu/gorm"

type RegionService struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;primary_key;"`
	Name         string `gorm:"not null;size:63;primary_key;"`
	Project      string `gorm:"not null;size:63;primary_key;"`
	Kind         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:100000;"`
}
