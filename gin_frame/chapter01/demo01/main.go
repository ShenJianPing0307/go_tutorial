package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/index.html", "hello home!")
}

func main() {
	router := gin.Default()
	// 加载模板路径
	router.LoadHTMLGlob("template/**/*")
	// 加载静态文件
	router.Static("/static", "static")

	router.GET("/index", Hello)

	router.Run(":8080")
}
