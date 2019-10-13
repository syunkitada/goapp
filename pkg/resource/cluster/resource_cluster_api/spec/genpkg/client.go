// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
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

type CreateComputeResponse struct {
	base_model.Response
	ResultMap CreateComputeResultMap
}

type CreateComputeResultMap struct {
	CreateCompute CreateComputeResult
}

type CreateComputeResult struct {
	Code  uint8
	Error string
	Data  spec.CreateComputeData
}
type DeleteComputeResponse struct {
	base_model.Response
	ResultMap DeleteComputeResultMap
}

type DeleteComputeResultMap struct {
	DeleteCompute DeleteComputeResult
}

type DeleteComputeResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteComputeData
}
type DeleteComputesResponse struct {
	base_model.Response
	ResultMap DeleteComputesResultMap
}

type DeleteComputesResultMap struct {
	DeleteComputes DeleteComputesResult
}

type DeleteComputesResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteComputesData
}
type GetComputeResponse struct {
	base_model.Response
	ResultMap GetComputeResultMap
}

type GetComputeResultMap struct {
	GetCompute GetComputeResult
}

type GetComputeResult struct {
	Code  uint8
	Error string
	Data  spec.GetComputeData
}
type GetComputesResponse struct {
	base_model.Response
	ResultMap GetComputesResultMap
}

type GetComputesResultMap struct {
	GetComputes GetComputesResult
}

type GetComputesResult struct {
	Code  uint8
	Error string
	Data  spec.GetComputesData
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
type SyncNodeResponse struct {
	base_model.Response
	ResultMap SyncNodeResultMap
}

type SyncNodeResultMap struct {
	SyncNode SyncNodeResult
}

type SyncNodeResult struct {
	Code  uint8
	Error string
	Data  spec.SyncNodeData
}
type UpdateComputeResponse struct {
	base_model.Response
	ResultMap UpdateComputeResultMap
}

type UpdateComputeResultMap struct {
	UpdateCompute UpdateComputeResult
}

type UpdateComputeResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateComputeData
}

func (client *Client) ResourceVirtualAdminGetCompute(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetComputeData, err error) {
	var res GetComputeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetCompute
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetComputes(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetComputesData, err error) {
	var res GetComputesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetComputes
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateCompute(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateComputeData, err error) {
	var res CreateComputeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateCompute
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateCompute(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateComputeData, err error) {
	var res UpdateComputeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateCompute
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteCompute(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteComputeData, err error) {
	var res DeleteComputeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteCompute
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteComputes(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteComputesData, err error) {
	var res DeleteComputesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteComputes
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetNodes(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodesData, err error) {
	var res GetNodesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNodes
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminSyncNode(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.SyncNodeData, err error) {
	var res SyncNodeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.SyncNode
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
