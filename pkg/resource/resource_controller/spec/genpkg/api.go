// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"fmt"
	"net/http"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
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

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, user *base_spec.UserAuthority, httpReq *http.Request, rw http.ResponseWriter,
	req *base_protocol.Request, rep *base_protocol.Response) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {

		default:
			err = fmt.Errorf("InvalidQueryName: %s", query.Name)
			return err
		}
	}
	return nil
}
