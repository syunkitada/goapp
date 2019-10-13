package server

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncNode(tctx); err != nil {
		return
	}
	return
}

func (srv *Server) SyncNode(tctx *logger.TraceContext) (err error) {
	fmt.Println("DEBUG SyncNode")
	// TODO
	// apiClient.ResourceVirtualAdminSyncNode
	return
}
