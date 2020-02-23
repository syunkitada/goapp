// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/gorilla/websocket"
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
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

type GetComputeConsoleResponse struct {
	base_protocol.Response
	ResultMap GetComputeConsoleResultMap
}

type GetComputeConsoleResultMap struct {
	GetComputeConsole GetComputeConsoleResult
}

type GetComputeConsoleResult struct {
	Code  uint8
	Error string
	Data  spec.GetComputeConsoleData
}

func (client *Client) ResourceVirtualAdminGetComputeConsole(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetComputeConsoleData, conn *websocket.Conn, err error) {
	var res GetComputeConsoleResponse
	conn, err = client.RequestWs(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetComputeConsole
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
