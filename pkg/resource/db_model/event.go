package db_model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type EventRule struct {
	gorm.Model
	Project string     `gorm:"not null;size:50;"`
	Node    string     `gorm:"not null;size:255;"`
	Name    string     `gorm:"not null;size:255;"`
	Msg     string     `gorm:"not null;size:255;"`
	Check   string     `gorm:"not null;size:255;"`
	Level   string     `gorm:"not null;size:50;"`
	Kind    string     `gorm:"not null;size:50;"` // Filter, Silence, Aggregate, Handler
	Until   *time.Time `gorm:""`
	Spec    string     `gorm:"not null;size:100000;"`
}
