package server

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncService(tctx); err != nil {
		return
	}

	nodeSpec := spec.NodeServiceSpec{}
	if err = srv.SyncNodeServiceByDb(tctx, &nodeSpec); err != nil {
		return
	}

	if err = srv.dbApi.SyncClusterClient(tctx); err != nil {
		return
	}
	return
}
