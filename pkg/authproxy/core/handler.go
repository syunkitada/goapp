package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (authproxy *Authproxy) NewHandler() http.Handler {
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(authproxy.ValidateHeaders())

	handler.POST("/token", authproxy.IssueToken)
	handler.GET("/health", authproxy.Health)
	handler.GET("/health-grpc", authproxy.HealthGrpc)

	authorized := handler.Group("/")
	authorized.Use(authproxy.AuthRequired())
	{
		authorized.GET("/auth-health", authproxy.AuthHealth)
	}

	return handler
}
