package resource

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseGetCluster struct {
	Clusters []resource_api_grpc_pb.Cluster
	Err      string
}

type ResponseGetNode struct {
	Nodes []resource_api_grpc_pb.Node
	Err   string
}

func (resource *Resource) CtlGetCluster(token string) (*ResponseGetCluster, error) {
	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: resource.conf.Ctl.Project,
			ServiceName: "Resource",
			Name:        "GetCluster",
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

func (resource *Resource) CtlGetNode(token string, cluster string, target string) (*ResponseGetNode, error) {
	reqData := resource_api_grpc_pb.GetNodeRequest{
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
			Name:        "GetNode",
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

	var resp ResponseGetNode
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
