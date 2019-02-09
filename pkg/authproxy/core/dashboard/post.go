package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (dashboard *Dashboard) Login(c *gin.Context) {
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

	tctx := logger.NewTraceContextWithTraceId(traceId, dashboard.host, dashboard.name)
	token, err := dashboard.Token.AuthAndIssueToken(tctx, &authRequest)
	if err != nil {
		glog.Error("Failed AuthAndIssueToken", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed IssueToken",
		})
		c.Abort()
		return
	}

	userAuthority, getUserAuthorityErr := dashboard.AuthproxyModelApi.GetUserAuthority(
		tctx, authRequest.Username, &authRequest.Action)
	if getUserAuthorityErr != nil {
		glog.Error(getUserAuthorityErr)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
	}

	c.SetCookie("token", token, 3600, "/", "192.168.10.103", true, true)
	c.JSON(http.StatusOK, gin.H{
		"Name":      authRequest.Username,
		"Authority": userAuthority,
	})

	glog.Info("Success Login: ", authRequest.Username)
}

func (dashboard *Dashboard) Logout(c *gin.Context) {
	username, usernameOk := c.Get("Username")
	_, userAuthorityOk := c.Get("UserAuthority")
	if !usernameOk || !userAuthorityOk {
		c.JSON(500, gin.H{
			"error": "Invalid request",
		})
		return
	}

	c.SetCookie("token", "", 0, "/", "192.168.10.103", true, true)
	c.JSON(200, gin.H{})

	glog.Info("Success Logout: ", username)
}
