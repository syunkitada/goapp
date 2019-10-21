package db_model

import "github.com/jinzhu/gorm"

type Rack struct {
	gorm.Model
	Datacenter string `gorm:"not null;size:50;index;"`
	Name       string `gorm:"not null;size:200;index;"` // Datacenter内でユニーク
	Floor      string `gorm:"not null;size:50;index;"`
	Kind       string `gorm:"not null;size:25;"`
	Unit       uint8  `gorm:"not null;"`
	Spec       string `gorm:"not null;size:1000;"`
}
