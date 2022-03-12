// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/home/home_api/spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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

type GetProjectUsersResponse struct {
	base_protocol.Response
	ResultMap GetProjectUsersResultMap
}

type GetProjectUsersResultMap struct {
	GetProjectUsers GetProjectUsersResult
}

type GetProjectUsersResult struct {
	Code  uint8
	Error string
	Data  spec.GetProjectUsersData
}
type UpdateUserPasswordResponse struct {
	base_protocol.Response
	ResultMap UpdateUserPasswordResultMap
}

type UpdateUserPasswordResultMap struct {
	UpdateUserPassword UpdateUserPasswordResult
}

type UpdateUserPasswordResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateUserPasswordData
}

func (client *Client) HomeUpdateUserPassword(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateUserPasswordData, err error) {
	var res UpdateUserPasswordResponse
	err = client.Request(tctx, "Home", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateUserPassword
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) HomeProjectGetProjectUsers(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetProjectUsersData, err error) {
	var res GetProjectUsersResponse
	err = client.Request(tctx, "HomeProject", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetProjectUsers
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
