package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetUpload(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func PostUpload(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		unix_time := time.Now().Unix() // int类型的时间戳
		unix_time_str := strconv.FormatInt(unix_time, 10)
		file_path := "./upload/" + unix_time_str + file.Filename
		ctx.SaveUploadedFile(file, file_path)
	}
	ctx.String(http.StatusOK, "上传成功")

}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/*")

	router.GET("/get_upload", GetUpload)
	router.POST("/post_upload", PostUpload)

	router.Run(":8080")
}
