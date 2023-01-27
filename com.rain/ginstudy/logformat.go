package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.POST("/foo", func(context *gin.Context) {
		context.JSON(200, "foo")
	})

	router.GET("/bar", func(context *gin.Context) {
		context.JSON(200, "bar")
	})

	router.GET("/status", func(context *gin.Context) {
		context.JSON(200, "ok")
	})

	router.Run(":8080")
}
