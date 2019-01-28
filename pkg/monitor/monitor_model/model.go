package monitor_model

import (
	"github.com/jinzhu/gorm"
)

type Node struct {
	gorm.Model
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Role         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	State        string `gorm:"not null;size:25;"`
	StateReason  string `gorm:"not null;size:50;"`
}

type IgnoreAlert struct {
	gorm.Model
	Index  string `gorm:"not null;size:100;"`
	Host   string `gorm:"not null;size:255;"`
	Name   string `gorm:"not null;size:255;"`
	Level  string `gorm:"not null;size:25;"`
	User   string `gorm:"size:50;"`
	Reason string `gorm:"size:255;"`
	Until  int64  `gorm:"not null;"`
}
