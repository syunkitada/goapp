// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceIndexData, uint8, error)
	GetProjectServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	GetProjectServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (*base_spec.GetServiceDashboardIndexData, uint8, error)
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

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, httpReq *http.Request, rw http.ResponseWriter,
	req *base_protocol.Request, rep *base_protocol.Response) (err error) {
	for _, query := range req.Queries {
		switch query.Name {
		case "Login":
			var input base_spec.Login
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.Login(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
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
			rep.ResultMap[query.Name] = base_protocol.Result{
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
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.LoginWithToken(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

		case "UpdateService":
			var input base_spec.UpdateService
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.UpdateService(tctx, &input)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}

		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetServiceIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetProjectServiceIndex":
			var input base_spec.GetServiceIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetProjectServiceIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetServiceDashboardIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap["GetServiceDashboardIndex"] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
				Code: code,
				Data: data,
			}
		case "GetProjectServiceDashboardIndex":
			var input base_spec.GetServiceDashboardIndex
			if tmpErr := json.Unmarshal([]byte(query.Data), &input); tmpErr != nil {
				rep.ResultMap[query.Name] = base_protocol.Result{
					Code:  base_const.CodeClientBadRequest,
					Error: tmpErr.Error(),
				}
				break
			}

			data, code, tmpErr := handler.resolver.GetProjectServiceDashboardIndex(tctx, &input, req.UserAuthority)
			if tmpErr != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.ResultMap["GetServiceDashboardIndex"] = base_protocol.Result{
					Code:  code,
					Error: tmpErr.Error(),
				}
				break
			}
			rep.ResultMap[query.Name] = base_protocol.Result{
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

func (handler *QueryHandler) ExecWs(tctx *logger.TraceContext, httpReq *http.Request, rw http.ResponseWriter,
	req *base_protocol.Request, rep *base_protocol.Response, conn *websocket.Conn) (err error) {
	for _, query := range req.Queries {
		switch query.Name {

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return
}
