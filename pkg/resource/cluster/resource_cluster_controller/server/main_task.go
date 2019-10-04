package server

import (
	"sync"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	// nodeSpec := spec.NodeSpec{}
	// if err = srv.SyncNodeByDb(tctx, &nodeSpec); err != nil {
	// 	return
	// }

	// var role string
	// if role, err = srv.SyncNodeRole(tctx); err != nil {
	// 	return
	// }

	// if role == resource_model.RoleMember {
	// 	return
	// }

	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go srv.SyncRegionService(tctx, &wg)
	// wg.Wait()

	return
}

func (srv *Server) SyncRegionService(tctx *logger.TraceContext, wg *sync.WaitGroup) {
	// defer func() { wg.Done() }()

	// fmt.Println("DEBUG SyncRegionService")

	// errChan := make(chan error)

	// ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, srv.syncResourceTimeout)
	// defer cancel()

	// go func() {
	// 	errChan <- srv.resourceModelApi.SyncRegionService(tctx)
	// }()

	// select {
	// case err = <-errChan:
	// 	break
	// case <-ctx.Done():
	// 	err = ctx.Err()
	// }
	return
}
