package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	c := gin.Default()
	c.GET("/jsonp", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSONP(http.StatusOK, data)
	})
	c.Run(":8080")
}
