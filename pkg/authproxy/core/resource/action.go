package resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"google.golang.org/grpc/status"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (resource *Resource) Action(c *gin.Context) {
	start := time.Now()
	tmpUsername, usernameOk := c.Get("Username")
	tmpUserAuthority, userAuthorityOk := c.Get("UserAuthority")
	tmpAction, actionOk := c.Get("Action")
	if !usernameOk || !userAuthorityOk || !actionOk {
		c.JSON(500, gin.H{
			"error": "Invalid request",
		})
		return
	}

	username := tmpUsername.(string)
	action := tmpAction.(authproxy_model.ActionRequest)
	userAuthority := tmpUserAuthority.(*authproxy_model.UserAuthority)

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
				"error": "Invalid Request",
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

	case "GetState":
		status, err := resource.resourceApiClient.Status()
		if err != nil {
			glog.Error("Failed HealthClient.Status", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid AuthRequest",
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"message": status,
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
			return
		}
		c.JSON(200, gin.H{
			"compute": rep.Compute,
			"err":     err,
		})
		return

	}

	c.JSON(200, gin.H{
		"message": "Health",
	})
}
