package main

import "github.com/gin-gonic/gin"

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(context *gin.Context) {
		var form LoginForm
		if context.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				context.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				context.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	router.Run(":8080")
}
