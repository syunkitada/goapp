package core

// import (
// 	"net/http"
// 
// 	"github.com/gin-gonic/gin"
// )
// 
// func (dashboard *Dashboard) NewHandler() http.Handler {
// 	handler := gin.New()
// 	handler.Use(gin.Logger())
// 	handler.Use(gin.Recovery())
// 	handler.Use(dashboard.ValidateHeaders())
// 
// 	handler.StaticFS("/", http.Dir(dashboard.BuildDir))
// 
// 	// authorized := handler.Group("/")
// 	// authorized.Use(dashboard.AuthRequired())
// 	// {
// 	// 	authorized.GET("/auth-health", dashboard.AuthHealth)
// 	// }
// 
// 	return handler
// }
