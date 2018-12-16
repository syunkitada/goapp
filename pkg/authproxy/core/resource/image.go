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

type ResponseGetImage struct {
	Images  []resource_api_grpc_pb.Image
	TraceId string
	Err     string
}

type ResponseCreateImage struct {
	Image   resource_api_grpc_pb.Image
	TraceId string
	Err     string
}

type ResponseUpdateImage struct {
	Image   resource_api_grpc_pb.Image
	TraceId string
	Err     string
}

type ResponseDeleteImage struct {
	TraceId string
	Err     string
}

func (resource *Resource) GetImage(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.GetImageRequest
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

	rep, err := resource.resourceApiClient.GetImage(&reqData)
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
		"Images":  rep.Images,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) CreateImage(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.CreateImageRequest
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

	rep, err := resource.resourceApiClient.CreateImage(&reqData)
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
		"Image":   rep.Image,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) UpdateImage(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.UpdateImageRequest
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

	rep, err := resource.resourceApiClient.UpdateImage(&reqData)
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
		"Image":   rep.Image,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) DeleteImage(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.DeleteImageRequest
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

	rep, err := resource.resourceApiClient.DeleteImage(&reqData)
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

func (resource *Resource) CtlGetImage(token string, cluster string, target string) (*ResponseGetImage, error) {
	reqData := resource_api_grpc_pb.GetImageRequest{
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
			Name:        "GetImage",
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

	var resp ResponseGetImage
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

func (resource *Resource) CtlCreateImage(token string, spec string) (*ResponseCreateImage, error) {
	reqData := resource_api_grpc_pb.CreateImageRequest{
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
			Name:        "CreateImage",
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

	var resp ResponseCreateImage
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

func (resource *Resource) CtlUpdateImage(token string, spec string) (*ResponseUpdateImage, error) {
	reqData := resource_api_grpc_pb.UpdateImageRequest{
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
			Name:        "UpdateImage",
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

	var resp ResponseUpdateImage
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

func (resource *Resource) CtlDeleteImage(token string, cluster string, target string) (*ResponseDeleteImage, error) {
	reqData := resource_api_grpc_pb.DeleteImageRequest{
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
			Name:        "DeleteImage",
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

	var resp ResponseDeleteImage
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
