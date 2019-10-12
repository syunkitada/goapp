package genpkg

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
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

type GetNodesResponse struct {
	base_model.Response
	ResultMap GetNodesResultMap
}

type GetNodesResultMap struct {
	GetNodes GetNodesResult
}

type GetNodesResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodesData
}

func (client *Client) GetNodes(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodesData, err error) {
	var res GetNodesResponse
	// client.Login(tctx, &base_spec.Login{})
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		fmt.Println("DEBUG GetNodes Request", err)
		return
	}

	if res.Code >= 100 || res.Error != "" {
		err = error_utils.NewInvalidResponseError(res.Code, res.Error)
		return
	}

	result := res.ResultMap.GetNodes
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
