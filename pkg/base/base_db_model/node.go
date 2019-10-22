package base_db_model

import "github.com/jinzhu/gorm"

type Node struct {
	gorm.Model
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Role         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	State        string `gorm:"not null;size:25;"`
	StateReason  string `gorm:"not null;size:50;"`
	Labels       string `gorm:"not null;size:500;"`
	Spec         string `gorm:"not null;size:5000;"`
}
