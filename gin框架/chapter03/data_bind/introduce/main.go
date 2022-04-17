package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
}

func Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index.html", nil)

}
func DoIndex(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Printf("username:%s,pasword:%s", username, password)
	ctx.String(http.StatusOK, "submit success!")
}

func BindIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "bind_index.html", nil)

}

func DoBindIndex(ctx *gin.Context) {
	var loginInfo LoginInfo
	_ = ctx.ShouldBind(&loginInfo)
	fmt.Printf("UserName:%s,PassWord:%s", loginInfo.UserName, loginInfo.PassWord)
	ctx.String(http.StatusOK, "submit success!")
}

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	router.GET("/index", Index)
	router.POST("/do_index", DoIndex)

	// 数据绑定
	router.GET("/bind_index", BindIndex)
	router.POST("/do_bind_index", DoBindIndex)

	router.Run(":8080")
}
