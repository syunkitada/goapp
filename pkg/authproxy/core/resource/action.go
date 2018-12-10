package resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"google.golang.org/grpc/status"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResourceContext struct {
	traceId       string
	userName      string
	userAuthority *authproxy_model.UserAuthority
	startTime     time.Time
	action        *authproxy_model.ActionRequest
}

func (resource *Resource) Action(c *gin.Context) {
	start := time.Now()
	tmpTraceId, traceIdOk := c.Get("TraceId")
	tmpUsername, usernameOk := c.Get("Username")
	tmpUserAuthority, userAuthorityOk := c.Get("UserAuthority")
	tmpAction, actionOk := c.Get("Action")
	if !traceIdOk || !usernameOk || !userAuthorityOk || !actionOk {
		c.JSON(500, gin.H{
			"err": "Invalid request",
		})
		return
	}

	username := tmpUsername.(string)
	action := tmpAction.(authproxy_model.ActionRequest)
	userAuthority := tmpUserAuthority.(*authproxy_model.UserAuthority)
	rc := &ResourceContext{
		traceId:       tmpTraceId.(string),
		userName:      tmpUsername.(string),
		action:        &action,
		userAuthority: tmpUserAuthority.(*authproxy_model.UserAuthority),
		startTime:     time.Now(),
	}
	logger.TraceInfo(rc.traceId, resource.host, resource.name, map[string]string{
		"Msg":             "Start",
		"TraceId":         rc.traceId,
		"Action":          action.Name,
		"User":            rc.userName,
		"Project":         rc.userAuthority.ActionProjectService.ProjectName,
		"RoleName":        rc.userAuthority.ActionProjectService.RoleName,
		"ProjectRoleName": rc.userAuthority.ActionProjectService.ProjectRoleName,
	})

	statusCode := 404
	errMsg := ""
	switch action.Name {
	case "GetCluster":
		rep, err := resource.resourceApiClient.GetCluster(&resource_api_grpc_pb.GetClusterRequest{
			Target: "%",
		})
		if err != nil {
			glog.Error("Failed GetCluster: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Failed GetCluster: %v", err),
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"clusters": rep.Clusters,
			"err":      err,
		})
		return

	case "GetNode":
		var reqData resource_api_grpc_pb.GetNodeRequest
		if err := json.Unmarshal([]byte(action.Data), &reqData); err != nil {
			glog.Errorf("Invalid Request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"err": "Invalid Request",
			})
			c.Abort()
			return
		}
		glog.Info(reqData)

		rep, err := resource.resourceApiClient.GetNode(&reqData)
		if err != nil {
			glog.Error("Failed GetNode: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Failed GetNode: %v", err),
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"nodes": rep.Nodes,
			"err":   err,
		})
		return

	case "GetCompute":
		var reqData resource_api_grpc_pb.GetComputeRequest
		if err := json.Unmarshal([]byte(action.Data), &reqData); err != nil {
			glog.Error("Failed GetNode: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Failed GetNode: %v", err),
			})
			c.Abort()
			return
		}
		glog.Info(reqData)

		rep, err := resource.resourceApiClient.GetCompute(&reqData)
		if err != nil {
			glog.Error("Failed GetNode: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("Failed GetNode: %v", err),
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"computes": rep.Computes,
			"err":      err,
		})
		return

	case "CreateCompute":
		var reqData resource_api_grpc_pb.CreateComputeRequest
		if err := json.Unmarshal([]byte(action.Data), &reqData); err != nil {
			glog.Errorf("Failed unmarshal request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": fmt.Sprintf("@@CreateCompute: Failed  unmarshal request: %v", err),
			})
			c.Abort()
			return
		}
		reqData.UserName = username
		reqData.RoleName = userAuthority.ActionProjectService.RoleName
		reqData.ProjectName = userAuthority.ActionProjectService.ProjectName
		reqData.ProjectRoleName = userAuthority.ActionProjectService.ProjectRoleName

		rep, err := resource.resourceApiClient.CreateCompute(&reqData)
		if err != nil {
			st := status.Convert(err)
			msg := fmt.Sprintf("@@ProxyApiCreateCompute: time=%v%v", time.Now().Sub(start), st.Message())
			glog.Warningf(msg)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": msg,
			})
			c.Abort()
		} else {
			c.JSON(200, gin.H{
				"compute": rep.Compute,
				"err":     err,
			})
		}

	case "GetNetworkV4":
		statusCode, errMsg = resource.GetNetworkV4(c, rc)
	case "CreateNetworkV4":
		statusCode, errMsg = resource.CreateNetworkV4(c, rc)
	case "UpdateNetworkV4":
		statusCode, errMsg = resource.UpdateNetworkV4(c, rc)
	case "DeleteNetworkV4":
		statusCode, errMsg = resource.DeleteNetworkV4(c, rc)

	case "GetState":
		status, err := resource.resourceApiClient.Status()
		if err != nil {
			glog.Error("Failed HealthClient.Status", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "Invalid AuthRequest",
			})
			c.Abort()
		} else {
			c.JSON(200, gin.H{
				"message": status,
			})
		}

	default:
		c.JSON(404, gin.H{
			"message": "NotFound",
		})
	}

	if statusCode == 0 {
		logger.TraceInfo(rc.traceId, resource.host, resource.name, map[string]string{
			"Msg":             "End",
			"TraceId":         rc.traceId,
			"Action":          action.Name,
			"User":            rc.userName,
			"Project":         rc.userAuthority.ActionProjectService.ProjectName,
			"RoleName":        rc.userAuthority.ActionProjectService.RoleName,
			"ProjectRoleName": rc.userAuthority.ActionProjectService.ProjectRoleName,
			"Latency":         strconv.FormatInt(time.Now().Sub(start).Nanoseconds()/1000000, 10),
			"RpcStatusCode":   strconv.Itoa(statusCode),
			"RpcErr":          errMsg,
		})
	} else {
		logger.TraceError(rc.traceId, resource.host, resource.name, map[string]string{
			"Msg":             "End",
			"TraceId":         rc.traceId,
			"Action":          action.Name,
			"User":            rc.userName,
			"Project":         rc.userAuthority.ActionProjectService.ProjectName,
			"RoleName":        rc.userAuthority.ActionProjectService.RoleName,
			"ProjectRoleName": rc.userAuthority.ActionProjectService.ProjectRoleName,
			"Latency":         strconv.FormatInt(time.Now().Sub(start).Nanoseconds()/1000000, 10),
			"RpcStatusCode":   strconv.Itoa(statusCode),
			"RpcErr":          errMsg,
		})
	}
}
