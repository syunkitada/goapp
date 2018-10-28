package resource_model_api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

func (modelApi *ResourceClusterModelApi) GetNode(req *resource_cluster_api_grpc_pb.GetNodeRequest) (*resource_cluster_api_grpc_pb.GetNodeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_cluster_model.Node
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	pbNodes := make([]*resource_cluster_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		pbNodes[i] = &resource_cluster_api_grpc_pb.Node{
			Name:         node.Name,
			Kind:         node.Kind,
			Role:         node.Role,
			Enable:       node.Enable,
			EnableReason: node.EnableReason,
			Status:       node.Status,
			StatusReason: node.StatusReason,
		}
	}

	reply := &resource_cluster_api_grpc_pb.GetNodeReply{
		Nodes: pbNodes,
	}

	return reply, nil
}

func (modelApi *ResourceClusterModelApi) UpdateNode(req *resource_cluster_api_grpc_pb.UpdateNodeRequest) error {
	var err error
	db, dbErr := gorm.Open("mysql", modelApi.cluser.Database.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var node resource_cluster_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		node = resource_cluster_model.Node{
			Name:         req.Name,
			Kind:         req.Kind,
			Role:         req.Role,
			Enable:       resource_cluster_model.StatusDisabled,
			EnableReason: "Registerd default status",
			Status:       req.Status,
			StatusReason: req.StatusReason,
		}
		if err = db.Create(&node).Error; err != nil {
			return err
		}
	} else {
		if req.Enable != "" && req.EnableReason != "" {
			node.Enable = req.Enable
			node.EnableReason = req.EnableReason
		}
		node.Status = req.Status
		node.StatusReason = req.StatusReason
		if err = db.Save(&node).Error; err != nil {
			return err
		}
	}

	glog.Info("Updated Node")
	return nil
}
