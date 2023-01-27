package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/long_async", func(context *gin.Context) {
		//创建在 goroutine 中使用的副本
		cCp := context.Copy()
		go func() {
			//用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)

			//请注意您使用的是复制的上下文 "cCp"
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	router.GET("/long_sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + context.Request.URL.Path)
	})
	router.Run(":8080")
}
