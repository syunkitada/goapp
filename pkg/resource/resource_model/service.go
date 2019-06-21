package resource_model

import "github.com/jinzhu/gorm"

type GlobalService struct {
	gorm.Model
	Domain string `gorm:"not null;"` // GSLB Domain
}

type Region struct {
	gorm.Model
	Name string `gorm:"not null;size:50;unique_index;"`
}

type RegionService struct {
	gorm.Model
	Region string `gorm:"not null;size:50;"`
	Domain string `gorm:"not null;"` // Vip Domain
}
