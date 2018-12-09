package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (authproxy *Authproxy) NewHandler() http.Handler {
	handler := gin.New()
	handler.Use(authproxy.Logger())
	handler.Use(gin.Recovery())
	if !authproxy.conf.Default.EnableTest {
		handler.Use(authproxy.ValidateHeaders())
	}

	handler.POST("/token", authproxy.Auth.IssueToken)
	handler.POST("/dashboard/login", authproxy.Dashboard.Login)

	authorized := handler.Group("/")
	authorized.Use(authproxy.AuthRequired())
	{
		authorized.POST("/dashboard/logout", authproxy.Dashboard.Logout)
		authorized.POST("/dashboard/state", authproxy.Dashboard.GetState)
		authorized.POST("/resource", authproxy.Resource.Action)
	}

	return handler
}
