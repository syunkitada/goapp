package db_model

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	Name  string
	Scope string
}
