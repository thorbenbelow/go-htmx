package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("public/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{ "msg": "Hello World"})
	})
	log.Fatal(router.Run(":8080"))

}
