package db_model

import "github.com/jinzhu/gorm"

type RegionService struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:63;"` // Vip Domain
	Project      string `gorm:"not null;size:63;"`
	Kind         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:100000;"`
}
