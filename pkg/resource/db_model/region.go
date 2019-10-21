package db_model

import "github.com/syunkitada/goapp/pkg/base/base_db_model"

type Region struct {
	base_db_model.Model
	Name string `gorm:"not null;size:25;primary_key;"`
	Kind string `gorm:"not null;size:25;"`
}
