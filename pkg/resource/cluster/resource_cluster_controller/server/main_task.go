package server

import (
	"sync"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	nodeSpec := spec.NodeServiceSpec{}
	if err = srv.SyncNodeServiceByDb(tctx, &nodeSpec); err != nil {
		return
	}

	var role string
	if role, err = srv.SyncNodeServiceRole(tctx); err != nil {
		return
	}

	if role == resource_model.RoleMember {
		return
	}

	if err = srv.SyncNodeServiceState(tctx); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	go srv.SyncCompute(tctx, &wg)
	wg.Wait()

	return
}

func (srv *Server) SyncCompute(tctx *logger.TraceContext, wg *sync.WaitGroup) {
	wg.Add(1)
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() {
		logger.EndTrace(tctx, startTime, err, 1)
		wg.Done()
	}()
	err = srv.dbApi.SyncCompute(tctx)
}
