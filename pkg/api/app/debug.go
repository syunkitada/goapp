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
	glog.Info(c)
	if authUser, ok := c.Get("AuthUser"); ok {
		c.JSON(200, gin.H{
			"user":    authUser,
			"message": "Pong",
		})
	} else {
		c.JSON(401, gin.H{
			"message": "Invalid User Pong",
		})
	}
}
