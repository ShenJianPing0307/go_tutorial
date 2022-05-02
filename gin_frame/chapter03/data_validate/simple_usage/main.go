package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func DoIndex(ctx *gin.Context) {
	var loginInfo LoginInfo
	err := ctx.ShouldBind(&loginInfo)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(400, gin.H{
			"msg":  "fail",
			"code": 400,
		})
	}

	ctx.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
	})
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	router.GET("/index", Index)
	router.POST("/do_index", DoIndex)

	router.Run(":8080")
}
