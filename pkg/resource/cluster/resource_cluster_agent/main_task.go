package resource_cluster_agent

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceClusterAgentServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNode(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterAgentServer) UpdateNode(tctx *logger.TraceContext) error {
	// systemMetricsReader := metricsReaderMap["system"]

	nodes := []resource_model.NodeSpec{
		resource_model.NodeSpec{
			Name:           srv.conf.Default.Host,
			Kind:           resource_model.KindResourceClusterAgent,
			Role:           resource_model.RoleMember,
			Labels:         srv.labels,
			ResourceLabels: srv.resourceLabels,
			Status:         resource_model.StatusEnabled,
			StatusReason:   "Default",
			State:          resource_model.StateUp,
			StateReason:    "UpdateNode",
		},
	}
	specs, err := json_utils.Marshal(nodes)
	if err != nil {
		return err
	}
	queries := []authproxy_model.Query{
		authproxy_model.Query{
			Kind: "update_node",
			StrParams: map[string]string{
				"Specs": string(specs),
			},
		},
	}
	rep, err := srv.resourceClusterApiClient.Action(
		logger.NewActionTraceContext(tctx, "system", "system", queries))
	if err != nil {
		return err
	}

	var response resource_model.UpdateNodeResponse
	if err = json_utils.Unmarshal(rep.Response, &response); err != nil {
		return err
	}

	if response.Tctx.StatusCode != codes.OkUpdated {
		err = fmt.Errorf("UnexpectedStatusCode: %v", response.Tctx.StatusCode)
		return err
	}

	var computeAssignmentReports []resource_model.AssignmentReport
	if computeAssignmentReports, err = srv.SyncComputeAssignments(tctx, response.Data.ComputeAssignments); err != nil {
		return err
	}

	assignmentReportMap := resource_model.AssignmentReportMap{
		ComputeAssignmentReports: computeAssignmentReports,
	}
	assignmentReportMapBytes, err := json_utils.Marshal(assignmentReportMap)
	queries = []authproxy_model.Query{
		authproxy_model.Query{
			Kind: "update_node_assignments",
			StrParams: map[string]string{
				"AssignmentReportMap": string(assignmentReportMapBytes),
			},
		},
	}
	rep, err = srv.resourceClusterApiClient.Action(
		logger.NewActionTraceContext(tctx, "system", "system", queries))
	if err != nil {
		return err
	}

	return nil
}
