package server

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
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
					Spec:         nodeSpec,
				},
			},
		},
	}

	var syncNodeServiceData *api_spec.SyncNodeServiceData
	if syncNodeServiceData, err = srv.apiClient.ResourceVirtualAdminSyncNodeService(tctx, queries); err != nil {
		return
	}

	var computeAssignmentReports []spec.AssignmentReport
	if computeAssignmentReports, err = srv.SyncComputeAssignments(tctx, syncNodeServiceData.Task.ComputeAssignments); err != nil {
		return err
	}
	fmt.Println("DEBUG reports", computeAssignmentReports)

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
