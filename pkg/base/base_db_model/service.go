package base_db_model

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	Name            string
	Scope           string
	Endpoints       string
	ProjectRoles    string
	QueryMap        string
	SyncRootCluster bool
}
