package base_client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type Client struct {
	httpClient   *http.Client
	localHandler http.Handler
	token        string
	service      string
	endpoints    []string
}

func NewClient(conf *base_config.ClientConfig) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.TlsInsecureSkipVerify},
	}
	httpClient := &http.Client{
		Transport: tr,
	}

	client := &Client{
		httpClient:   httpClient,
		localHandler: conf.LocalHandler,
		endpoints:    conf.Endpoints,
	}

	return client
}

type Query struct {
	Name string
	Data interface{}
}

func (client *Client) Request(tctx *logger.TraceContext, queries []Query, resp interface{}, requiredAuth bool) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	reqQueries := []base_model.ReqQuery{}
	var queryBytes []byte
	for _, query := range queries {
		if queryBytes, err = json.Marshal(query.Data); err != nil {
			return err
		} else {
			reqQueries = append(reqQueries, base_model.ReqQuery{
				Name: query.Name,
				Data: string(queryBytes),
			})
		}
	}

	req := base_model.Request{
		Tctx:    tctx,
		Service: "Auth",
		Queries: reqQueries,
	}
	if requiredAuth {
		req.Token = client.token
	}

	var reqJson []byte
	if reqJson, err = json.Marshal(req); err != nil {
		return err
	}

	var body []byte
	var statusCode int
	var httpReq *http.Request
	reqBuffer := bytes.NewBuffer(reqJson)
	if client.localHandler != nil {
		if httpReq, err = http.NewRequest(
			"POST", "http://127.0.0.1/q", reqBuffer); err != nil {
			return err
		}
		w := httptest.NewRecorder()
		client.localHandler.ServeHTTP(w, httpReq)
		body = w.Body.Bytes()
		statusCode = w.Code
	} else {
		var httpResp *http.Response
		for _, target := range client.endpoints {
			if httpReq, err = http.NewRequest("POST", target+"/q", reqBuffer); err != nil {
				return err
			}

			httpResp, err = client.httpClient.Do(httpReq)
			if err != nil {
				return err
			}
			break
		}

		defer func() {
			if tmpErr := httpResp.Body.Close(); tmpErr != nil {
				logger.Errorf(tctx, err, "Failed httpResp.Body.Close()")
			}
		}()
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		statusCode = httpResp.StatusCode
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return err
	}

	if statusCode != 200 {
		return fmt.Errorf("Invalid StatusCode: get=%v, want=%v", statusCode, 200)
	}

	return nil
}

type LoginResponse struct {
	base_model.Response
	Data LoginResponseData
}

type LoginResponseData struct {
	Login base_spec.LoginData
}

func (client *Client) Login(tctx *logger.TraceContext, input *base_spec.Login) (data *base_spec.LoginData, err error) {
	queries := []Query{Query{Name: "Login", Data: input}}
	var reply LoginResponse
	err = client.Request(tctx, queries, &reply, false)
	if err != nil {
		return
	}
	if reply.Code != base_const.CodeOk || reply.Error != "" {
		err = error_utils.NewInvalidResponseError(reply.Code, reply.Error)
		return
	}
	data = &reply.Data.Login
	return
}

type GetServiceIndexResponse struct {
	base_model.Response
	Data GetServiceIndexResponseData
}

type GetServiceIndexResponseData struct {
	GetServiceIndex base_spec.GetServiceIndexData
}

func (client *Client) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, err error) {
	queries := []Query{Query{Name: "GetServiceIndex", Data: input}}
	var reply GetServiceIndexResponse
	err = client.Request(tctx, queries, &reply, false)
	if err != nil {
		return
	}
	if reply.Code != base_const.CodeOk || reply.Error != "" {
		err = error_utils.NewInvalidResponseError(reply.Code, reply.Error)
		return
	}
	data = &reply.Data.GetServiceIndex
	return
}

type UpdateServiceResponse struct {
	base_model.Response
	Data UpdateServiceResponseData
}

type UpdateServiceResponseData struct {
	UpdateService base_spec.UpdateServiceData
}

func (client *Client) UpdateServices(tctx *logger.TraceContext, queries []Query) (data *base_spec.UpdateServiceData, err error) {
	var reply UpdateServiceResponse
	err = client.Request(tctx, queries, &reply, false)
	if err != nil {
		return
	}
	if reply.Code != base_const.CodeOk || reply.Error != "" {
		err = error_utils.NewInvalidResponseError(reply.Code, reply.Error)
		return
	}
	data = &reply.Data.UpdateService
	return
}
