package genpkg

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

// AuthproxyClient is http client for authproxy
type Client struct {
	httpClient   *http.Client
	localHandler http.Handler
	token        string
	service      string
	targets      []string
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
		targets:      conf.Targets,
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
		for _, target := range client.targets {
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

	if err = json.Unmarshal(body, &resp); err != nil {
		return err
	}

	if statusCode != 200 {
		return fmt.Errorf("Invalid StatusCode: get=%v, want=%v", statusCode, 200)
	}

	return nil
}
