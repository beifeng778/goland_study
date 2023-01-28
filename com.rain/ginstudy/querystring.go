package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/welcome", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Guest")
		lastname := context.Query("lastname")
		context.String(200, "Hello %s %s", firstname, lastname)
	})
	router.Run()
}
