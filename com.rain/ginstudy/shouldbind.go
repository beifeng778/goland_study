package main

import "github.com/gin-gonic/gin"

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	//c.ShouldBind 使用了 c.Request.Body，不可重用
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(200, `the body should be formA`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(200, `the body should be formB`)
	} else {
		return
	}
}
