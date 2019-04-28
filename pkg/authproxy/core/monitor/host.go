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

type ResponseGetHost struct {
	HostMap map[string]monitor_api_grpc_pb.Host
	TraceId string
	Err     string
}

type ResponseCreateHost struct {
	Host    monitor_api_grpc_pb.Host
	TraceId string
	Err     string
}

func (monitor *Monitor) GetHost(c *gin.Context, rc *MonitorContext) (int, string) {
	return 0, ""

	// var reqData monitor_api_grpc_pb.GetHostRequest
	// if err := json.Unmarshal([]byte(rc.action.Data), &reqData); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"TraceId": rc.traceId,
	// 		"Err":     err,
	// 	})
	// 	return -1, err.Error()
	// }
	// reqData.TraceId = rc.traceId
	// reqData.UserName = rc.userName
	// reqData.RoleName = rc.userAuthority.ActionProjectService.RoleName
	// reqData.ProjectName = rc.userAuthority.ActionProjectService.ProjectName
	// reqData.ProjectRoleName = rc.userAuthority.ActionProjectService.ProjectRoleName

	// rep, err := monitor.monitorApiClient.GetHost(&reqData)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"TraceId": rc.traceId,
	// 		"Err":     err,
	// 	})
	// 	return int(rep.StatusCode), err.Error()
	// }

	// if rep.Err != "" {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"TraceId": rc.traceId,
	// 		"Err":     rep.Err,
	// 	})
	// 	return int(rep.StatusCode), rep.Err
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"TraceId": rc.traceId,
	// 	"HostMap": rep.HostMap,
	// })
	// return int(rep.StatusCode), rep.Err
}

func (monitor *Monitor) CtlGetHost(token string, index string) (*ResponseGetHost, error) {
	// reqData := monitor_api_grpc_pb.GetHostRequest{
	// 	Index: index,
	// }
	// reqDataJson, err := json.Marshal(reqData)
	// if err != nil {
	// 	return nil, fmt.Errorf("Err: %v", err)
	// }

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: monitor.conf.Ctl.Project,
			ServiceName: "Monitor",
			Queries:     []authproxy_model.Query{},
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

	var resp ResponseGetHost
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
		defer func() { err = httpResp.Body.Close() }()
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
