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
		//simple file
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		err := context.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}
		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
