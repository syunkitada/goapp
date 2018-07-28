package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
	})
}

func AuthTest(c *gin.Context) {
	username, usernameOk := c.Get("Username")
	roleName, roleNameOk := c.Get("RoleName")
	if !usernameOk || !roleNameOk {
		c.JSON(500, gin.H{
			"message": "Invalid request",
		})
		return
	}

	glog.Info(username, roleName)

	c.JSON(200, gin.H{
		"user":    username,
		"role":    roleName,
		"message": "Pong",
	})
}
