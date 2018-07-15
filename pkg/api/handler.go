package api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/syunkitada/goapp/pkg/api/app"
	"github.com/syunkitada/goapp/pkg/api/services"
	"net/http"
)

func NewHandler() http.Handler {
	handler := gin.Default()

	handler.POST("/token", app.IssueToken)
	handler.GET("/ping", services.Ping)

	authorized := handler.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/authtest", services.Ping)
	}

	return handler
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		glog.Info("hoge")
		// TODO validate token
		// Some authorization in Authorization
		// user := Authorization()

		c.Set("AuthorizedUser", "user")
	}
}
