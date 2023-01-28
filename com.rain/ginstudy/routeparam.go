package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()
	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(200, "Hello %s", name)
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		action = strings.Replace(action, "/", "", -1)
		message := name + " is " + action
		context.String(200, message)
	})
	router.Run()
}
