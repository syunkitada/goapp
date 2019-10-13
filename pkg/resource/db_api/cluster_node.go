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
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) SyncClusterNode(tctx *logger.TraceContext) (err error) {
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
				Name: "GetNodes",
				Data: resource_api_spec.GetNodes{},
			},
		}
		res, tmpErr := client.ResourceVirtualAdminGetNodes(tctx, queries)
		if tmpErr != nil {
			logger.Warningf(tctx, "Failed GetNodes: %s", tmpErr.Error())
			continue
		}
		if tmpErr := api.CreateOrUpdateClusterNode(tctx, res.Nodes); tmpErr != nil {
			logger.Warningf(tctx, "Failed CreateOrUpdateClusterNode: %s", tmpErr.Error())
			continue
		}
	}
	return
}

func (api *Api) CreateOrUpdateClusterNode(tctx *logger.TraceContext, nodes []base_spec.Node) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, node := range nodes {
			var tmp db_model.ClusterNode
			if err = tx.Where("name = ? and kind = ?", node.Name, node.Kind).First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.ClusterNode{
					Node: base_db_model.Node{
						Name:         node.Name,
						Kind:         node.Kind,
						Role:         node.Role,
						Status:       node.Status,
						StatusReason: node.StatusReason,
						State:        node.State,
						StateReason:  node.StateReason,
					},
				}
				if err = tx.Create(&tmp).Error; err != nil {
					return
				}
			} else {
				tmp.State = node.State
				tmp.StateReason = node.StateReason
				if err = tx.Save(&tmp).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}
