package db_model

import "github.com/syunkitada/goapp/pkg/base/base_db_model"

type RegionService struct {
	base_db_model.Model
	Region       string `gorm:"not null;size:25;primary_key;"`
	Name         string `gorm:"not null;size:60;primary_key;"`
	Project      string `gorm:"not null;size:60;primary_key;"`
	Kind         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:100000;"`
}
