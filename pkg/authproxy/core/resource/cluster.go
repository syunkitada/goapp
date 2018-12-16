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

type ResponseGetCluster struct {
	Clusters []resource_api_grpc_pb.Cluster
	TraceId  string
	Err      string
}

type ResponseCreateCluster struct {
	Cluster resource_api_grpc_pb.Cluster
	TraceId string
	Err     string
}

type ResponseUpdateCluster struct {
	Cluster resource_api_grpc_pb.Cluster
	TraceId string
	Err     string
}

type ResponseDeleteCluster struct {
	TraceId string
	Err     string
}

func (resource *Resource) GetCluster(c *gin.Context, rc *ResourceContext) (int, string) {
	var reqData resource_api_grpc_pb.GetClusterRequest
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

	rep, err := resource.resourceApiClient.GetCluster(&reqData)
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
		"Clusters": rep.Clusters,
	})
	return int(rep.StatusCode), rep.Err
}

func (resource *Resource) CtlGetCluster(token string, cluster string, target string) (*ResponseGetCluster, error) {
	reqData := resource_api_grpc_pb.GetClusterRequest{
		Target: target,
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
			Name:        "GetCluster",
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

	var resp ResponseGetCluster
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
