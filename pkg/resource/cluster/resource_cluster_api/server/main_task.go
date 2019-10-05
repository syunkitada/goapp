package server

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncCluster(tctx); err != nil {
		return
	}
	// if err = srv.SyncService(tctx); err != nil {
	// 	return
	// }

	// nodeSpec := spec.NodeSpec{}
	// if err = srv.SyncNodeByDb(tctx, &nodeSpec); err != nil {
	// 	return
	// }
	return
}

func (srv *Server) SyncCluster(tctx *logger.TraceContext) (err error) {
	var token string
	if token, err = srv.dbApi.IssueToken(srv.clusterConf.Api.Name); err != nil {
		return
	}
	queries := []base_client.Query{
		base_client.Query{
			Name: "UpdateCluster",
			Data: resource_api_spec.UpdateCluster{
				Name:      srv.clusterConf.Api.Name,
				Region:    srv.clusterConf.RegionName,
				Token:     token,
				Endpoints: srv.clusterConf.Api.Endpoints,
			},
		},
	}
	if _, err = srv.rootClient.UpdateCluster(tctx, queries); err != nil {
		return
	}

	return
}
