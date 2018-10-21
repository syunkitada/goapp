package resource_region_model

import (
	"github.com/jinzhu/gorm"
)

type Node struct {
	gorm.Model
	Name string `gorm:"not null;size:255;unique_index;"`
}

type NodeAssignment struct {
	gorm.Model
	Node   Node `gorm:"foreignkey:NetworkV4ID;association_foreignkey:Refer;"`
	NodeID uint `gorm:"not null;"`
}

type NetworkAvailabilityZone struct {
	gorm.Model
	Name string `gorm:"not null;"`
}

type NodeAvailabilityZone struct {
	gorm.Model
	Name string `gorm:"not null;"`
}
