package server

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncService(tctx, true); err != nil {
		return
	}

	return
}
