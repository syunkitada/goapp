// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type QueryResolver interface {
	GetComputeConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, user *base_spec.UserAuthority, conn *websocket.Conn) (*spec.GetComputeConsoleData, uint8, error)
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
		case "GetComputeConsole":
			var input spec.GetComputeConsole
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}
			data, code, tmpErr := handler.resolver.GetComputeConsole(tctx, &input, req.UserAuthority, conn)
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

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return
}
