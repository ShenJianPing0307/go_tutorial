package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	//router := gin.Default()
	router := gin.New()

	router.GET("/login", Login)

	router.Run(":8080")

}

func Login(context *gin.Context) {
	context.String(http.StatusOK, "登陆!")
}
