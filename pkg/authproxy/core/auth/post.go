package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (auth *Auth) Login(c *gin.Context) {
	var authRequest authproxy_model.AuthRequest
	traceId := c.GetString("TraceId")

	if err := c.ShouldBindWith(&authRequest, binding.JSON); err != nil {
		glog.Warningf("Invalid AuthRequest: Failed ShouldBindJSON: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
		c.Abort()
		return
	}

	tctx := logger.NewTraceContextWithTraceId(traceId, auth.host, auth.name)
	token, err := auth.token.AuthAndIssueToken(tctx, &authRequest)
	if err != nil {
		glog.Error("Failed AuthAndIssueToken", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed IssueToken",
		})
		c.Abort()
		return
	}

	userAuthority, getUserAuthorityErr := auth.authproxyModelApi.GetUserAuthority(
		tctx, authRequest.Username, &authRequest.Action)
	if getUserAuthorityErr != nil {
		glog.Error(getUserAuthorityErr)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":      authRequest.Username,
		"Authority": userAuthority,
		"Token":     token,
	})
}
