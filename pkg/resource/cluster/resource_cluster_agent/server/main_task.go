package server

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (srv *Server) MainTask(tctx *logger.TraceContext) (err error) {
	if err = srv.SyncNodeService(tctx); err != nil {
		return
	}
	return
}

func (srv *Server) SyncNodeService(tctx *logger.TraceContext) (err error) {
	nodeSpec := resource_api_spec.NodeServiceSpec{}
	var token string
	if token, err = srv.dbApi.IssueToken("service"); err != nil {
		return
	}
	queries := []base_client.Query{
		base_client.Query{
			Name: "SyncNodeService",
			Data: resource_api_spec.SyncNodeService{
				NodeService: base_spec.NodeService{
					Name:         srv.baseConf.Host,
					Kind:         srv.clusterConf.Agent.Name,
					Role:         base_const.RoleMember,
					Status:       base_const.StatusEnabled,
					StatusReason: "Default",
					State:        base_const.StateUp,
					StateReason:  "SyncNodeService",
					Token:        token,
					Endpoints:    srv.clusterConf.Agent.AppConfig.Endpoints,
					Spec:         nodeSpec,
				},
			},
		},
	}

	var syncNodeServiceData *api_spec.SyncNodeServiceData
	if syncNodeServiceData, err = srv.apiClient.ResourceVirtualAdminSyncNodeService(tctx, queries); err != nil {
		return
	}

	srv.computeAssignmentsMutex.Lock()
	srv.computeAssignments = syncNodeServiceData.Task.ComputeAssignments
	srv.computeAssignmentsMutex.Unlock()

	srv.computeAssignmentReportsMutex.Lock()
	computeAssignmentReports := srv.computeAssignmentReports
	srv.computeAssignmentReportsMutex.Unlock()

	reportNodeServiceTaskQueries := []base_client.Query{
		base_client.Query{
			Name: "ReportNodeServiceTask",
			Data: resource_api_spec.ReportNodeServiceTask{
				ComputeAssignmentReports: computeAssignmentReports,
			},
		},
	}
	if _, err = srv.apiClient.ResourceVirtualAdminReportNodeServiceTask(tctx, reportNodeServiceTaskQueries); err != nil {
		return
	}
	fmt.Println("DEBUG reported")
	return
}
