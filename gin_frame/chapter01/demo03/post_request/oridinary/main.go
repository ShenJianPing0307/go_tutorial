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
	fmt.Printf("用户名：%s, 密码：%s", username, password)
	ctx.String(http.StatusOK, "提交成功！")
}

func GetIndex1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index1.html", nil)
}

func PostIndex1(ctx *gin.Context) {
	userMap := ctx.PostFormMap("user")
	fmt.Println(userMap)
	ctx.String(http.StatusOK, "提交成功！")
}

func GetIndex2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index2.html", nil)
}

func PostIndex2(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.DefaultPostForm("age", "0")
	hobby := ctx.PostFormArray("hobby")
	fmt.Printf("用户名：%s, 密码：%s, 年龄：%s, 爱好：%s", username, password, age, hobby)
	ctx.String(http.StatusOK, "提交成功！")
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	// ctx.PostForm
	router.GET("/get_index", GetIndex)
	router.POST("/post_index", PostIndex)
	// ctx.PostFormMap
	router.GET("/get_index1", GetIndex1)
	router.POST("/post_index1", PostIndex1)
	// ctx.DefaultPostForm、ctx.PostFormArray
	router.GET("/get_index2", GetIndex2)
	router.POST("/post_index2", PostIndex2)

	router.Run(":8080")
}
