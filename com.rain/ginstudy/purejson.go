package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//按照unicode编码
	router.GET("/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"html": "<b>hello world!</b>",
		})
	})
	//按照当前字面编码
	router.GET("/purejson", func(context *gin.Context) {
		context.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello world!</b>",
		})
	})
	router.Run(":8080")
}
