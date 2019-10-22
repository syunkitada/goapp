package db_model

import "github.com/jinzhu/gorm"

type PhysicalResource struct {
	gorm.Model
	Datacenter    string `gorm:"not null;size:50;index;"`
	Rack          string `gorm:"not null;size:50;index;"`
	Cluster       string `gorm:"not null;size:50;index;"`  // 仮想リソースを扱う場合はClusterに紐図かせる
	Name          string `gorm:"not null;size:200;index;"` // Datacenter内でユニーク
	Kind          string `gorm:"not null;size:25;"`        // Server, Pdu, L2Switch, L3Switch, RootSwitch
	PhysicalModel string `gorm:"not null;size:200;"`
	RackPosition  uint8  `gorm:"not null;"`
	Status        string `gorm:"not null;size:25;"`
	StatusReason  string `gorm:"not null;size:50;"`
	PowerLinks    string `gorm:"not null;size:5000;"`
	NetLinks      string `gorm:"not null;size:5000;"`
	Spec          string `gorm:"not null;size:5000;"`
}
