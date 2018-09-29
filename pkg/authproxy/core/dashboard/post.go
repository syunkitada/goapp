package dashboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/core/auth"
	"github.com/syunkitada/goapp/pkg/authproxy/model"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
)

func (dashboard *Dashboard) Login(c *gin.Context) {
	var authRequest model.AuthRequest

	if err := c.ShouldBindWith(&authRequest, binding.JSON); err != nil {
		glog.Warningf("Invalid AuthRequest: Failed ShouldBindJSON: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
		c.Abort()
		return
	}

	token, err := auth.AuthAndIssueToken(&authRequest)
	if err != nil {
		glog.Error("Failed AuthAndIssueToken", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Failed IssueToken",
		})
		c.Abort()
		return
	}

	userAuthority, getUserAuthorityErr := model_api.GetUserAuthority(authRequest.Username)
	if getUserAuthorityErr != nil {
		glog.Error(getUserAuthorityErr)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
	}

	glog.Info("Success Login: ", authRequest)
	c.SetCookie("token", token, 3600, "/", "192.168.10.103", true, true)
	c.JSON(http.StatusOK, gin.H{
		"Name":      authRequest.Username,
		"Authority": userAuthority,
	})

	return
}
