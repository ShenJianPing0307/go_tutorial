package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"go_tutorial/gin_frame/chapter03/data_validate/custom_validator/handler"
	"net/http"
)

type LoginInfo struct {
	UserName string `form:"username" json:"username" binding:"len_valid"`
	PassWord string `form:"password" json:"password" binding:"len_valid"`
}

func ValidIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "valid_index.html", nil)

}

func DoValidIndex(ctx *gin.Context) {
	var loginInfo LoginInfo
	_ = ctx.ShouldBind(&loginInfo)
	fmt.Printf("UserName:%s,PassWord:%s", loginInfo.UserName, loginInfo.PassWord)
	ctx.String(http.StatusOK, "submit success!")
}

func main() {

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("len_valid", handler.CustomValid)
	}

	router.LoadHTMLGlob("template/*")

	// 数据校验
	router.GET("/valid_index", ValidIndex)
	router.POST("/do_valid_index", DoValidIndex)

	router.Run(":8080")

}
