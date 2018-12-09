package resource

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseGetNetworkV4 struct {
	Networks   []resource_api_grpc_pb.NetworkV4
	StackTrace []string
	Err        string
}

type ResponseCreateNetworkV4 struct {
	Network    resource_api_grpc_pb.NetworkV4
	StackTrace []string
	Err        string
}

type ResponseUpdateNetworkV4 struct {
	Network    resource_api_grpc_pb.NetworkV4
	StackTrace []string
	Err        string
}

type ResponseDeleteNetworkV4 struct {
	Network    resource_api_grpc_pb.NetworkV4
	StackTrace []string
	Err        string
}

func (resource *Resource) GetNetworkV4(c *gin.Context, username string, userAuthority *authproxy_model.UserAuthority, action *authproxy_model.ActionRequest) error {
	var reqData resource_api_grpc_pb.GetNetworkV4Request
	if err := json.Unmarshal([]byte(action.Data), &reqData); err != nil {
		return err
	}

	rep, err := resource.resourceApiClient.GetNetworkV4(&reqData)
	if err != nil {
		st := status.Convert(err)
		return fmt.Errorf(st.Message())
	}

	c.JSON(200, gin.H{
		"networks": rep.Networks,
		"err":      err,
	})
	return nil
}

func (resource *Resource) CreateNetworkV4(c *gin.Context, rc *ResourceContext) {
	traceFormat := "ProxyApi.Resource.CreateNetworkV4: time=%v"

	var reqData resource_api_grpc_pb.CreateNetworkV4Request
	if err := json.Unmarshal([]byte(rc.action.Data), &reqData); err != nil {
		t := fmt.Sprintf(traceFormat, time.Now().Sub(rc.startTime))
		logger.Error(resource.name, map[string]string{"msg": err.Error()})
		c.JSON(http.StatusInternalServerError, gin.H{
			"stackTrace": []string{t},
			"err":        err,
		})
		return
	}
	reqData.UserName = rc.userName
	reqData.RoleName = rc.userAuthority.ActionProjectService.RoleName
	reqData.ProjectName = rc.userAuthority.ActionProjectService.ProjectName
	reqData.ProjectRoleName = rc.userAuthority.ActionProjectService.ProjectRoleName

	rep, err := resource.resourceApiClient.CreateNetworkV4(&reqData)
	t := fmt.Sprintf(traceFormat, time.Now().Sub(rc.startTime))
	if err != nil {
		logger.Error(resource.name, map[string]string{"msg": err.Error()})
		st := status.Convert(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"stackTrace": []string{t},
			"err":        st.Message(),
		})
		return
	}

	if rep.Err != "" {
		logger.TraceError(resource.name, rc.traceId, map[string]string{"err": rep.Err})
		c.JSON(http.StatusInternalServerError, gin.H{
			"stackTrace": append(rep.StackTrace, t),
			"err":        rep.Err,
		})
		return
	}

	c.JSON(200, gin.H{
		"stackTrace": append(rep.StackTrace, t),
		"network":    rep.Network,
	})
	return
}

func (resource *Resource) UpdateNetworkV4(c *gin.Context, username string, userAuthority *authproxy_model.UserAuthority, action *authproxy_model.ActionRequest) error {
	var reqData resource_api_grpc_pb.UpdateNetworkV4Request
	if err := json.Unmarshal([]byte(action.Data), &reqData); err != nil {
		return err
	}
	reqData.UserName = username
	reqData.RoleName = userAuthority.ActionProjectService.RoleName
	reqData.ProjectName = userAuthority.ActionProjectService.ProjectName
	reqData.ProjectRoleName = userAuthority.ActionProjectService.ProjectRoleName

	rep, err := resource.resourceApiClient.UpdateNetworkV4(&reqData)
	if err != nil {
		st := status.Convert(err)
		return fmt.Errorf(st.Message())
	}

	c.JSON(200, gin.H{
		"network": rep.Network,
		"err":     err,
	})
	return nil
}

func (resource *Resource) DeleteNetworkV4(c *gin.Context, username string, userAuthority *authproxy_model.UserAuthority, action *authproxy_model.ActionRequest) error {
	var reqData resource_api_grpc_pb.DeleteNetworkV4Request
	if err := json.Unmarshal([]byte(action.Data), &reqData); err != nil {
		return err
	}
	reqData.UserName = username
	reqData.RoleName = userAuthority.ActionProjectService.RoleName
	reqData.ProjectName = userAuthority.ActionProjectService.ProjectName
	reqData.ProjectRoleName = userAuthority.ActionProjectService.ProjectRoleName

	rep, err := resource.resourceApiClient.DeleteNetworkV4(&reqData)
	if err != nil {
		st := status.Convert(err)
		return fmt.Errorf(st.Message())
	}

	c.JSON(200, gin.H{
		"network": rep.Network,
		"err":     err,
	})
	return nil
}

func (resource *Resource) CtlGetNetworkV4(token string, cluster string, target string) (*ResponseGetNetworkV4, error) {
	reqData := resource_api_grpc_pb.GetNetworkV4Request{
		Cluster: cluster,
		Target:  target,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, err
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
			return nil, err
		}
		defer httpResp.Body.Close()
		var readAllErr error
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, readAllErr
		}
		statusCode = httpResp.StatusCode
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	if statusCode != 200 {
		return &resp, fmt.Errorf("Invalid StatusCode: %v", statusCode)
	}

	return &resp, nil
}

func (resource *Resource) CtlCreateNetworkV4(token string, spec string) (*ResponseCreateNetworkV4, error) {
	start := time.Now()
	reqData := resource_api_grpc_pb.CreateNetworkV4Request{
		Spec: spec,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("@@CtlCreateNetworkV4: time=%v: %v", time.Now().Sub(start), err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("@@CtlCreateNetworkV4: time=%v: %v", time.Now().Sub(start), err)
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
			return nil, err
		}
		defer httpResp.Body.Close()
		var readAllErr error
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, readAllErr
		}
		statusCode = httpResp.StatusCode
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("@@CtlCreateNetworkV4: time=%v: %v", time.Now().Sub(start), err)
	}

	traceLog := fmt.Sprintf("CtlCreateNetworkV4: time=%v", time.Now().Sub(start))
	resp.StackTrace = append(resp.StackTrace, traceLog)

	if statusCode != 200 {
		return &resp, fmt.Errorf("statusCode=%v, %v", statusCode, resp.Err)
	}

	return &resp, nil
}

func (resource *Resource) CtlUpdateNetworkV4(token string, spec string) (*ResponseUpdateNetworkV4, error) {
	start := time.Now()
	reqData := resource_api_grpc_pb.UpdateNetworkV4Request{
		Spec: spec,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("@@CtlUpdateNetworkV4: time=%v: %v", time.Now().Sub(start), err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("@@CtlUpdateNetworkV4: time=%v: %v", time.Now().Sub(start), err)
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
			return nil, err
		}
		defer httpResp.Body.Close()
		var readAllErr error
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, readAllErr
		}
		statusCode = httpResp.StatusCode
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("@@CtlUpdateNetworkV4: time=%v: %v", time.Now().Sub(start), err)
	}

	if statusCode != 200 {
		return &resp, fmt.Errorf("@@CtlUpdateNetworkV4: time=%v, statusCode=%v %v", time.Now().Sub(start), statusCode, resp.Err)
	}

	return &resp, nil
}

func (resource *Resource) CtlDeleteNetworkV4(token string, cluster string, target string) (*ResponseDeleteNetworkV4, error) {
	start := time.Now()
	reqData := resource_api_grpc_pb.DeleteNetworkV4Request{
		Cluster: cluster,
		Target:  target,
	}
	reqDataJson, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("@@CtlDeleteNetworkV4: time=%v: %v", time.Now().Sub(start), err)
	}

	httpReq, err := http.NewRequest("POST", resource.conf.Ctl.ApiUrl+"/resource", bytes.NewBuffer(reqJson))
	if err != nil {
		return nil, fmt.Errorf("@@CtlDeleteNetworkV4: time=%v: %v", time.Now().Sub(start), err)
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
			return nil, err
		}
		defer httpResp.Body.Close()
		var readAllErr error
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return nil, readAllErr
		}
		statusCode = httpResp.StatusCode
	}

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("@@CtlDeleteNetworkV4: time=%v: %v", time.Now().Sub(start), err)
	}

	if statusCode != 200 {
		return &resp, fmt.Errorf("@@CtlDeleteNetworkV4: time=%v, statusCode=%v %v", time.Now().Sub(start), statusCode, resp.Err)
	}

	return &resp, nil
}
