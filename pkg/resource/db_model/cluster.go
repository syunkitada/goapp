package db_model

import "github.com/jinzhu/gorm"

type Cluster struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;"`
	Datacenter   string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:50;unique_index;"`
	Kind         string `gorm:"not null;size:25;"`
	Description  string `gorm:"not null;size:200;"`
	DomainSuffix string `gorm:"not null;size:255;unique;"`
	Labels       string `gorm:"not null;size:500;"`
	Spec         string `gorm:"not null;size:1000;"`
	Weight       int    `gorm:"not null;"`
}
