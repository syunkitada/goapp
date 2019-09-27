// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type QueryResolver interface {
	Login(tctx *logger.TraceContext, input *base_spec.Login) (*base_spec.LoginData, uint8, error)
	LoginWithToken(tctx *logger.TraceContext, input *base_spec.LoginWithToken, user *base_spec.UserAuthority) (*base_spec.LoginWithTokenData, uint8, error)
	UpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (*base_spec.UpdateServiceData, uint8, error)
	GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex) (*base_spec.GetServiceIndexData, uint8, error)
	GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex) (*base_spec.GetServiceDashboardIndexData, uint8, error)
	GetAllUsers(tctx *logger.TraceContext, input *base_spec.GetAllUsers) (*base_spec.GetAllUsersData, uint8, error)
	GetUser(tctx *logger.TraceContext, input *base_spec.GetUser) (*base_spec.GetUserData, uint8, error)
	GetUsers(tctx *logger.TraceContext, input *base_spec.GetUsers) (*base_spec.GetUsersData, uint8, error)
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

			data, code, tmpErr := handler.resolver.GetServiceIndex(tctx, &input)
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

			data, code, tmpErr := handler.resolver.GetServiceDashboardIndex(tctx, &input)
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
		case "GetAllUsers":
			var input base_spec.GetAllUsers
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetAllUsers(tctx, &input)
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
		case "GetUser":
			var input base_spec.GetUser
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetUser(tctx, &input)
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
		case "GetUsers":
			var input base_spec.GetUsers
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			data, code, tmpErr := handler.resolver.GetUsers(tctx, &input)
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
		}
	}
	return nil
}
