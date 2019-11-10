package db_model

import "github.com/syunkitada/goapp/pkg/base/base_db_model"

type ClusterNodeService struct {
	base_db_model.NodeService
	Cluster string `gorm:"not null;size:50;"`
	Weight  int
}

type NodeServiceMeta struct {
	NodeService   base_db_model.NodeService `gorm:"foreignkey:NodeServiceID;association_foreignkey:Refer;"`
	NodeServiceID uint                      `gorm:"not null;primary_key"`
	Weight        int
}

type NodeServiceWithMeta struct {
	base_db_model.NodeService
	Weight int
}
