package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (authproxy *Authproxy) NewHandler() http.Handler {
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	if !authproxy.Conf.Default.TestMode {
		handler.Use(authproxy.ValidateHeaders())
	}

	handler.POST("/token", authproxy.Auth.IssueToken)
	handler.POST("/dashboard/login", authproxy.Dashboard.Login)
	handler.GET("/health", authproxy.Health)
	handler.GET("/health-grpc", authproxy.HealthGrpc)

	authorized := handler.Group("/")
	authorized.Use(authproxy.AuthRequired())
	{
		authorized.POST("/dashboard/logout", authproxy.Dashboard.Logout)
		authorized.GET("/dashboard/state", authproxy.Dashboard.GetState)
		authorized.GET("/auth-health", authproxy.AuthHealth)
	}

	return handler
}
