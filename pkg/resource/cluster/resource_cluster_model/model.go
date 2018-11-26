package resource_cluster_model

import (
	"github.com/jinzhu/gorm"
)

type Node struct {
	gorm.Model
	Name               string `gorm:"not null;size:255;"`
	Kind               string `gorm:"not null;size:25;"`
	Role               string `gorm:"not null;size:25;"`
	Status             string `gorm:"not null;size:25;"`
	StatusReason       string `gorm:"not null;size:50;"`
	State              string `gorm:"not null;size:25;"`
	StateReason        string `gorm:"not null;size:50;"`
	ComputeDriver      string `gorm:"not null;size:25;"`
	ContainerDriver    string `gorm:"not null;size:25;"`
	LoadbalancerDriver string `gorm:"not null;size:25;"`
}

type Region struct {
	gorm.Model
	Name string `gorm:"not null;"`
}

type Compute struct {
	gorm.Model
	Name         string `gorm:"not null;size:200;unique_index"`
	FullName     string `gorm:"not null;size:255;unique_index"`
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}
