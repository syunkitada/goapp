package core

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"
// 
// 	"github.com/gin-gonic/gin"
// )
// 
// // ValidateHeaders validate http headers
// // SecureHeaders adds secure headers to the API
// // func (a *API) SecureHeaders(next http.Handler) http.Handler {
// // return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// func (dashboard *Dashboard) ValidateHeaders() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Check AllowedHosts
// 		var err error
// 		if len(dashboard.AllowedHosts) > 0 {
// 			isGoodHost := false
// 			for _, allowedHost := range dashboard.AllowedHosts {
// 				if strings.EqualFold(allowedHost, c.Request.Host) {
// 					isGoodHost = true
// 					break
// 				}
// 			}
// 			if !isGoodHost {
// 				c.JSON(http.StatusForbidden, gin.H{
// 					"error": fmt.Sprintf("Bad host name: %s", c.Request.Host),
// 				})
// 				c.Abort()
// 			}
// 		}
// 		// If there was an error, do not continue request
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": fmt.Sprintf("Failed to check allowed hosts"),
// 			})
// 			c.Abort()
// 		}
// 
// 		// Add X-XSS-Protection header
// 		// Enables XSS filtering. Rather than sanitizing the page, the browser will prevent rendering of the page if an attack is detected.
// 		c.Writer.Header().Add("X-XSS-Protection", "1; mode=block")
// 
// 		// Add Content-Type header
// 		// Content type tells the browser what type of content you are sending. If you do not include it, the browser will try to guess the type and may get it wrong.
// 		// w.Header().Add("Content-Type", "application/json")
// 
// 		// Add X-Content-Type-Options header
// 		// Content Sniffing is the inspecting the content of a byte stream to attempt to deduce the file format of the data within it.
// 		// Browsers will do this to try to guess at the content type you are sending.
// 		// By setting this header to “nosniff”, it prevents IE and Chrome from content sniffing a response away from its actual content type. This reduces exposure to drive-by download attacks.
// 		c.Writer.Header().Add("X-Content-Type-Options", "nosniff")
// 
// 		// Prevent page from being displayed in an iframe
// 		c.Writer.Header().Add("X-Frame-Options", "DENY")
// 	}
// }
// 
// // func (dashboard *Dashboard) AuthRequired() gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		var tokenAuthRequest authproxy_model.TokenAuthRequest
// // 		if err := c.Bind(&tokenAuthRequest); err != nil {
// // 			glog.Warning("Invalid AuthRequest: Failed ParseToken")
// // 			c.JSON(http.StatusUnauthorized, gin.H{
// // 				"error": "Invalid AuthRequest",
// // 			})
// // 			c.Abort()
// // 		}
// //
// // 		claims, err := dashboard.ParseToken(tokenAuthRequest)
// // 		if err != nil {
// // 			glog.Warning("Invalid AuthRequest: Failed ParseToken")
// // 			c.JSON(http.StatusUnauthorized, gin.H{
// // 				"error": "Invalid AuthRequest",
// // 			})
// // 			c.Abort()
// // 		}
// //
// // 		c.Set("UserName", claims["UserName"])
// // 		c.Set("RoleName", claims["RoleName"])
// // 		c.Set("ProjectName", claims["ProjectName"])
// // 		c.Set("ProjectRoleName", claims["ProjectRoleName"])
// // 	}
// // }
// 
// // func (dashboard *Dashboard) ParseToken(request authproxy_model.TokenAuthRequest) (jwt.MapClaims, error) {
// // 	token, err := jwt.Parse(request.Token, func(token *jwt.Token) (interface{}, error) {
// // 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// // 			msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// // 			return nil, msg
// // 		}
// // 		return []byte(Conf.Admin.TokenSecret), nil
// // 	})
// //
// // 	if err != nil {
// // 		return nil, err
// // 	}
// //
// // 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// // 		return claims, nil
// // 	}
// //
// // 	return nil, nil
// // }
