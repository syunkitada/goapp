package db_model

import "github.com/jinzhu/gorm"

type GlobalService struct {
	gorm.Model
	Domain string `gorm:"not null;"` // GSLB Domain
}
