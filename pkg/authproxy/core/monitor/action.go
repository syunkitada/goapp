package monitor

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
)

type MonitorContext struct {
	traceId       string
	userName      string
	userAuthority *authproxy_model.UserAuthority
	startTime     time.Time
	action        *authproxy_model.ActionRequest
}

func (monitor *Monitor) Action(c *gin.Context) {
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
	rc := &MonitorContext{
		traceId:       tmpTraceId.(string),
		userName:      tmpUsername.(string),
		action:        &action,
		userAuthority: tmpUserAuthority.(*authproxy_model.UserAuthority),
		startTime:     time.Now(),
	}
	fmt.Println(rc.traceId, monitor.host, monitor.name, map[string]string{
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
	case "GetNode":
		statusCode, errMsg = monitor.GetNode(c, rc)
	case "GetHost":
		statusCode, errMsg = monitor.GetHost(c, rc)
	default:
		c.JSON(404, gin.H{
			"Err":     "NotFoundAction",
			"TraceId": rc.traceId,
		})
	}

	if statusCode == 0 {
		fmt.Println(rc.traceId, monitor.host, monitor.name, map[string]string{
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
		fmt.Println(rc.traceId, monitor.host, monitor.name, map[string]string{
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
