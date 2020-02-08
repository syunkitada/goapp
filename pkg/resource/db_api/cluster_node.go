package db_api

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	resource_cluster_api "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) GetClusterNodeServices(tctx *logger.TraceContext, input *spec.GetNodeServices, user *base_spec.UserAuthority) (data []base_spec.NodeService, err error) {
	query := api.DB.Table("cluster_node_services").
		Select("name, kind, role, status, status_reason, state, state_reason").
		Where("deleted_at IS NULL AND cluster = ?", input.Cluster)
	err = query.Find(&data).Error
	return
}

func (api *Api) SyncClusterNodeService(tctx *logger.TraceContext) (err error) {
	var clusters []db_model.Cluster
	if err = api.DB.Find(&clusters).Error; err != nil {
		return
	}

	for _, cluster := range clusters {
		endpoints := strings.Split(cluster.Endpoints, ",")
		client := resource_cluster_api.NewClient(&base_config.ClientConfig{
			Endpoints:             endpoints,
			Token:                 cluster.Token,
			Project:               cluster.Project,
			TlsInsecureSkipVerify: true,
		})

		queries := []base_client.Query{
			base_client.Query{
				Name: "GetNodeServices",
				Data: resource_api_spec.GetNodeServices{},
			},
		}
		res, tmpErr := client.ResourceVirtualAdminGetNodeServices(tctx, queries)
		if tmpErr != nil {
			logger.Warningf(tctx, "Failed GetNodeServices: %s", tmpErr.Error())
			continue
		}
		if tmpErr := api.CreateOrUpdateClusterNodeService(tctx, &cluster, res.NodeServices); tmpErr != nil {
			logger.Warningf(tctx, "Failed CreateOrUpdateClusterNodeService: %s", tmpErr.Error())
			continue
		}
	}
	return
}

func (api *Api) CreateOrUpdateClusterNodeService(tctx *logger.TraceContext, cluster *db_model.Cluster, nodes []base_spec.NodeService) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, node := range nodes {
			var tmp db_model.ClusterNodeService
			if err = tx.Where("name = ? and kind = ?", node.Name, node.Kind).First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.ClusterNodeService{
					NodeService: base_db_model.NodeService{
						Name:         node.Name,
						Kind:         node.Kind,
						Role:         node.Role,
						Status:       node.Status,
						StatusReason: node.StatusReason,
						State:        node.State,
						StateReason:  node.StateReason,
					},
					Cluster: cluster.Name,
				}
				if err = tx.Create(&tmp).Error; err != nil {
					return
				}
			} else {
				if err = tx.Model(&tmp).Updates(map[string]interface{}{
					"Role":         node.Role,
					"Status":       node.Status,
					"StatusReason": node.StatusReason,
					"State":        node.State,
					"StateReason":  node.StateReason,
					"Cluster":      cluster.Name,
				}).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}
