package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func SomeHandler1(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	//c.ShouldBind 使用了 c.Request.Body，不可重用
	if errA := c.ShouldBindWith(&objA, binding.JSON); errA == nil {
		c.String(200, `the body should be formA`)
	} else if errB := c.ShouldBindWith(&objB, binding.JSON); errB == nil {
		c.String(200, `the body should be formB`)
	} else {
		return
	}
}
