package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
)

func (resource *Resource) Action(c *gin.Context) {
	username, usernameOk := c.Get("Username")
	userAuthority, userAuthorityOk := c.Get("UserAuthority")
	tmpAction, actionOk := c.Get("Action")
	if !usernameOk || !userAuthorityOk || !actionOk {
		c.JSON(500, gin.H{
			"error": "Invalid request",
		})
		return
	}

	action := tmpAction.(authproxy_model.ActionRequest)

	glog.Info(username)
	glog.Info(userAuthority)
	glog.Info(action)

	switch action.Name {
	case "GetState":
		glog.Info("itest")
		status, err := resource.ResourceApiClient.Status()
		if err != nil {
			glog.Error("Failed HealthClient.Status", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid AuthRequest",
			})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{
			"message": status,
		})

		return

	}

	// status, err := resource.ResourceClient.Status()
	// if err != nil {
	// 	glog.Error("Failed HealthClient.Status", err)
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Invalid AuthRequest",
	// 	})
	// 	c.Abort()
	// }
	// glog.Info(status)

	c.JSON(200, gin.H{
		"message": "Health",
	})
}
