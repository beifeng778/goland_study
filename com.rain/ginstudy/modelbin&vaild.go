package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

func main() {
	router := gin.Default()
	// 绑定 JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		context.JSON(200, gin.H{
			"status": "ypu are logged in",
		})
	})

	// 绑定 XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(context *gin.Context) {
		var xml Login
		if err := context.ShouldBindXML(&xml); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if xml.User != "manu" || xml.Password != "123" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}
		context.JSON(200, gin.H{"status": "you are logged in"})
	})

	// 绑定 HTML 表单 (user=manu&password=123)
	router.POST("/loginForm", func(context *gin.Context) {
		var form Login
		// 根据 Content-Type Header 推断使用哪个绑定器
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if form.User != "manu" || form.Password != "123" {
			context.JSON(401, gin.H{"status": "unauthorized"})
		}
		context.JSON(200, gin.H{"status": "you are logged in"})
	})
	router.Run()
}
