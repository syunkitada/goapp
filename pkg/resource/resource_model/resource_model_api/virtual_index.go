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
	statusCode = codes.ClientNotFound
	for _, query := range req.Queries {
		switch query.Kind {
		case "get_index":
			response.Index = *modelApi.getVirtualIndex()
		case "get_dashboard-index":
			response.Index = *modelApi.getVirtualIndex()
			statusCode, err = modelApi.GetClusters(tctx, db, query, data)

		case "get_region-service":
			statusCode, err = modelApi.GetRegionService(tctx, db, query, data)
		case "get_region-services":
			statusCode, err = modelApi.GetRegionServices(tctx, db, query, data)
		case "create_region-service":
			statusCode, err = modelApi.CreateRegionService(tctx, db, query)
		case "update_region-service":
			statusCode, err = modelApi.UpdateRegionService(tctx, db, query)
		case "delete_region-service":
			statusCode, err = modelApi.DeleteRegionService(tctx, db, query)

		case "get_cluster":
			statusCode, err = modelApi.GetCluster(tctx, db, query, data)
		case "get_clusters":
			statusCode, err = modelApi.GetClusters(tctx, db, query, data)
		case "create_cluster":
			statusCode, err = modelApi.CreateCluster(tctx, db, query)
		case "update_cluster":
			statusCode, err = modelApi.UpdateCluster(tctx, db, query)
		case "delete_cluster":
			statusCode, err = modelApi.DeleteCluster(tctx, db, query)

		case "get_node":
			statusCode, err = modelApi.GetNode(tctx, db, query, data)
		case "get_nodes":
			statusCode, err = modelApi.GetNodes(tctx, db, query, data)
		// case "create_node":
		// 	statusCode, err = modelApi.CreateNode(tctx, db, query)
		// case "update_node":
		// 	statusCode, err = modelApi.UpdateNode(tctx, db, query)
		case "delete_node":
			statusCode, err = modelApi.DeleteNode(tctx, db, query)

		case "get_compute":
			statusCode, err = modelApi.GetCompute(tctx, db, query, data)
		case "get_computes":
			statusCode, err = modelApi.GetComputes(tctx, db, query, data)
		case "create_compute":
			statusCode, err = modelApi.CreateCompute(tctx, db, query)
		case "update_compute":
			statusCode, err = modelApi.UpdateCompute(tctx, db, query)
		case "delete_compute":
			statusCode, err = modelApi.DeleteCompute(tctx, db, query)

		case "get_image":
			statusCode, err = modelApi.GetImage(tctx, db, query, data)
		case "get_images":
			statusCode, err = modelApi.GetImages(tctx, db, query, data)
		case "create_image":
			statusCode, err = modelApi.CreateImage(tctx, db, query)
		case "update_image":
			statusCode, err = modelApi.UpdateImage(tctx, db, query)
		case "delete_image":
			statusCode, err = modelApi.DeleteImage(tctx, db, query)

		case "get_network-v4":
			statusCode, err = modelApi.GetNetworkV4(tctx, db, query, data)
		case "get_network-v4s":
			statusCode, err = modelApi.GetNetworkV4s(tctx, db, query, data)
		case "create_network-v4":
			statusCode, err = modelApi.CreateNetworkV4(tctx, db, query)
		case "update_network-v4":
			statusCode, err = modelApi.UpdateNetworkV4(tctx, db, query)
		case "delete_network-v4":
			statusCode, err = modelApi.DeleteNetworkV4(tctx, db, query)
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
		resource_model.RegionCmd,
		resource_model.GlobalServiceCmd,
		resource_model.RegionServiceCmd,
		resource_model.ImageCmd,
		resource_model.NetworkV4Cmd,
		resource_model.ComputeCmd,
		resource_model.NodeCmd,
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
					GetQueries:       []string{"get_computes", "get_images"},
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
