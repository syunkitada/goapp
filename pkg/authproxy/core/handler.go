package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syunkitada/goapp/pkg/authproxy/core/home"
	"github.com/syunkitada/goapp/pkg/resource/resource_authproxy"
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

	ws_authorized := handler.Group("/ws")
	ws_authorized.Use(authproxy.WsAuthRequired())
	{
		ws_authorized.GET("/monitor", authproxy.Monitor.Ws)
	}

	resourceAuthproxy := resource_authproxy.New(authproxy.conf)
	homeSrv := home.New(authproxy.conf)

	authorized := handler.Group("/")
	authorized.Use(authproxy.AuthRequired())
	{
		authorized.POST("/dashboard/logout", authproxy.Dashboard.Logout)
		authorized.POST("/dashboard/state", authproxy.Dashboard.GetState)
		authorized.POST("/Home", homeSrv.Action)
		authorized.POST("/Chat", homeSrv.Chat)
		authorized.POST("/Wiki", homeSrv.Wiki)
		authorized.POST("/Resource.Physical", resourceAuthproxy.PhysicalAction)
		authorized.POST("/Resource.Virtual", resourceAuthproxy.VirtualAction)
		authorized.POST("/monitor", authproxy.Monitor.Action)
	}

	return handler
}
