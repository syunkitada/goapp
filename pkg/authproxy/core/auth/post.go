package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (auth *Auth) IssueToken(c *gin.Context) {
	var authRequest authproxy_model.AuthRequest
	tmpTraceId, traceIdOk := c.Get("TraceId")
	if !traceIdOk {
		c.JSON(500, gin.H{
			"err": "Invalid request",
		})
		c.Abort()
	}
	traceId := tmpTraceId.(string)

	if err := c.ShouldBindWith(&authRequest, binding.JSON); err != nil {
		glog.Errorf("Failed IssueToken for user=%v: Failed ShouldBindJSON: %v", authRequest.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
		c.Abort()
		return
	}

	token, err := auth.token.AuthAndIssueToken(&authRequest)
	if err != nil {
		glog.Errorf("Failed IssueToken for user=%v: Failed AuthAndIssueToken: %v", authRequest.Username, err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed IssueToken",
		})
		c.Abort()
		return
	}

	logger.TraceInfo(auth.name, traceId, map[string]string{
		"user": authRequest.Username,
		"msg":  "Success IssueToken",
	})
	c.JSON(http.StatusOK, gin.H{
		"Token": token,
	})

	return
}
