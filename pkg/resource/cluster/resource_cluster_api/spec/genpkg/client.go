// This code is auto generated.
// Don't modify this code.

package genpkg

import (
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

type CreateComputeResponse struct {
	base_protocol.Response
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
type CreateEventRulesResponse struct {
	base_protocol.Response
	ResultMap CreateEventRulesResultMap
}

type CreateEventRulesResultMap struct {
	CreateEventRules CreateEventRulesResult
}

type CreateEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.CreateEventRulesData
}
type DeleteComputeResponse struct {
	base_protocol.Response
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
	base_protocol.Response
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
type DeleteEventRulesResponse struct {
	base_protocol.Response
	ResultMap DeleteEventRulesResultMap
}

type DeleteEventRulesResultMap struct {
	DeleteEventRules DeleteEventRulesResult
}

type DeleteEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.DeleteEventRulesData
}
type GetComputeResponse struct {
	base_protocol.Response
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
	base_protocol.Response
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
type GetEventRuleResponse struct {
	base_protocol.Response
	ResultMap GetEventRuleResultMap
}

type GetEventRuleResultMap struct {
	GetEventRule GetEventRuleResult
}

type GetEventRuleResult struct {
	Code  uint8
	Error string
	Data  spec.GetEventRuleData
}
type GetEventRulesResponse struct {
	base_protocol.Response
	ResultMap GetEventRulesResultMap
}

type GetEventRulesResultMap struct {
	GetEventRules GetEventRulesResult
}

type GetEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.GetEventRulesData
}
type GetEventsResponse struct {
	base_protocol.Response
	ResultMap GetEventsResultMap
}

type GetEventsResultMap struct {
	GetEvents GetEventsResult
}

type GetEventsResult struct {
	Code  uint8
	Error string
	Data  spec.GetEventsData
}
type GetLogParamsResponse struct {
	base_protocol.Response
	ResultMap GetLogParamsResultMap
}

type GetLogParamsResultMap struct {
	GetLogParams GetLogParamsResult
}

type GetLogParamsResult struct {
	Code  uint8
	Error string
	Data  spec.GetLogParamsData
}
type GetLogsResponse struct {
	base_protocol.Response
	ResultMap GetLogsResultMap
}

type GetLogsResultMap struct {
	GetLogs GetLogsResult
}

type GetLogsResult struct {
	Code  uint8
	Error string
	Data  spec.GetLogsData
}
type GetNodeResponse struct {
	base_protocol.Response
	ResultMap GetNodeResultMap
}

type GetNodeResultMap struct {
	GetNode GetNodeResult
}

type GetNodeResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodeData
}
type GetNodeMetricsResponse struct {
	base_protocol.Response
	ResultMap GetNodeMetricsResultMap
}

type GetNodeMetricsResultMap struct {
	GetNodeMetrics GetNodeMetricsResult
}

type GetNodeMetricsResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodeMetricsData
}
type GetNodeServicesResponse struct {
	base_protocol.Response
	ResultMap GetNodeServicesResultMap
}

type GetNodeServicesResultMap struct {
	GetNodeServices GetNodeServicesResult
}

type GetNodeServicesResult struct {
	Code  uint8
	Error string
	Data  spec.GetNodeServicesData
}
type GetNodesResponse struct {
	base_protocol.Response
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
type ReportNodeResponse struct {
	base_protocol.Response
	ResultMap ReportNodeResultMap
}

type ReportNodeResultMap struct {
	ReportNode ReportNodeResult
}

type ReportNodeResult struct {
	Code  uint8
	Error string
	Data  spec.ReportNodeData
}
type ReportNodeServiceTaskResponse struct {
	base_protocol.Response
	ResultMap ReportNodeServiceTaskResultMap
}

type ReportNodeServiceTaskResultMap struct {
	ReportNodeServiceTask ReportNodeServiceTaskResult
}

type ReportNodeServiceTaskResult struct {
	Code  uint8
	Error string
	Data  spec.ReportNodeServiceTaskData
}
type SyncNodeServiceResponse struct {
	base_protocol.Response
	ResultMap SyncNodeServiceResultMap
}

type SyncNodeServiceResultMap struct {
	SyncNodeService SyncNodeServiceResult
}

type SyncNodeServiceResult struct {
	Code  uint8
	Error string
	Data  spec.SyncNodeServiceData
}
type UpdateComputeResponse struct {
	base_protocol.Response
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
type UpdateEventRulesResponse struct {
	base_protocol.Response
	ResultMap UpdateEventRulesResultMap
}

type UpdateEventRulesResultMap struct {
	UpdateEventRules UpdateEventRulesResult
}

type UpdateEventRulesResult struct {
	Code  uint8
	Error string
	Data  spec.UpdateEventRulesData
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
func (client *Client) ResourceVirtualAdminGetNodeServices(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodeServicesData, err error) {
	var res GetNodeServicesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNodeServices
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminSyncNodeService(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.SyncNodeServiceData, err error) {
	var res SyncNodeServiceResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.SyncNodeService
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminReportNodeServiceTask(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.ReportNodeServiceTaskData, err error) {
	var res ReportNodeServiceTaskResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.ReportNodeServiceTask
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
func (client *Client) ResourceVirtualAdminGetNode(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodeData, err error) {
	var res GetNodeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNode
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetNodeMetrics(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetNodeMetricsData, err error) {
	var res GetNodeMetricsResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetNodeMetrics
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminReportNode(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.ReportNodeData, err error) {
	var res ReportNodeResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.ReportNode
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetLogs(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetLogsData, err error) {
	var res GetLogsResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetLogs
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetLogParams(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetLogParamsData, err error) {
	var res GetLogParamsResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetLogParams
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetEvents(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetEventsData, err error) {
	var res GetEventsResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetEvents
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetEventRule(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetEventRuleData, err error) {
	var res GetEventRuleResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetEventRule
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminGetEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.GetEventRulesData, err error) {
	var res GetEventRulesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.GetEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminCreateEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.CreateEventRulesData, err error) {
	var res CreateEventRulesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.CreateEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminUpdateEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.UpdateEventRulesData, err error) {
	var res UpdateEventRulesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
func (client *Client) ResourceVirtualAdminDeleteEventRules(tctx *logger.TraceContext, queries []base_client.Query) (data *spec.DeleteEventRulesData, err error) {
	var res DeleteEventRulesResponse
	err = client.Request(tctx, "ResourceVirtualAdmin", queries, &res, true)
	if err != nil {
		return
	}
	result := res.ResultMap.DeleteEventRules
	if result.Code >= 100 || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
