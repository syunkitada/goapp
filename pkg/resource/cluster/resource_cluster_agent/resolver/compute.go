package resolver

import (
	"fmt"

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
	var messageType int
	var message []byte
	for {
		fmt.Println("Waiting Messages on WebSocket")
		messageType, message, err = conn.ReadMessage()
		if err != nil {
			logger.Warningf(tctx, "Faild ReadMessage: %s", err.Error())
			return
		}
		fmt.Println("DEBUG message", messageType, string(message))
		if err = conn.WriteMessage(messageType, message); err != nil {
			logger.Warningf(tctx, "Faild WriteMessage: %s", err.Error())
			return
		}
	}
}
