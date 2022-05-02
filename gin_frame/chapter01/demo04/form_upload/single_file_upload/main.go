package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUpload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func PostUpload(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	file_path := "./upload/" + file.Filename // 设置文件路径
	fmt.Println(file_path)
	_ = ctx.SaveUploadedFile(file, file_path) // 保存文件
	ctx.String(http.StatusOK, "上传成功")

}

func main() {
	router := gin.Default()
	// 加载模板文件
	router.LoadHTMLGlob("template/*")

	router.GET("/get_upload", GetUpload)
	router.POST("/post_upload", PostUpload)

	router.Run(":8080")
}
