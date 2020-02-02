package core

// import (
// 	"net/http"
// 
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/glog"
// )
// 
// func (dashboard *Dashboard) LoginIndex(c *gin.Context) {
// 	if pusher := c.Writer.Pusher(); pusher != nil {
// 		// use pusher.Push() to do server push
// 		options := &http.PushOptions{
// 			Header: http.Header{
// 				"Accept-Encoding": c.Request.Header["Accept-Encoding"],
// 			},
// 		}
// 		if err := pusher.Push("/static/js/app.js", options); err != nil {
// 			glog.Warning("Failed to push: %v", err)
// 		}
// 		glog.Info("push app.js")
// 	} else {
// 		glog.Warning("pusher not supported")
// 	}
// 
// 	glog.Info("DEBUG")
// 
// 	c.HTML(http.StatusOK, "login/index.tmpl", gin.H{
// 		"title": "Posts",
// 	})
// }
// 
// func (dashboard *Dashboard) Login(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Health",
// 	})
// }
