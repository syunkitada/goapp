package monitor

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type ResponseGetNode struct {
	Nodes   []monitor_api_grpc_pb.Node
	TraceId string
	Err     string
}

type ResponseCreateNode struct {
	Node    monitor_api_grpc_pb.Node
	TraceId string
	Err     string
}

type ResponseUpdateNode struct {
	Node    monitor_api_grpc_pb.Node
	TraceId string
	Err     string
}

type ResponseDeleteNode struct {
	TraceId string
	Err     string
}

func (monitor *Monitor) GetNode(c *gin.Context, rc *MonitorContext) (int, string) {
	var reqData monitor_api_grpc_pb.GetNodeRequest
	if err := json.Unmarshal([]byte(rc.action.Data), &reqData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     err,
		})
		return -1, err.Error()
	}
	reqData.TraceId = rc.traceId
	reqData.UserName = rc.userName
	reqData.RoleName = rc.userAuthority.ActionProjectService.RoleName
	reqData.ProjectName = rc.userAuthority.ActionProjectService.ProjectName
	reqData.ProjectRoleName = rc.userAuthority.ActionProjectService.ProjectRoleName

	rep, err := monitor.monitorApiClient.GetNode(&reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     err,
		})
		return int(rep.StatusCode), err.Error()
	}

	if rep.Err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     rep.Err,
		})
		return int(rep.StatusCode), rep.Err
	}

	c.JSON(http.StatusOK, gin.H{
		"TraceId": rc.traceId,
		"Nodes":   rep.Nodes,
	})
	return int(rep.StatusCode), rep.Err
}

func (monitor *Monitor) CtlGetNode(token string, target string) (*ResponseGetNode, error) {
	reqData := monitor_api_grpc_pb.GetNodeRequest{
		Target: target,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: monitor.conf.Ctl.Project,
			ServiceName: "Monitor",
			Name:        "GetNode",
			Data:        string(reqDataJson),
		},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	httpReq, err := http.NewRequest("POST", monitor.conf.Ctl.ApiUrl+"/monitor", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	var resp ResponseGetNode
	var body []byte
	var statusCode int
	if monitor.conf.Default.EnableTest {
		handler := monitor.conf.Authproxy.HttpServer.TestHandler
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, httpReq)
		body = w.Body.Bytes()
		statusCode = w.Code
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Transport: tr,
		}
		httpResp, err := client.Do(httpReq)
		if err != nil {
			return nil, fmt.Errorf("Err: %v", err)
		}
		defer httpResp.Body.Close()
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, fmt.Errorf("Err: %v", err)
		}
		statusCode = httpResp.StatusCode
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	if statusCode != 200 {
		return &resp, fmt.Errorf("Err: %v\nStatusCode: %v\nTraceID: %v", resp.Err, statusCode, resp.TraceId)
	}

	return &resp, nil
}
