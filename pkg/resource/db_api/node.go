package db_api

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) GetNodes(tctx *logger.TraceContext, input *spec.GetNodes, user *base_spec.UserAuthority) (data []spec.Node, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "GetNodes",
			Data: *input,
		},
	}

	getNodesData, tmpErr := client.ResourceVirtualAdminGetNodes(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetNodes: %s", tmpErr.Error())
		return
	}
	data = getNodesData.Nodes
	return
}

func (api *Api) GetNode(tctx *logger.TraceContext, input *spec.GetNode, user *base_spec.UserAuthority) (data *spec.Node, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "GetNode",
			Data: *input,
		},
	}

	getNodeData, tmpErr := client.ResourceVirtualAdminGetNode(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetNode: %s", tmpErr.Error())
		return
	}
	data = &getNodeData.Node
	return
}

func (api *Api) GetLogParams(tctx *logger.TraceContext, input *spec.GetLogParams, user *base_spec.UserAuthority) (data *spec.GetLogParamsData, err error) {
	client, ok := api.clusterClientMap[input.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "GetLogParams",
			Data: *input,
		},
	}

	getLogParamsData, tmpErr := client.ResourceVirtualAdminGetLogParams(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed GetLogParams: %s", tmpErr.Error())
		return
	}
	data = getLogParamsData
	return
}
