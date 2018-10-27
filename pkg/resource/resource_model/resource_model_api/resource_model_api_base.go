package resource_model_api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) UpdateNode(req *resource_api_grpc_pb.UpdateNodeRequest) error {
	var err error
	db, dbErr := gorm.Open("mysql", modelApi.Conf.Resource.Database.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}
	db.LogMode(modelApi.Conf.Default.EnableDatabaseLog)

	var node resource_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		node = resource_model.Node{
			Name:         req.Name,
			Kind:         req.Kind,
			Role:         req.Role,
			Status:       req.Status,
			StatusReason: req.StatusReason,
		}
		if err = db.Create(&node).Error; err != nil {
			return err
		}
	} else {
		node.Status = req.Status
		node.StatusReason = req.StatusReason
		if err = db.Save(&node).Error; err != nil {
			return err
		}
	}

	glog.Info("Updated Node")
	return nil
}
