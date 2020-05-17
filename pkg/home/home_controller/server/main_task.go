package server

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	fmt.Println("Exec Contoller Task")

	return
}
