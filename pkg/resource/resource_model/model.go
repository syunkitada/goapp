package resource_model

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

type Region struct {
	gorm.Model
	Name string `gorm:"not null;size:50;unique_index;"`
}

type NetworkV4 struct {
	gorm.Model
	Cluster            string           `gorm:"not null;size:50;"`
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Name               string           `gorm:"not null;size:200;"`
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
	NetworkV4   NetworkV4 `gorm:"foreignkey:NetworkV4ID;association_foreignkey:Refer;"`
	NetworkV4ID uint      `gorm:"not null;"`
	IP          string    `gorm:"not null;"`
	Mac         string    `gorm:"not null;"`
}

type GlobalService struct {
	Domain string `gorm:"not null;"` // GSLB Domain
}

type RegionService struct {
	Region string `gorm:"not null;size:50;"`
	Domain string `gorm:"not null;"` // Vip Domain
}

type Compute struct {
	gorm.Model
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Cluster            string           `gorm:"not null;size:50;"`
	Name               string           `gorm:"not null;size:200;"`
	Description        string           `gorm:"not null;size:200;"`
	Kind               string           `gorm:"not null;size:25;"`
	Labels             string           `gorm:"not null;size:255;"`
	Status             string           `gorm:"not null;size:25;"`
	StatusReason       string           `gorm:"not null;size:50;"`
	Spec               string           `gorm:"not null;size:5000;"`
	Domain             string           `gorm:"not null;size:255;"`
	LinkSpec           string           `gorm:"not null;size:2500;"`
}

type Container struct {
	gorm.Model
	Cluster      string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}

type Volume struct {
	gorm.Model
	Cluster      string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}

type Image struct {
	gorm.Model
	Cluster      string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}

type Loadbalancer struct {
	gorm.Model
	Cluster      string `gorm:"not null;size:25;"`
	Name         string `gorm:"not null;size:255;"`
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}
