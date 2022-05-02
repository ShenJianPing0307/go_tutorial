package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"code":    2001,
		"message": "success",
	})

}

func main() {

	router := gin.Default()

	router.GET("/index", Index)

	router.Run(":8080")

}
