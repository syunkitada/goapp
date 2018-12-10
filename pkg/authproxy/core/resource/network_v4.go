package resource

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
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseGetNetworkV4 struct {
	Networks []resource_api_grpc_pb.NetworkV4
	TraceId  string
	Err      string
}

type ResponseCreateNetworkV4 struct {
	Network resource_api_grpc_pb.NetworkV4
	TraceId string
	Err     string
}

type ResponseUpdateNetworkV4 struct {
	Network resource_api_grpc_pb.NetworkV4
	TraceId string
	Err     string
}

type ResponseDeleteNetworkV4 struct {
	Network resource_api_grpc_pb.NetworkV4
	TraceId string
	Err     string
}

func (resource *Resource) GetNetworkV4(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.GetNetworkV4Request
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

	rep, err := resource.resourceApiClient.GetNetworkV4(&reqData)
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
		"TraceId":  rc.traceId,
		"Networks": rep.Networks,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) CreateNetworkV4(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.CreateNetworkV4Request
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

	rep, err := resource.resourceApiClient.CreateNetworkV4(&reqData)
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
		"TraceId": rc.traceId,
		"Network": rep.Network,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) UpdateNetworkV4(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.UpdateNetworkV4Request
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

	rep, err := resource.resourceApiClient.UpdateNetworkV4(&reqData)
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
		"Network": rep.Network,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) DeleteNetworkV4(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.DeleteNetworkV4Request
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

	rep, err := resource.resourceApiClient.DeleteNetworkV4(&reqData)
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
		"Network": rep.Network,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) CtlGetNetworkV4(token string, cluster string, target string) (*ResponseGetNetworkV4, error) {
	reqData := resource_api_grpc_pb.GetNetworkV4Request{
		Cluster: cluster,
		Target:  target,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: resource.conf.Ctl.Project,
			ServiceName: "Resource",
			Name:        "GetNetworkV4",
			Data:        string(reqDataJson),
		},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	var resp ResponseGetNetworkV4
	var body []byte
	var statusCode int
	if resource.conf.Default.EnableTest {
		handler := resource.conf.Authproxy.HttpServer.TestHandler
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

func (resource *Resource) CtlCreateNetworkV4(token string, spec string) (*ResponseCreateNetworkV4, error) {
	reqData := resource_api_grpc_pb.CreateNetworkV4Request{
		Spec: spec,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: resource.conf.Ctl.Project,
			ServiceName: "Resource",
			Name:        "CreateNetworkV4",
			Data:        string(reqDataJson),
		},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	var resp ResponseCreateNetworkV4
	var body []byte
	var statusCode int
	if resource.conf.Default.EnableTest {
		handler := resource.conf.Authproxy.HttpServer.TestHandler
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

func (resource *Resource) CtlUpdateNetworkV4(token string, spec string) (*ResponseUpdateNetworkV4, error) {
	reqData := resource_api_grpc_pb.UpdateNetworkV4Request{
		Spec: spec,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: resource.conf.Ctl.Project,
			ServiceName: "Resource",
			Name:        "UpdateNetworkV4",
			Data:        string(reqDataJson),
		},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	var resp ResponseUpdateNetworkV4
	var body []byte
	var statusCode int
	if resource.conf.Default.EnableTest {
		handler := resource.conf.Authproxy.HttpServer.TestHandler
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

func (resource *Resource) CtlDeleteNetworkV4(token string, cluster string, target string) (*ResponseDeleteNetworkV4, error) {
	reqData := resource_api_grpc_pb.DeleteNetworkV4Request{
		Cluster: cluster,
		Target:  target,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: resource.conf.Ctl.Project,
			ServiceName: "Resource",
			Name:        "DeleteNetworkV4",
			Data:        string(reqDataJson),
		},
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}

	var resp ResponseDeleteNetworkV4
	var body []byte
	var statusCode int
	if resource.conf.Default.EnableTest {
		handler := resource.conf.Authproxy.HttpServer.TestHandler
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
