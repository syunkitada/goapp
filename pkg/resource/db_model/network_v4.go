package db_model

import "github.com/jinzhu/gorm"

type NetworkV4 struct {
	gorm.Model
	Cluster            string           `gorm:"not null;size:50;"`
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Name               string           `gorm:"not null;size:200;"`
	Description        string           `gorm:"not null;size:255;"`
	Kind               string           `gorm:"not null;size:25;"`
	Labels             string           `gorm:"not null;size:255;"`
	Status             string           `gorm:"not null;size:25;"`
	StatusReason       string           `gorm:"not null;size:50;"`
	Spec               string           `gorm:"not null;size:1000;"`
	Subnet             string           `gorm:"not null;"`
	StartIp            string           `gorm:"not null;"`
	EndIp              string           `gorm:"not null;"`
	Gateway            string           `gorm:"not null;"`
}

type NetworkV4Port struct {
	gorm.Model
	NetworkV4    NetworkV4 `gorm:"foreignkey:NetworkV4ID;association_foreignkey:Refer;"`
	NetworkV4ID  uint      `gorm:"not null;"`
	Ip           string    `gorm:"not null;"`
	Mac          string    `gorm:"not null;"`
	ResourceKind string    `gorm:"not null;"`
	ResourceName string    `gorm:"not null;"`
}
