package resource_model_api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_utils"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) VirtualAction(tctx *logger.TraceContext,
	req *authproxy_grpc_pb.ActionRequest, rep *authproxy_grpc_pb.ActionReply) {
	var err error
	var statusCode int64
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	response := authproxy_model.ActionResponse{
		Tctx: *req.Tctx,
	}

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		authproxy_utils.MergeResponse(rep, &response, nil, err, codes.RemoteDbError)
		return
	}
	defer func() {
		tmpErr := db.Close()
		if tmpErr != nil {
			logger.Error(tctx, tmpErr)
		}
	}()

	data := map[string]interface{}{}
	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "GetIndex":
			response.Index = *modelApi.getVirtualIndex()
		case "GetDashboardIndex":
			response.Index = *modelApi.getVirtualIndex()
			statusCode, err = modelApi.GetClusters(tctx, db, query, data)

		case "GetCluster":
			statusCode, err = modelApi.GetCluster(tctx, db, query, data)
		case "GetClusters":
			statusCode, err = modelApi.GetClusters(tctx, db, query, data)
		case "CreateCluster":
			statusCode, err = modelApi.CreateCluster(tctx, db, query)
		case "UpdateCluster":
			statusCode, err = modelApi.UpdateCluster(tctx, db, query)
		case "DeleteCluster":
			statusCode, err = modelApi.DeleteCluster(tctx, db, query)

		case "GetCompute":
			statusCode, err = modelApi.GetCompute(tctx, db, query, data)
		case "GetComputes":
			statusCode, err = modelApi.GetComputes(tctx, db, query, data)
		case "CreateCompute":
			statusCode, err = modelApi.CreateCompute(tctx, db, query)
		case "UpdateCompute":
			statusCode, err = modelApi.UpdateCompute(tctx, db, query)
		case "DeleteCompute":
			statusCode, err = modelApi.DeleteCompute(tctx, db, query)
		}

		if err != nil {
			break
		}
	}

	authproxy_utils.MergeResponse(rep, &response, data, err, statusCode)
}

func (modelApi *ResourceModelApi) getVirtualIndex() *index_model.Index {
	cmdMap := map[string]index_model.Cmd{}
	cmdMaps := []map[string]index_model.Cmd{
		resource_model.ClusterCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	return &index_model.Index{
		SyncDelay: 20000,
		CmdMap:    cmdMap,
		View: index_model.Panels{
			Name: "Root",
			Kind: "RoutePanels",
			Panels: []interface{}{
				resource_model.ClustersTable,
				index_model.Tabs{
					Name:             "Resources",
					Kind:             "RouteTabs",
					Subname:          "kind",
					Route:            "/Clusters/:datacenter/Resources/:kind",
					TabParam:         "kind",
					GetQueries:       []string{"GetComputes", "GetImages"},
					ExpectedDataKeys: []string{"Computes", "Images"},
					IsSync:           true,
					Tabs: []interface{}{
						resource_model.ComputesTable,
						resource_model.ImagesTable,
					}, // Tabs
				},
				gin.H{
					"Name":      "Resource",
					"Subname":   "resource",
					"Route":     "/Clusters/:datacenter/Resources/:kind/Detail/:resource/:subkind",
					"Kind":      "RoutePanes",
					"PaneParam": "kind",
					"Panes": []interface{}{
						resource_model.ComputesDetail,
						resource_model.ImagesDetail,
					},
				},
			},
		},
	}
}
