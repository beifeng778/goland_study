package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		//设置example变量
		context.Set("example", "12345")

		//请求前
		context.Next()

		//请求后
		latency := time.Since(t)
		log.Print(latency)

		//获取发送的status
		status := context.Writer.Status()
		log.Println(status)
	}
}

func main() {
	route := gin.New()
	route.Use(Logger())

	route.GET("/test", func(context *gin.Context) {
		example := context.MustGet("example").(string)

		//打印：”12345“
		log.Println(example)
	})
	route.Run()
}
