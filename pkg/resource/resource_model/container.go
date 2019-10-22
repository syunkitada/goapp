package resource_model

import "github.com/jinzhu/gorm"

type Container struct {
	gorm.Model
	Cluster      string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}
