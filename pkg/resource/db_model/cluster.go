package db_model

import "github.com/jinzhu/gorm"

type Cluster struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;"`
	Datacenter   string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:50;unique_index;"`
	Kind         string `gorm:"not null;size:25;"`
	Project      string `gorm:"not null;size:255;"`
	Description  string `gorm:"not null;size:200;"`
	DomainSuffix string `gorm:"not null;size:255;unique;"`
	Labels       string `gorm:"not null;size:500;"`
	Spec         string `gorm:"not null;size:1000;"`
	Weight       int    `gorm:"not null;"`
	Endpoints    string `gorm:"not null;"`
	Token        string `gorm:"not null;"`
}

type ClusterStatistic struct {
	Warnings  int `gorm:"not null;"`
	Criticals int `gorm:"not null;"`
	Nodes     int `gorm:"not null;"`
	Instances int `gorm:"not null;"`
}
