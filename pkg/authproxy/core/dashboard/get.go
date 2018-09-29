package dashboard

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func (dashboard *Dashboard) GetState(c *gin.Context) {
	username, usernameOk := c.Get("Username")
	userAuthority, userAuthorityOk := c.Get("UserAuthority")
	if !usernameOk || !userAuthorityOk {
		c.JSON(500, gin.H{
			"error": "Invalid request",
		})
		return
	}

	glog.Info("Success AuthHealth: username(%v)", username)
	glog.Info(userAuthority)

	c.JSON(200, gin.H{
		"Name":      username,
		"Authority": userAuthority,
	})
}
