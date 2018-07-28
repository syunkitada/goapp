package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/model"
	"github.com/syunkitada/goapp/pkg/model/model_api"

	"github.com/syunkitada/goapp/pkg/util"
	"net/http"
)

func IssueToken(c *gin.Context) {
	var authRequest model.AuthRequest

	if err := c.ShouldBindWith(&authRequest, binding.JSON); err != nil {
		glog.Warningf("Invalid AuthRequest: Failed ShouldBindJSON: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid AuthRequest",
		})
		c.Abort()
		return
	}

	if token, err := model_api.IssueToken(&authRequest); err != nil {
		glog.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		c.Abort()
		return
	} else {
		if token != "" {
			glog.Info("Success Login: ", authRequest)
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid authRequest",
			})
			c.Abort()
			return
		}
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenAuthRequest model.TokenAuthRequest
		c.Bind(&tokenAuthRequest)

		claims, err := util.ParseToken(tokenAuthRequest)
		if err != nil {
			glog.Warning("Invalid AuthRequest: Failed ParseToken")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid AuthRequest",
			})
			c.Abort()
		}

		c.Set("Username", claims["Username"])
		c.Set("RoleName", claims["RoleName"])
		c.Set("ProjectName", claims["ProjectName"])
		c.Set("ProjectRoleName", claims["ProjectRoleName"])
	}
}
