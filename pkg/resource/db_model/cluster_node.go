package db_model

import "github.com/syunkitada/goapp/pkg/base/base_db_model"

type ClusterNode struct {
	base_db_model.Node
	Cluster string `gorm:"not null;size:50;"`
	Weight  int
}

type NodeMeta struct {
	Node   base_db_model.Node `gorm:"foreignkey:NodeID;association_foreignkey:Refer;"`
	NodeID uint               `gorm:"not null;primary_key"`
	Weight int
}

type NodeWithMeta struct {
	base_db_model.Node
	Weight int
}
