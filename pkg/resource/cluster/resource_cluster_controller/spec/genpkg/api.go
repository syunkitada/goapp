// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type QueryResolver interface {
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

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return
}
