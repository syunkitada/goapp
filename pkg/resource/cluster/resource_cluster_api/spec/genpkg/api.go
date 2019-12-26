// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	CreateCompute(tctx *logger.TraceContext, input *spec.CreateCompute, user *base_spec.UserAuthority) (*spec.CreateComputeData, uint8, error)
	CreateEventRules(tctx *logger.TraceContext, input *spec.CreateEventRules, user *base_spec.UserAuthority) (*spec.CreateEventRulesData, uint8, error)
	DeleteCompute(tctx *logger.TraceContext, input *spec.DeleteCompute, user *base_spec.UserAuthority) (*spec.DeleteComputeData, uint8, error)
	DeleteComputes(tctx *logger.TraceContext, input *spec.DeleteComputes, user *base_spec.UserAuthority) (*spec.DeleteComputesData, uint8, error)
	DeleteEventRules(tctx *logger.TraceContext, input *spec.DeleteEventRules, user *base_spec.UserAuthority) (*spec.DeleteEventRulesData, uint8, error)
	GetCompute(tctx *logger.TraceContext, input *spec.GetCompute, user *base_spec.UserAuthority) (*spec.GetComputeData, uint8, error)
	GetComputes(tctx *logger.TraceContext, input *spec.GetComputes, user *base_spec.UserAuthority) (*spec.GetComputesData, uint8, error)
	GetEventRule(tctx *logger.TraceContext, input *spec.GetEventRule, user *base_spec.UserAuthority) (*spec.GetEventRuleData, uint8, error)
	GetEventRules(tctx *logger.TraceContext, input *spec.GetEventRules, user *base_spec.UserAuthority) (*spec.GetEventRulesData, uint8, error)
	GetEvents(tctx *logger.TraceContext, input *spec.GetEvents, user *base_spec.UserAuthority) (*spec.GetEventsData, uint8, error)
	GetLogParams(tctx *logger.TraceContext, input *spec.GetLogParams, user *base_spec.UserAuthority) (*spec.GetLogParamsData, uint8, error)
	GetLogs(tctx *logger.TraceContext, input *spec.GetLogs, user *base_spec.UserAuthority) (*spec.GetLogsData, uint8, error)
	GetNode(tctx *logger.TraceContext, input *spec.GetNode, user *base_spec.UserAuthority) (*spec.GetNodeData, uint8, error)
	GetNodeServices(tctx *logger.TraceContext, input *spec.GetNodeServices, user *base_spec.UserAuthority) (*spec.GetNodeServicesData, uint8, error)
	GetNodes(tctx *logger.TraceContext, input *spec.GetNodes, user *base_spec.UserAuthority) (*spec.GetNodesData, uint8, error)
	ReportNode(tctx *logger.TraceContext, input *spec.ReportNode, user *base_spec.UserAuthority) (*spec.ReportNodeData, uint8, error)
	ReportNodeServiceTask(tctx *logger.TraceContext, input *spec.ReportNodeServiceTask, user *base_spec.UserAuthority) (*spec.ReportNodeServiceTaskData, uint8, error)
	SyncNodeService(tctx *logger.TraceContext, input *spec.SyncNodeService, user *base_spec.UserAuthority) (*spec.SyncNodeServiceData, uint8, error)
	UpdateCompute(tctx *logger.TraceContext, input *spec.UpdateCompute, user *base_spec.UserAuthority) (*spec.UpdateComputeData, uint8, error)
	UpdateEventRules(tctx *logger.TraceContext, input *spec.UpdateEventRules, user *base_spec.UserAuthority) (*spec.UpdateEventRulesData, uint8, error)
}

type QueryHandler struct {
	baseConf *base_config.Config
	appConf  *base_config.AppConfig
	resolver QueryResolver
}

func NewQueryHandler(baseConf *base_config.Config, appConf *base_config.AppConfig, resolver QueryResolver) *QueryHandler {
	return &QueryHandler{
		baseConf: baseConf,
		appConf:  appConf,
		resolver: resolver,
	}
}

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, user *base_spec.UserAuthority, httpReq *http.Request, rw http.ResponseWriter,
	req *base_model.Request, rep *base_model.Response) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "Login":
			var input base_spec.Login
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.Login(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    data.Token,
				Secure:   true,
				HttpOnly: true,
				Expires:  time.Now().Add(1 * time.Hour), // TODO Configurable
			} // FIXME SameSite
			http.SetCookie(rw, &cookie)

		case "Logout":
			rep.ResultMap[query.Name] = base_model.Result{
				Code: base_const.CodeOk,
			}
			cookie := http.Cookie{
				Name:     "X-Auth-Token",
				Value:    "",
				Secure:   true,
				HttpOnly: true,
			}
			http.SetCookie(rw, &cookie)

		case "LoginWithToken":
			var input base_spec.LoginWithToken
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.LoginWithToken(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

		case "UpdateService":
			var input base_spec.UpdateService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.UpdateService(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetServiceIndex(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetServiceDashboardIndex(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap["GetServiceDashboardIndex"] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateCompute":
			var input spec.CreateCompute
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateCompute(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "CreateEventRules":
			var input spec.CreateEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.CreateEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteCompute":
			var input spec.DeleteCompute
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteCompute(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteComputes":
			var input spec.DeleteComputes
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteComputes(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "DeleteEventRules":
			var input spec.DeleteEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.DeleteEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetCompute":
			var input spec.GetCompute
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetCompute(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetComputes":
			var input spec.GetComputes
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetComputes(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetEventRule":
			var input spec.GetEventRule
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetEventRule(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetEventRules":
			var input spec.GetEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetEvents":
			var input spec.GetEvents
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetEvents(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetLogParams":
			var input spec.GetLogParams
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetLogParams(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetLogs":
			var input spec.GetLogs
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetLogs(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNode":
			var input spec.GetNode
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNode(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNodeServices":
			var input spec.GetNodeServices
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNodeServices(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "GetNodes":
			var input spec.GetNodes
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetNodes(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "ReportNode":
			var input spec.ReportNode
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.ReportNode(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "ReportNodeServiceTask":
			var input spec.ReportNodeServiceTask
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.ReportNodeServiceTask(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "SyncNodeService":
			var input spec.SyncNodeService
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.SyncNodeService(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateCompute":
			var input spec.UpdateCompute
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateCompute(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}
		case "UpdateEventRules":
			var input spec.UpdateEventRules
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.UpdateEventRules(tctx, &input, user)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_model.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_model.Result{
				Code: code,
				Data: data,
			}

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return nil
}
