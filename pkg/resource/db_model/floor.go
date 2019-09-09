package db_model

import "github.com/jinzhu/gorm"

type Floor struct {
	gorm.Model
	Datacenter string `gorm:"not null;size:50;index;"`
	Name       string `gorm:"not null;size:50;index;"` // Datacenter内でユニーク
	Kind       string `gorm:"not null;size:25;"`
	Zone       string `gorm:"not null;size:50;"`
	Floor      uint8  `gorm:"not null;"`
	Spec       string `gorm:"not null;size:1000;"`
}
