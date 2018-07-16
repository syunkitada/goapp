package api

import (
	// "github.com/golang/glog"
	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/api/app"
	"net/http"
)

func NewHandler() http.Handler {
	handler := gin.Default()

	handler.POST("/token", app.IssueToken)
	handler.GET("/ping", app.Ping)

	authorized := handler.Group("/")
	authorized.Use(app.AuthRequired())
	{
		authorized.GET("/authtest", app.AuthTest)
	}

	return handler
}
