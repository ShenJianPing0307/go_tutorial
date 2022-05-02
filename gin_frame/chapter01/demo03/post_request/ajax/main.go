package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func PostIndex(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username, password)
	data := map[string]interface{}{
		"code":    2000,
		"message": "成功",
	}
	ctx.JSON(http.StatusOK, data)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")
	router.Static("/static", "static")

	router.GET("/get_index", GetIndex)
	router.POST("/post_index", PostIndex)

	router.Run(":8080")
}
