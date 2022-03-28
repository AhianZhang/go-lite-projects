package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(ctx *gin.Context) {

		file, err := ctx.FormFile("file")
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println(file.Filename)
		dst := "./" + file.Filename
		ctx.SaveUploadedFile(file, dst)
		ctx.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))
	})
	router.Run(":8888")
}
