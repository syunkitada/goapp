package resource_model

import (
	"github.com/jinzhu/gorm"
)

type ComputeResource struct {
	gorm.Model
	Name         string
	Labels       string
	Kind         string
	Status       string
	StatusReason string
}

type VolumeResource struct {
	gorm.Model
	Name         string
	Labels       string
	Kind         string
	Status       string
	StatusReason string
}

type ImageResource struct {
	gorm.Model
	Name         string
	Labels       string
	Kind         string
	Status       string
	StatusReason string
}

type LoadbalancerResource struct {
	gorm.Model
	Name         string
	Labels       string
	Kind         string
	Status       string
	StatusReason string
}

type RegionAvailabilityZone struct {
	gorm.Model
	Name string
}

type NetworkV4 struct {
	gorm.Model
	Name string
}

type NetworkV4Port struct {
	gorm.Model
	NetworkV4   NetworkV4 `gorm:"foreignkey:NetworkV4ID;association_foreignkey:Refer;"`
	NetworkV4ID uint
	IP          string
}
