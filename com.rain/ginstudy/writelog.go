package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	//禁用控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()

	//记录到文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})
	router.Run(":8080")
}
