package authproxy_client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

// AuthproxyClient is http client for authproxy
type AuthproxyClient struct {
	conf         *config.Config
	httpClient   *http.Client
	localHandler http.Handler
	apiUrl       string
	serviceName  string
}

func New(conf *config.Config, authproxy *core.Authproxy) *AuthproxyClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Transport: tr,
	}

	client := &AuthproxyClient{
		conf:       conf,
		httpClient: httpClient,
		apiUrl:     conf.Ctl.ApiUrl,
	}

	if conf.Default.EnableTest {
		client.localHandler = authproxy.NewHandler()
	}

	return client
}

func (client *AuthproxyClient) Request(tctx *logger.TraceContext, responseLogin *ResponseLogin, action string, reqData interface{}, resp interface{}) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var reqDataJson []byte
	if reqDataJson, err = json.Marshal(reqData); err != nil {
		return err
	}
	fmt.Println(reqDataJson)

	req := authproxy_model.TokenAuthRequest{
		Token: responseLogin.Token,
		Action: authproxy_model.ActionRequest{
			ProjectName: client.conf.Ctl.Project,
			ServiceName: client.serviceName,
			Queries:     []authproxy_model.Query{},
		},
	}

	return client.request(tctx, client.serviceName, req, resp)
}

func (client *AuthproxyClient) request(tctx *logger.TraceContext, path string, req interface{}, resp interface{}) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var reqJson []byte
	if reqJson, err = json.Marshal(req); err != nil {
		return err
	}

	var httpReq *http.Request
	if httpReq, err = http.NewRequest(
		"POST", client.apiUrl+"/"+path, bytes.NewBuffer(reqJson)); err != nil {
		return err
	}

	var body []byte
	var statusCode int
	if client.localHandler != nil {
		w := httptest.NewRecorder()
		client.localHandler.ServeHTTP(w, httpReq)
		body = w.Body.Bytes()
		statusCode = w.Code
	} else {
		httpResp, err := client.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer func() { err = httpResp.Body.Close() }()
		body, err = ioutil.ReadAll(httpResp.Body)
		if err != nil {
			return err
		}
		statusCode = httpResp.StatusCode
	}

	if err = json.Unmarshal(body, &resp); err != nil {
		return err
	}

	if statusCode != 200 {
		return fmt.Errorf("Invalid StatusCode: get=%v, want=%v", statusCode, 200)
	}

	return nil
}

func (client *AuthproxyClient) Action(tctx *logger.TraceContext, token string, serviceName string, queries []authproxy_model.Query) (*authproxy_model.ActionResponse, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: client.conf.Ctl.Project,
			ServiceName: serviceName,
			Queries:     queries,
		},
	}

	var resp authproxy_model.ActionResponse
	err = client.request(tctx, serviceName, &req, &resp)
	return &resp, err
}
