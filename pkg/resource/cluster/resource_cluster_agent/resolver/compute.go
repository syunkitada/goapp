package resolver

import (
	"github.com/gorilla/websocket"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (resolver *Resolver) GetComputeConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, user *base_spec.UserAuthority, conn *websocket.Conn) (data *spec.GetComputeConsoleData, code uint8, err error) {
	code = base_const.CodeOk
	data = &spec.GetComputeConsoleData{}
	if conn == nil {
		return
	}

	err = resolver.computeDriver.ProxyConsole(tctx, input, conn)
	return
}
