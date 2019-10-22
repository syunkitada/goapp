package db_model

import "github.com/jinzhu/gorm"

type Datacenter struct {
	gorm.Model
	Name         string    `gorm:"not null;size:50;unique_index;"`
	Kind         string    `gorm:"not null;size:25;"`
	Description  string    `gorm:"not null;size:200;"`
	Region       string    `gorm:"not null;size:50;unique_index;"`
	DomainSuffix string    `gorm:"not null;size:255;unique;"`
	Spec         string    `gorm:"not null;size:1000;"`
	Clusters     []Cluster `gorm:"-"`
}
