package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/model"
	"github.com/syunkitada/goapp/pkg/model/model_api"

	"net/http"
)

func IssueToken(c *gin.Context) {
	var authRequest model.AuthRequest
	c.Bind(&authRequest)
	glog.Info(authRequest)

	if token, err := model_api.IssueToken(&authRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{})
		}
	}
}
