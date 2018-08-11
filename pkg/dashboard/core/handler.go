package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (dashboard *Dashboard) NewHandler() http.Handler {
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(dashboard.ValidateHeaders())

	handler.StaticFS("/static", http.Dir(dashboard.StaticDir))
	handler.LoadHTMLGlob(dashboard.TemplatesDir)

	handler.GET("/", dashboard.Login)
	handler.GET("/login", dashboard.LoginIndex)

	// authorized := handler.Group("/")
	// authorized.Use(dashboard.AuthRequired())
	// {
	// 	authorized.GET("/auth-health", dashboard.AuthHealth)
	// }

	return handler
}
