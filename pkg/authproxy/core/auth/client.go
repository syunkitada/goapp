package auth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
)

type ResponseIssueToken struct {
	Token string
}

func (auth *Auth) CtlIssueToken() (*ResponseIssueToken, error) {
	authRequest := authproxy_model.AuthRequest{
		Username: auth.conf.Ctl.Username,
		Password: auth.conf.Ctl.Password,
	}
	authRequestJson, marshalErr := json.Marshal(authRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}

	req, newRequestErr := http.NewRequest("POST", auth.conf.Ctl.ApiUrl+"/token", bytes.NewBuffer(authRequestJson))
	if newRequestErr != nil {
		return nil, newRequestErr
	}

	var responseIssueToken ResponseIssueToken
	var body []byte
	var statusCode int
	if auth.conf.Default.EnableTest {
		handler := auth.conf.Authproxy.HttpServer.TestHandler
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		body = w.Body.Bytes()
		statusCode = w.Code
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Transport: tr,
		}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		var readAllErr error
		body, readAllErr = ioutil.ReadAll(resp.Body)
		if readAllErr != nil {
			return nil, readAllErr
		}
		statusCode = resp.StatusCode
	}

	if err := json.Unmarshal(body, &responseIssueToken); err != nil {
		return nil, err
	}

	if statusCode != 200 {
		return &responseIssueToken, fmt.Errorf("Invalid StatusCode: %v", statusCode)
	}

	return &responseIssueToken, nil
}
