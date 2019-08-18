package base_app

import (
	"github.com/gin-gonic/gin"
)

func (app *BaseApp) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var req Request
		// traceId := c.GetString("TraceId")
		// tctx := logger.NewTraceContextWithTraceId(traceId, app.host, app.name)

		// if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"Err": "Invalid Auth Request: You need to login.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// value, cookieErr := c.Cookie("token")
		// if cookieErr == nil {
		// 	tokenAuthRequest.Token = value
		// }

		// claims, err := app.Token.ParseToken(tokenAuthRequest)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"Err": "Invalid Auth Token: You need to login.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// username := claims["Username"].(string)
		// userAuthority, getUserAuthorityErr := app.BaseAppModelApi.GetUserAuthority(
		// 	tctx, username, &tokenAuthRequest.Action)
		// if getUserAuthorityErr != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"Err": "Invalid AuthAction: This request is not allowed.",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// c.Set("Username", claims["Username"])
		// c.Set("UserAuthority", userAuthority)
		// c.Set("Action", tokenAuthRequest.Action)
		return
	}
}

func (app *BaseApp) WsAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.Header["X-Auth-Token"]
		// projectName := c.Request.Header["X-Auth-Project"]
		// serviceName := c.Request.Header["X-Auth-Service"]
		// // actionName := c.Request.Header["X-Auth-Action"]

		// traceId := c.GetString("TraceId")
		// tctx := logger.NewTraceContextWithTraceId(traceId, app.host, app.name)

		// tokenAuthRequest := app_model.TokenAuthRequest{
		// 	Token: token[0],
		// 	Action: app_model.ActionRequest{
		// 		ProjectName: projectName[0],
		// 		ServiceName: serviceName[0],
		// 		// Name:        actionName[0],
		// 	},
		// }

		// value, cookieErr := c.Cookie("token")
		// if cookieErr == nil {
		// 	tokenAuthRequest.Token = value
		// }

		// claims, err := app.Token.ParseToken(tokenAuthRequest)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"Err": "Invalid Auth Token",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// username := claims["Username"].(string)
		// userAuthority, getUserAuthorityErr := app.BaseAppModelApi.GetUserAuthority(
		// 	tctx, username, &tokenAuthRequest.Action)
		// if getUserAuthorityErr != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"Err": "Invalid Auth Action",
		// 	})
		// 	c.Abort()
		// 	return
		// }

		// c.Set("Username", claims["Username"])
		// c.Set("UserAuthority", userAuthority)
		// c.Set("Action", tokenAuthRequest.Action)
		return
	}
}
