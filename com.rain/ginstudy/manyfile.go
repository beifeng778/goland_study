package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			dst := "./" + file.Filename
			err := context.SaveUploadedFile(file, dst)
			if err != nil {
				return
			}
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
