package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/someJSON", func(context *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		context.SecureJSON(http.StatusOK, names)
	})
	router.Run(":8080")
}
