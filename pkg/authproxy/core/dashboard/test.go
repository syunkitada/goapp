package dashboard

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/core/auth"
)

type ResponseLogin struct {
	Name      string
	Authority authproxy_model.UserAuthority
}

func (dashboard *Dashboard) TestLogin(t *testing.T) *ResponseLogin {
	handler := dashboard.Conf.Authproxy.HttpServer.TestHandler

	authRequest := authproxy_model.AuthRequest{
		Username: dashboard.Conf.Admin.Username,
		Password: dashboard.Conf.Admin.Password,
	}
	authRequestJson, marshalErr := json.Marshal(authRequest)
	if marshalErr != nil {
		t.Fatalf("Failed json.Marshal: %v", marshalErr)
	}

	req, newRequestErr := http.NewRequest("POST", "/dashboard/login", bytes.NewBuffer(authRequestJson))
	if newRequestErr != nil {
		t.Fatalf("Failed http.NewRequest: %v", newRequestErr)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseLogin ResponseLogin
	unmarshalErr := json.Unmarshal(w.Body.Bytes(), &responseLogin)
	if unmarshalErr != nil {
		t.Fatalf("Failed json.Unmarshal: %v", unmarshalErr)
	}

	return &responseLogin
}

func (dashboard *Dashboard) TestGetState(t *testing.T, token *auth.ResponseIssueToken) *ResponseLogin {
	handler := dashboard.Conf.Authproxy.HttpServer.TestHandler

	request := authproxy_model.TokenAuthRequest{}
	requestJson, marshalErr := json.Marshal(request)
	if marshalErr != nil {
		t.Fatalf("Failed json.Marshal: %v", marshalErr)
	}

	req, newRequestErr := http.NewRequest("POST", "/dashboard/state", bytes.NewBuffer(requestJson))
	if newRequestErr != nil {
		t.Fatalf("Failed http.NewRequest: %v", newRequestErr)
	}

	cookie := http.Cookie{Name: "token", Value: token.Token}
	req.AddCookie(&cookie)

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseLogin ResponseLogin
	unmarshalErr := json.Unmarshal(w.Body.Bytes(), &responseLogin)
	if unmarshalErr != nil {
		t.Fatalf("Failed json.Unmarshal: %v", unmarshalErr)
	}

	return &responseLogin
}
