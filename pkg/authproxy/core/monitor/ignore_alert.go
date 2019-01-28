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

type ResponseGetIgnoreAlert struct {
	IgnoreAlerts []monitor_api_grpc_pb.IgnoreAlert
	TraceId      string
	Err          string
}

type ResponseCreateIgnoreAlert struct {
	IgnoreAlert monitor_api_grpc_pb.IgnoreAlert
	TraceId     string
	Err         string
}

type ResponseUpdateIgnoreAlert struct {
	IgnoreAlert monitor_api_grpc_pb.IgnoreAlert
	TraceId     string
	Err         string
}

type ResponseDeleteIgnoreAlert struct {
	TraceId string
	Err     string
}

func (monitor *Monitor) GetIgnoreAlert(c *gin.Context, rc *MonitorContext) (int, string) {
	var reqData monitor_api_grpc_pb.GetIgnoreAlertRequest
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

	rep, err := monitor.monitorApiClient.GetIgnoreAlert(&reqData)
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
		"TraceId":      rc.traceId,
		"IgnoreAlerts": rep.IgnoreAlerts,
	})
	return int(rep.StatusCode), rep.Err
}

func (monitor *Monitor) CreateIgnoreAlert(c *gin.Context, rc *MonitorContext) (int, string) {
	var reqData monitor_api_grpc_pb.CreateIgnoreAlertRequest
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

	rep, err := monitor.monitorApiClient.CreateIgnoreAlert(&reqData)
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

	c.JSON(200, gin.H{
		"TraceId":     rc.traceId,
		"IgnoreAlert": rep.IgnoreAlert,
	})
	return int(rep.StatusCode), rep.Err
}

func (monitor *Monitor) UpdateIgnoreAlert(c *gin.Context, rc *MonitorContext) (int, string) {
	var reqData monitor_api_grpc_pb.UpdateIgnoreAlertRequest
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

	rep, err := monitor.monitorApiClient.UpdateIgnoreAlert(&reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     err,
		})
		return int(rep.StatusCode), rep.Err
	}

	if rep.Err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     rep.Err,
		})
		return int(rep.StatusCode), rep.Err
	}

	c.JSON(200, gin.H{
		"TraceId":     rc.traceId,
		"IgnoreAlert": rep.IgnoreAlert,
	})
	return int(rep.StatusCode), rep.Err
}

func (monitor *Monitor) DeleteIgnoreAlert(c *gin.Context, rc *MonitorContext) (int, string) {
	var reqData monitor_api_grpc_pb.DeleteIgnoreAlertRequest
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

	rep, err := monitor.monitorApiClient.DeleteIgnoreAlert(&reqData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     err,
		})
		return int(rep.StatusCode), rep.Err
	}

	if rep.Err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"TraceId": rc.traceId,
			"Err":     rep.Err,
		})
		return int(rep.StatusCode), rep.Err
	}

	c.JSON(200, gin.H{
		"TraceId": rc.traceId,
	})
	return int(rep.StatusCode), rep.Err
}

func (monitor *Monitor) CtlGetIgnoreAlert(token string, index string) (*ResponseGetIgnoreAlert, error) {
	if index == "" {
		index = "%"
	}
	reqData := monitor_api_grpc_pb.GetIgnoreAlertRequest{
		Index: index,
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
			Name:        "GetIgnoreAlert",
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

	var resp ResponseGetIgnoreAlert
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

func (monitor *Monitor) CtlCreateIgnoreAlert(token string, spec string) (*ResponseCreateIgnoreAlert, error) {
	reqData := monitor_api_grpc_pb.CreateIgnoreAlertRequest{
		Spec: spec,
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
			Name:        "CreateIgnoreAlert",
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

	var resp ResponseCreateIgnoreAlert
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

func (monitor *Monitor) CtlUpdateIgnoreAlert(token string, spec string) (*ResponseUpdateIgnoreAlert, error) {
	reqData := monitor_api_grpc_pb.UpdateIgnoreAlertRequest{
		Spec: spec,
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
			Name:        "UpdateIgnoreAlert",
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

	var resp ResponseUpdateIgnoreAlert
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

func (monitor *Monitor) CtlDeleteIgnoreAlert(token string, id uint64) (*ResponseDeleteIgnoreAlert, error) {
	reqData := monitor_api_grpc_pb.DeleteIgnoreAlertRequest{
		Id: id,
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
			Name:        "DeleteIgnoreAlert",
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

	var resp ResponseDeleteIgnoreAlert
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
