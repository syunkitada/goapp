package db_api

import (
	"fmt"
	"strings"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"

	resource_cluster_api "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) SyncClusterNode(tctx *logger.TraceContext) (err error) {
	fmt.Println("DEBUG SyncClusterNode")

	var clusters []db_model.Cluster
	if err = api.DB.Find(&clusters).Error; err != nil {
		return
	}

	// TODO FIX project
	for _, cluster := range clusters {
		endpoints := strings.Split(cluster.Endpoints, ",")
		client := resource_cluster_api.NewClient(&base_config.ClientConfig{
			Endpoints:             endpoints,
			Token:                 cluster.Token,
			Project:               "service",
			TlsInsecureSkipVerify: true,
		})

		queries := []base_client.Query{
			base_client.Query{
				Name: "GetNodes",
				Data: resource_api_spec.GetNodes{},
			},
		}
		res, tmpErr := client.GetNodes(tctx, queries)
		if tmpErr != nil {
			fmt.Println("DEBUG tmpErr", tmpErr)
		}
		fmt.Println("DEBUG GetNodes", endpoints, cluster.Token, res)
	}
	// TODO sync cluster node
	return
}
