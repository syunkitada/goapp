// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
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

type GetAllUsersResponse struct {
	base_model.Response
	ResultMap GetAllUsersResultMap
}

type GetAllUsersResultMap struct {
	GetAllUsers GetAllUsersResult
}

type GetAllUsersResult struct {
	Code  uint8
	Error string
	Data  spec.GetAllUsersData
}
type GetUserResponse struct {
	base_model.Response
	ResultMap GetUserResultMap
}

type GetUserResultMap struct {
	GetUser GetUserResult
}

type GetUserResult struct {
	Code  uint8
	Error string
	Data  spec.GetUserData
}
type GetUsersResponse struct {
	base_model.Response
	ResultMap GetUsersResultMap
}

type GetUsersResultMap struct {
	GetUsers GetUsersResult
}

type GetUsersResult struct {
	Code  uint8
	Error string
	Data  spec.GetUsersData
}

func (client *Client) HomeGetAllUsers(tctx *logger.TraceContext, queries []base_client.Query) (data *base_spec.GetAllUsersData, err error) {
	var res GetAllUsersResponse
	err = client.Request(tctx, "Home", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetAllUsers
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) HomeGetUser(tctx *logger.TraceContext, queries []base_client.Query) (data *base_spec.GetUserData, err error) {
	var res GetUserResponse
	err = client.Request(tctx, "Home", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetUser
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) HomeProjectGetUsers(tctx *logger.TraceContext, queries []base_client.Query) (data *base_spec.GetUsersData, err error) {
	var res GetUsersResponse
	err = client.Request(tctx, "HomeProject", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetUsers
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
