package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/model"
	"github.com/syunkitada/goapp/pkg/model/model_api"

	"github.com/syunkitada/goapp/pkg/util"
	"net/http"
)

var tokenMap = map[string]string{}

func IssueToken(c *gin.Context) {
	var authRequest model.AuthRequest
	c.Bind(&authRequest)
	glog.Info(authRequest)

	if token, err := model_api.IssueToken(&authRequest); err != nil {
		glog.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		if token != "" {
			tokenMap[token] = token
			glog.Info(tokenMap)
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Invalid authRequest",
			})
		}
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenAuthRequest model.TokenAuthRequest
		c.Bind(&tokenAuthRequest)
		glog.Info(tokenAuthRequest)

		if val, ok := tokenMap[tokenAuthRequest.Token]; ok {
			c.Set("AuthorizedUser", val)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authRequest",
			})
			c.Abort()
		}

		claims, err := util.ParseToken(tokenAuthRequest)
		glog.Info(err)
		glog.Info(claims)
		c.Set("AuthUser", claims["username"])
	}
}
