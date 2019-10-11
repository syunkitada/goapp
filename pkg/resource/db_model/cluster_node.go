package db_model

import "github.com/syunkitada/goapp/pkg/base/base_db_model"

type ClusterNode struct {
	base_db_model.Node
	Cluster string `gorm:"not null;size:50;"`
}
