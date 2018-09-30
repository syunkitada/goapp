package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (auth *Auth) TestAuthAndIssueToken(t *testing.T) {
	handler := auth.Conf.Authproxy.TestHandler

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
