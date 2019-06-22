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

func (modelApi *ResourceModelApi) PhysicalAction(tctx *logger.TraceContext,
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
	defer func() {
		tmpErr := db.Close()
		if tmpErr != nil {
			logger.Error(tctx, tmpErr)
		}
	}()

	tmpStatusCode := codes.Unknown
	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "get_index":
			response.Index = *modelApi.getPhysicalIndex()
		case "get_dashboard-index":
			response.Index = *modelApi.getPhysicalIndex()
			tmpStatusCode, err = modelApi.GetDatacenters(tctx, db, query, data)

		case "get_region":
			tmpStatusCode, err = modelApi.GetRegion(tctx, db, query, data)
		case "get_regions":
			tmpStatusCode, err = modelApi.GetRegions(tctx, db, query, data)
		case "create_region":
			tmpStatusCode, err = modelApi.CreateRegion(tctx, db, query)
		case "update_region":
			tmpStatusCode, err = modelApi.UpdateRegion(tctx, db, query)
		case "delete_region":
			tmpStatusCode, err = modelApi.DeleteRegion(tctx, db, query)

		case "get_datacenter":
			tmpStatusCode, err = modelApi.GetDatacenter(tctx, db, query, data)
		case "get_datacenters":
			tmpStatusCode, err = modelApi.GetDatacenters(tctx, db, query, data)
		case "create_datacenter":
			tmpStatusCode, err = modelApi.CreateDatacenter(tctx, db, query)
		case "update_datacenter":
			tmpStatusCode, err = modelApi.UpdateDatacenter(tctx, db, query)
		case "delete_datacenter":
			tmpStatusCode, err = modelApi.DeleteDatacenter(tctx, db, query)

		case "get_floor":
			tmpStatusCode, err = modelApi.GetFloor(tctx, db, query, data)
		case "get_floors":
			tmpStatusCode, err = modelApi.GetFloors(tctx, db, query, data)
		case "create_floor":
			tmpStatusCode, err = modelApi.CreateFloor(tctx, db, query)
		case "update_floor":
			tmpStatusCode, err = modelApi.UpdateFloor(tctx, db, query)
		case "delete_floor":
			tmpStatusCode, err = modelApi.DeleteFloor(tctx, db, query)

		case "get_rack":
			tmpStatusCode, err = modelApi.GetRack(tctx, db, query, data)
		case "get_racks":
			tmpStatusCode, err = modelApi.GetRacks(tctx, db, query, data)
		case "create_rack":
			tmpStatusCode, err = modelApi.CreateRack(tctx, db, query)
		case "update_rack":
			tmpStatusCode, err = modelApi.UpdateRack(tctx, db, query)
		case "delete_rack":
			tmpStatusCode, err = modelApi.DeleteRack(tctx, db, query)

		case "get_physical-resource":
			tmpStatusCode, err = modelApi.GetPhysicalResource(tctx, db, query, data)
		case "get_physical-resources":
			tmpStatusCode, err = modelApi.GetPhysicalResources(tctx, db, query, data)
		case "create_physical-resource":
			tmpStatusCode, err = modelApi.CreatePhysicalResource(tctx, db, query)
		case "update_physical-resource":
			tmpStatusCode, err = modelApi.UpdatePhysicalResource(tctx, db, query)
		case "delete_physical-resource":
			tmpStatusCode, err = modelApi.DeletePhysicalResource(tctx, db, query)

		case "get_physical-model":
			tmpStatusCode, err = modelApi.GetPhysicalModel(tctx, db, query, data)
		case "get_physical-models":
			tmpStatusCode, err = modelApi.GetPhysicalModels(tctx, db, query, data)
		case "create_physical-model":
			tmpStatusCode, err = modelApi.CreatePhysicalModel(tctx, db, query)
		case "update_physical-model":
			tmpStatusCode, err = modelApi.UpdatePhysicalModel(tctx, db, query)
		case "delete_physical-model":
			tmpStatusCode, err = modelApi.DeletePhysicalModel(tctx, db, query)
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

func (modelApi *ResourceModelApi) getPhysicalIndex() *index_model.Index {
	cmdMap := map[string]index_model.Cmd{}
	cmdMaps := []map[string]index_model.Cmd{
		resource_model.RegionCmd,
		resource_model.DatacenterCmd,
		resource_model.RackCmd,
		resource_model.FloorCmd,
		resource_model.PhysicalModelCmd,
		resource_model.PhysicalResourceCmd,
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
				resource_model.DatacentersTable,
				index_model.Tabs{
					Name:             "Resources",
					Kind:             "RouteTabs",
					Subname:          "kind",
					Route:            "/Datacenters/:datacenter/Resources/:kind",
					TabParam:         "kind",
					GetQueries:       []string{"get_physical-resources", "get_racks", "get_floors", "get_physical-models"},
					ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
					IsSync:           true,
					Tabs: []interface{}{
						resource_model.PhysicalResourcesTable,
						resource_model.RacksTable,
						resource_model.FloorsTable,
						resource_model.PhysicalModelsTable,
					}, // Tabs
				},
				gin.H{
					"Name":      "Resource",
					"Subname":   "resource",
					"Route":     "/Datacenters/:datacenter/Resources/:kind/Detail/:resource/:subkind",
					"Kind":      "RoutePanes",
					"PaneParam": "kind",
					"Panes": []interface{}{
						resource_model.PhysicalModelsDetail,
						resource_model.PhysicalResourcesDetail,
					},
				},
			},
		},
	}
}
