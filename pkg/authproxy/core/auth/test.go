package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/syunkitada/goapp/pkg/authproxy/model"
)

type ResponseIssueToken struct {
	Token string
}

func (auth *Auth) TestIssueToken(t *testing.T) *ResponseIssueToken {
	handler := auth.Conf.Authproxy.TestHandler

	authRequest := model.AuthRequest{
		Username: auth.Conf.Admin.Username,
		Password: auth.Conf.Admin.Password,
	}
	authRequestJson, marshalErr := json.Marshal(authRequest)
	if marshalErr != nil {
		t.Fatalf("Failed json.Marshal: %v", marshalErr)
	}

	req, newRequestErr := http.NewRequest("POST", "/token", bytes.NewBuffer(authRequestJson))
	if newRequestErr != nil {
		t.Fatalf("Failed http.NewRequest: %v", newRequestErr)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var responseIssueToken ResponseIssueToken
	unmarshalErr := json.Unmarshal(w.Body.Bytes(), &responseIssueToken)
	if unmarshalErr != nil {
		t.Fatalf("Failed json.Unmarshal: %v", unmarshalErr)
	}

	return &responseIssueToken
}
