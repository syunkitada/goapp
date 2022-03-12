package base_client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_protocol"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type Client struct {
	httpClient     *http.Client
	localHandler   http.Handler
	token          string
	isRefreshToken bool
	service        string
	user           string
	password       string
	project        string
	endpoints      []string
}

func NewClient(conf *base_config.ClientConfig) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.TlsInsecureSkipVerify},
	}
	httpClient := &http.Client{
		Transport: tr,
	}

	isRefreshToken := false
	if conf.User != "" && conf.Password != "" {
		isRefreshToken = true
	}

	client := &Client{
		httpClient:     httpClient,
		localHandler:   conf.LocalHandler,
		user:           conf.User,
		password:       conf.Password,
		project:        conf.Project,
		service:        conf.Service,
		token:          conf.Token,
		endpoints:      conf.Endpoints,
		isRefreshToken: isRefreshToken,
	}

	return client
}

func (client *Client) SetProject(project string) {
	client.project = project
}

type Query struct {
	Name string
	Data interface{}
}

func (client *Client) Request(tctx *logger.TraceContext, service string, queries []Query, resp interface{}, requiredAuth bool) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	reqQueries := []base_protocol.ReqQuery{}
	var queryBytes []byte
	for _, query := range queries {
		if queryBytes, err = json.Marshal(query.Data); err != nil {
			return err
		} else {
			reqQueries = append(reqQueries, base_protocol.ReqQuery{
				Name: query.Name,
				Data: string(queryBytes),
			})
		}
	}

	req := base_protocol.Request{
		Tctx:    tctx,
		Service: service,
		Project: client.project,
		Queries: reqQueries,
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
		if requiredAuth {
			if err = client.Auth(tctx); err != nil {
				return err
			}
			httpReq.Header.Add("X-Auth-Token", client.token)
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
			if requiredAuth {
				if err = client.Auth(tctx); err != nil {
					return err
				}
				httpReq.Header.Add("X-Auth-Token", client.token)
			}

			httpResp, err = client.httpClient.Do(httpReq)
			if err != nil {
				return err
			} else {
				break
			}
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
		var baseResponse base_protocol.Response
		if err = json.Unmarshal(body, &baseResponse); err != nil {
			return fmt.Errorf("Invalid StatusCode: got=%d, want=%d", statusCode, 200)
		}
		return fmt.Errorf("Invalid StatusCode: got=%d, want=%d, err=%v", statusCode, 200, baseResponse.ResultMap)
	}

	return nil
}

func (client *Client) RequestWs(tctx *logger.TraceContext, service string, queries []Query, resp interface{}, requiredAuth bool) (wsConn *websocket.Conn, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	reqQueries := []base_protocol.ReqQuery{}
	var queryBytes []byte
	for _, query := range queries {
		if queryBytes, err = json.Marshal(query.Data); err != nil {
			return
		} else {
			reqQueries = append(reqQueries, base_protocol.ReqQuery{
				Name: query.Name,
				Data: string(queryBytes),
			})
		}
	}

	req := base_protocol.Request{
		Tctx:    tctx,
		Service: service,
		Project: client.project,
		Queries: reqQueries,
	}

	var reqJson []byte
	if reqJson, err = json.Marshal(req); err != nil {
		return
	}

	var body []byte
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if client.localHandler != nil {
		fmt.Println("TODO not supported")
	} else {
		for _, target := range client.endpoints {
			endpoint := strings.Replace(target, "http", "ws", -1) + "/ws"
			header := http.Header{}
			if requiredAuth {
				if err = client.Auth(tctx); err != nil {
					return
				}

				header.Add("X-Auth-Token", client.token)
			}
			wsConn, _, err = dialer.Dial(endpoint, header)
			if err != nil {
				logger.Errorf(tctx, err, "Failed dial to %s", endpoint)
				continue
			}
			break
		}

		if err = wsConn.WriteMessage(websocket.TextMessage, reqJson); err != nil {
			logger.Warningf(tctx, "Failed WriteMessage: %s", err.Error())
			return
		}

	}

	_, body, err = wsConn.ReadMessage()
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	return
}

func (client *Client) Auth(tctx *logger.TraceContext) (err error) {
	if client.token != "" {
		if !client.isRefreshToken {
			return
		}
		splitedToken := strings.Split(client.token, ".")
		if decoded, tmpErr := jwt.DecodeSegment(splitedToken[1]); tmpErr != nil {
			logger.Warningf(tctx, "Failed decode token: %s", tmpErr.Error())
		} else {
			var claim jwt.StandardClaims
			if tmpErr := json.Unmarshal(decoded, &claim); tmpErr != nil {
				logger.Warningf(tctx, "Failed unmarshal token: %s", tmpErr.Error())
			} else {
				if time.Now().Add(60 * time.Second).Before(time.Unix(claim.ExpiresAt, 0)) {
					return
				}
			}
		}
	}

	var data *base_spec.LoginData
	data, err = client.Login(tctx, &base_spec.Login{
		User:     client.user,
		Password: client.password,
	})
	if err != nil {
		return
	}
	client.token = data.Token
	return
}

type LoginResponse struct {
	base_protocol.Response
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

	if res.Code >= 100 || res.Error != "" {
		err = error_utils.NewInvalidResponseError(res.Code, res.Error)
		return
	}

	result := res.ResultMap.Login
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}
	data = &result.Data
	client.token = data.Token
	return
}

type GetServiceIndexResponse struct {
	base_protocol.Response
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
	err = client.Request(tctx, input.Name, queries, &res, true)
	if err != nil {
		return
	}

	if res.Code >= 100 || res.Error != "" {
		err = error_utils.NewInvalidResponseError(res.Code, res.Error)
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

type GetProjectServiceIndexResponse struct {
	base_protocol.Response
	ResultMap GetProjectServiceIndexResultMap
}

type GetProjectServiceIndexResultMap struct {
	GetProjectServiceIndex GetServiceIndexResult
}

func (client *Client) GetProjectServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, err error) {
	queries := []Query{Query{Name: "GetProjectServiceIndex", Data: input}}
	var res GetProjectServiceIndexResponse
	err = client.Request(tctx, input.Name, queries, &res, true)
	if err != nil {
		return
	}

	if res.Code >= 100 || res.Error != "" {
		err = error_utils.NewInvalidResponseError(res.Code, res.Error)
		return
	}

	result := res.ResultMap.GetProjectServiceIndex
	if result.Code != base_const.CodeOk || result.Error != "" {
		err = error_utils.NewInvalidResponseError(result.Code, result.Error)
		return
	}

	data = &result.Data
	return
}

type UpdateServiceResponse struct {
	base_protocol.Response
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
