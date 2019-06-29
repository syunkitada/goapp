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

	data := map[string]interface{}{}
	response := authproxy_model.ActionResponse{
		Tctx: *req.Tctx,
	}

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		authproxy_utils.MergeResponse(rep, &response, data, err, codes.RemoteDbError)
		return
	}
	defer modelApi.close(tctx, db)

	tmpStatusCode := codes.Unknown
	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "get_index":
			response.Index = *modelApi.getVirtualIndex()
		case "get_dashboard-index":
			response.Index = *modelApi.getVirtualIndex()
			tmpStatusCode, err = modelApi.GetClusters(tctx, db, query, data)

		case "get_region-service":
			tmpStatusCode, err = modelApi.GetRegionService(tctx, db, query, data)
		case "get_region-services":
			tmpStatusCode, err = modelApi.GetRegionServices(tctx, db, query, data)
		case "create_region-service":
			tmpStatusCode, err = modelApi.CreateRegionService(tctx, db, req, query)
		case "update_region-service":
			tmpStatusCode, err = modelApi.UpdateRegionService(tctx, db, query)
		case "delete_region-service":
			tmpStatusCode, err = modelApi.DeleteRegionService(tctx, db, query)

		case "get_cluster":
			tmpStatusCode, err = modelApi.GetCluster(tctx, db, query, data)
		case "get_clusters":
			tmpStatusCode, err = modelApi.GetClusters(tctx, db, query, data)
		case "create_cluster":
			tmpStatusCode, err = modelApi.CreateCluster(tctx, db, query)
		case "update_cluster":
			tmpStatusCode, err = modelApi.UpdateCluster(tctx, db, query)
		case "delete_cluster":
			tmpStatusCode, err = modelApi.DeleteCluster(tctx, db, query)

		case "get_node":
			tmpStatusCode, err = modelApi.GetNode(tctx, db, query, data)
		case "get_nodes":
			tmpStatusCode, err = modelApi.GetNodes(tctx, db, query, data)
		case "update_node":
			tmpStatusCode, err = modelApi.UpdateNode(tctx, db, query)
		case "delete_node":
			tmpStatusCode, err = modelApi.DeleteNode(tctx, db, query)

		case "get_compute":
			tmpStatusCode, err = modelApi.GetCompute(tctx, db, query, data)
		case "get_computes":
			tmpStatusCode, err = modelApi.GetComputes(tctx, db, query, data)

		case "get_image":
			tmpStatusCode, err = modelApi.GetImage(tctx, db, query, data)
		case "get_images":
			tmpStatusCode, err = modelApi.GetImages(tctx, db, query, data)
		case "create_image":
			tmpStatusCode, err = modelApi.CreateImage(tctx, db, query)
		case "update_image":
			tmpStatusCode, err = modelApi.UpdateImage(tctx, db, query)
		case "delete_image":
			tmpStatusCode, err = modelApi.DeleteImage(tctx, db, query)

		case "get_network-v4":
			tmpStatusCode, err = modelApi.GetNetworkV4(tctx, db, query, data)
		case "get_network-v4s":
			tmpStatusCode, err = modelApi.GetNetworkV4s(tctx, db, query, data)
		case "create_network-v4":
			tmpStatusCode, err = modelApi.CreateNetworkV4(tctx, db, query)
		case "update_network-v4":
			tmpStatusCode, err = modelApi.UpdateNetworkV4(tctx, db, query)
		case "delete_network-v4":
			tmpStatusCode, err = modelApi.DeleteNetworkV4(tctx, db, query)
		}

		if err != nil {
			break
		}

		if tmpStatusCode > statusCode {
			statusCode = tmpStatusCode
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
