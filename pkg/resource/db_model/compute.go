package db_model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Compute struct {
	gorm.Model
	PhysicalResource   PhysicalResource `gorm:"foreignkey:PhysicalResourceID;association_foreignkey:Refer;"`
	PhysicalResourceID uint             `gorm:"not null;"`
	Region             string           `gorm:"not null;size:50;"`
	Cluster            string           `gorm:"not null;size:50;"`
	RegionService      string           `gorm:"not null;size:63;"`
	Name               string           `gorm:"not null;size:255;unique_index"`
	Kind               string           `gorm:"not null;size:25;"`
	Labels             string           `gorm:"not null;size:255;"`
	Status             string           `gorm:"not null;size:25;"`
	StatusReason       string           `gorm:"not null;size:50;"`
	Spec               string           `gorm:"not null;size:5000;"`
	Project            string           `gorm:"not null;size:63;"`
	LinkSpec           string           `gorm:"not null;size:2500;"`
	Image              string           `gorm:"not null;size:255;"`
	Vcpus              uint             `gorm:"not null;"`
	Memory             uint             `gorm:"not null;"`
	Disk               uint             `gorm:"not null;"`
}

type ComputeAssignment struct {
	gorm.Model
	Compute       Compute                   `gorm:"foreignkey:ComputeID;association_foreignkey:Refer;"`
	ComputeID     uint                      `gorm:"not null;"`
	NodeService   base_db_model.NodeService `gorm:"foreignkey:NodeServiceID;association_foreignkey:Refer;"`
	NodeServiceID uint                      `gorm:"not null;"`
	Status        string                    `gorm:"not null;size:25;"`
	StatusReason  string                    `gorm:"not null;size:50;"`
}

type ComputeAssignmentWithComputeAndNodeService struct {
	ID               uint
	UpdatedAt        time.Time
	ComputeID        uint
	ComputeSpec      string
	ComputeName      string
	NodeServiceID    uint
	NodeServiceName  string
	Status           string
	StatusReason     string
	Cost             int // calcurated by vcpu, memory, disk
	ServiceToken     string
	ServiceEndpoints string
}

type ComputeAssignmentEx struct {
	ID        uint
	Status    string
	Spec      spec.RegionServiceComputeSpec
	UpdatedAt time.Time
}
