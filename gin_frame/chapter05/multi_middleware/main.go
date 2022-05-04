package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddlewareA() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("MiddlewareA before request")
		// before request
		context.Next()
		// after request
		fmt.Println("MiddlewareA after request")
	}
}

func MiddlewareB() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("MiddlewareB before request")
		// before request
		context.Next()
		// after request
		fmt.Println("MiddlewareB after request")
	}
}

func main() {

	router := gin.Default()

	router.GET("/login", Login)
	router.Use(MiddlewareA(), MiddlewareB())

	router.Run(":8080")

}

func Login(context *gin.Context) {
	context.String(http.StatusOK, "登陆!")
}
