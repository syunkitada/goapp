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
		service:      conf.Service,
		endpoints:    conf.Endpoints,
	}

	return client
}

type Query struct {
	Name string
	Data interface{}
}

func (client *Client) Request(tctx *logger.TraceContext, service string, queries []Query, resp interface{}, requiredAuth bool) error {
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
		Service: service,
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
		var baseResponse base_model.Response
		if err = json.Unmarshal(body, &baseResponse); err != nil {
			return fmt.Errorf("Invalid StatusCode: got=%d, want=%d", statusCode, 200)
		}
		return fmt.Errorf("Invalid StatusCode: got=%d, want=%d, err=%v", statusCode, 200, baseResponse.ResultMap)
	}

	return nil
}

type LoginResponse struct {
	base_model.Response
	ResultMap LoginResultMap
}

type LoginResultMap struct {
	Login LoginResult
}

type LoginResult struct {
	Code  uint8
	Error string
	Data  base_spec.LoginData
}

func (client *Client) Login(tctx *logger.TraceContext, input *base_spec.Login) (data *base_spec.LoginData, err error) {
	queries := []Query{Query{Name: "Login", Data: input}}
	var res LoginResponse
	err = client.Request(tctx, "Auth", queries, &res, false)
	if err != nil {
		return
	}
	result := res.ResultMap.Login
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}
	data = &result.Data
	return
}

type GetServiceIndexResponse struct {
	base_model.Response
	ResultMap GetServiceIndexResultMap
}
type GetServiceIndexResultMap struct {
	GetServiceIndex GetServiceIndexResult
}

type GetServiceIndexResult struct {
	Code  uint8
	Error string
	Data  base_spec.GetServiceIndexData
}

func (client *Client) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, err error) {
	queries := []Query{Query{Name: "GetServiceIndex", Data: input}}
	var res GetServiceIndexResponse
	err = client.Request(tctx, input.Name, queries, &res, false)
	if err != nil {
		return
	}

	result := res.ResultMap.GetServiceIndex
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}

type UpdateServiceResponse struct {
	base_model.Response
	ResultMap UpdateServiceResultMap
}

type UpdateServiceResultMap struct {
	UpdateService UpdateServiceResult
}

type UpdateServiceResult struct {
	Code  uint8
	Error string
	Data  base_spec.UpdateServiceData
}

func (client *Client) UpdateServices(tctx *logger.TraceContext, queries []Query) (data *base_spec.UpdateServiceData, err error) {
	var res UpdateServiceResponse
	err = client.Request(tctx, "Auth", queries, &res, false)
	if err != nil {
		return
	}
	result := res.ResultMap.UpdateService
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}
