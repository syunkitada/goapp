package resource_model_api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetCluster(tctx *logger.TraceContext, req *resource_api_grpc_pb.GetClusterRequest) *resource_api_grpc_pb.GetClusterReply {
	rep := &resource_api_grpc_pb.GetClusterReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var clusters []resource_model.Cluster
	if err = db.Where("name like ?", req.Target).Find(&clusters).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	rep.Clusters = modelApi.convertClusters(tctx, clusters)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) ValidateClusterName(db *gorm.DB, name string) (bool, error) {
	var err error
	var clusters []resource_model.Cluster
	if err = db.Where("name = ?", name).Find(&clusters).Error; err != nil {
		return false, err
	}

	if len(clusters) == 1 {
		return true, nil
	}
	return false, nil
}

func (modelApi *ResourceModelApi) convertClusters(tctx *logger.TraceContext, clusters []resource_model.Cluster) []*resource_api_grpc_pb.Cluster {
	pbClusters := make([]*resource_api_grpc_pb.Cluster, len(clusters))
	for i, cluster := range clusters {
		updatedAt, err := ptypes.TimestampProto(cluster.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", cluster.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(cluster.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", cluster.Model.CreatedAt)
			continue
		}

		pbClusters[i] = &resource_api_grpc_pb.Cluster{
			Name:      cluster.Name,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbClusters
}

func (modelApi *ResourceModelApi) convertCluster(cluster *resource_model.Cluster) (*resource_api_grpc_pb.Cluster, error) {
	updatedAt, err := ptypes.TimestampProto(cluster.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(cluster.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	clusterPb := &resource_api_grpc_pb.Cluster{
		Name:      cluster.Name,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return clusterPb, nil
}
