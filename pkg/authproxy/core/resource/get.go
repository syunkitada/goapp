package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func (resource *Resource) GetStatus(c *gin.Context) {
	status, err := Resource.ResourceClient.Status()
	if err != nil {
		glog.Error("Failed HealthClient.Status", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
		c.Abort()
	}
	glog.Info(status)

	c.JSON(200, gin.H{
		"message": "Health",
	})
}
