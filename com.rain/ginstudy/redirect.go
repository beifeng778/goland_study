package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//router.GET("/test", func(context *gin.Context) {
	//	context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	//})
	//router.POST("/test1", func(context *gin.Context) {
	//	context.Redirect(http.StatusFound, "/foo")
	//})
	router.GET("/test1", func(context *gin.Context) {
		context.Request.URL.Path = "/test2"
		router.HandleContext(context)
	})
	router.GET("/test2", func(context *gin.Context) {
		context.JSON(200, gin.H{"hello": "world"})
	})
	router.Run()
}
