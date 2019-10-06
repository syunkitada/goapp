package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Client struct {
	*base_client.Client
}

func NewClient(conf *base_config.ClientConfig) *Client {
	client := Client{
		Client: base_client.NewClient(conf),
	}
	return &client
}

type UpdateClusterResponse struct {
	base_model.Response
	ResultMap UpdateClusterResultMap
}

type UpdateClusterResultMap struct {
	UpdateCluster UpdateClusterResult
}

type UpdateClusterResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateClusterData
}

func (client *Client) UpdateCluster(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateClusterData, err error) {
	var res UpdateClusterResponse
	client.Login(tctx, &base_spec.Login{})
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateCluster
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
