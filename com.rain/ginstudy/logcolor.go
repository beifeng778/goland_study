package main

import "github.com/gin-gonic/gin"

func main() {
	//禁用颜色
	//gin.DisableConsoleColor()
	//使用颜色
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	router.Run()
}
