package main

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	UserName string `form:"username" valid:"Required"`
	PassWord string `form:"password" valid:"Required"`
}

func ValidIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "valid_index.html", nil)

}

func DoValidIndex(ctx *gin.Context) {

	// 1、声明结构体变量
	var loginInfo LoginInfo
	// 2、绑定结构体变量
	//_ = ctx.ShouldBind(&loginInfo)
	if err := ctx.ShouldBind(&loginInfo); err != nil {
		//handle error
		return
	}
	// 3、初始化beego验证器
	valid := validation.Validation{}
	// 定制错误信息
	var messages = map[string]string{
		"Required": "不能为空",
		"MinSize":  "最短长度为 %d",
		"Length":   "长度必须为 %d",
		"Numeric":  "必须是有效的数字",
		"Email":    "必须是有效的电子邮件地址",
		"Mobile":   "必须是有效的手机号码",
	}
	validation.SetDefaultMessage(messages)

	// 4、开始验证
	b, _ := valid.Valid(&loginInfo)
	if !b {
		// 5、输出未验证通过的错误信息
		// UserName.Required.Can not be empty
		// PassWord.Required.Can not be empty
		for _, err1 := range valid.Errors {
			fmt.Println(err1.Key)
			fmt.Println(err1.Message)
		}
	}
	fmt.Printf("UserName:%s,PassWord:%s", loginInfo.UserName, loginInfo.PassWord)
	ctx.String(http.StatusOK, "submit success!")
}

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	// 数据校验
	router.GET("/valid_index", ValidIndex)
	router.POST("/do_valid_index", DoValidIndex)

	router.Run(":8080")

}
