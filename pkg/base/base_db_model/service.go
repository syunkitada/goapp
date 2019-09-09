package base_db_model

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	Name            string `gorm:"not null;size:50;unique_index;"`
	Scope           string `gorm:"not null;size:50;"`
	Endpoints       string `gorm:"not null;size:1000;"`
	ProjectRoles    string `gorm:"not null;size:1000;"`
	QueryMap        string `gorm:"not null;size:10000;"`
	SyncRootCluster bool   `gorm:"not null;"`
}
