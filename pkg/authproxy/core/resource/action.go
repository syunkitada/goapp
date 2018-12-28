package resource

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	// "github.com/syunkitada/goapp/pkg/lib/logger"
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

	action := tmpAction.(authproxy_model.ActionRequest)
	rc := &ResourceContext{
		traceId:       tmpTraceId.(string),
		userName:      tmpUsername.(string),
		action:        &action,
		userAuthority: tmpUserAuthority.(*authproxy_model.UserAuthority),
		startTime:     time.Now(),
	}
	// TODO FIX
	fmt.Println(rc.traceId, resource.host, resource.name, map[string]string{
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
		statusCode, errMsg = resource.GetCluster(c, rc)
	case "GetNode":
		statusCode, errMsg = resource.GetNode(c, rc)

	case "GetNetworkV4":
		statusCode, errMsg = resource.GetNetworkV4(c, rc)
	case "CreateNetworkV4":
		statusCode, errMsg = resource.CreateNetworkV4(c, rc)
	case "UpdateNetworkV4":
		statusCode, errMsg = resource.UpdateNetworkV4(c, rc)
	case "DeleteNetworkV4":
		statusCode, errMsg = resource.DeleteNetworkV4(c, rc)

	case "GetCompute":
		statusCode, errMsg = resource.GetCompute(c, rc)
	case "CreateCompute":
		statusCode, errMsg = resource.CreateCompute(c, rc)
	case "UpdateCompute":
		statusCode, errMsg = resource.UpdateCompute(c, rc)
	case "DeleteCompute":
		statusCode, errMsg = resource.DeleteCompute(c, rc)

	case "GetImage":
		statusCode, errMsg = resource.GetImage(c, rc)
	case "CreateImage":
		statusCode, errMsg = resource.CreateImage(c, rc)
	case "UpdateImage":
		statusCode, errMsg = resource.UpdateImage(c, rc)
	case "DeleteImage":
		statusCode, errMsg = resource.DeleteImage(c, rc)

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
		// TODO FIX
		fmt.Println(rc.traceId, resource.host, resource.name, map[string]string{
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
		// TODO FIX
		fmt.Println(rc.traceId, resource.host, resource.name, map[string]string{
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
