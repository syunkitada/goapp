package server

import (
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncService(tctx); err != nil {
		return
	}

	nodeSpec := spec.NodeServiceSpec{}
	if err = srv.SyncNodeServiceByDb(tctx, &nodeSpec); err != nil {
		return
	}

	if err = srv.SyncCluster(tctx); err != nil {
		return
	}

	if err = srv.SyncEventRules(tctx); err != nil {
		return
	}

	return
}

func (srv *Server) SyncCluster(tctx *logger.TraceContext) (err error) {
	var token string
	// TODO make username configurable
	if token, err = srv.dbApi.IssueToken("service"); err != nil {
		return
	}
	queries := []base_client.Query{
		base_client.Query{
			Name: "UpdateCluster",
			Data: resource_api_spec.UpdateCluster{
				Name:         srv.clusterName,
				Region:       srv.clusterConf.Region,
				Datacenter:   srv.clusterConf.Datacenter,
				Kind:         srv.clusterConf.Kind,
				DomainSuffix: srv.clusterConf.DomainSuffix,
				Weight:       srv.clusterConf.Weight,
				Token:        token,
				Project:      "service", // TODO make project configurable
				Endpoints:    srv.clusterConf.Api.Endpoints,
			},
		},
	}
	if _, err = srv.rootClient.ResourceVirtualAdminUpdateCluster(tctx, queries); err != nil {
		return
	}

	return
}

func (srv *Server) SyncEventRules(tctx *logger.TraceContext) (err error) {
	var eventRules []db_model.EventRule
	if eventRules, err = srv.dbApi.GetFilterEventRules(tctx); err != nil {
		return
	}
	srv.resolver.SetFilterEventRules(tctx, eventRules)
	return
}
