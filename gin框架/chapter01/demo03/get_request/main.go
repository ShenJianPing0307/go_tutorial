package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	ctx.String(http.StatusOK, "success!")
}

func Index1(ctx *gin.Context) {
	id := ctx.Query("id")
	fmt.Println(id) // 12
	ctx.String(http.StatusOK, "success!")
}

func Index2(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "0")
	fmt.Println(id) // 0
	ctx.String(http.StatusOK, "success!")
}

func Index3(ctx *gin.Context) {
	idList := ctx.QueryArray("id")
	fmt.Println(idList) // [1,2,3,4,5]
	ctx.String(http.StatusOK, "success!")
}

func Index4(ctx *gin.Context) {
	user := ctx.QueryMap("user")
	fmt.Println(user) // map[age:15 name:"lily"]
	ctx.String(http.StatusOK, "success!")
}

func main() {
	router := gin.Default()

	// 路径参数获取，如：http://127.0.0.1:8080/index/12，获取12
	router.GET("/index/:id", Index)

	// 查询参数获取，如：http://127.0.0.1:8080/index?id=12，获取12
	router.GET("/index1", Index1)
	router.GET("/index2", Index2)
	router.GET("/index3", Index3)
	router.GET("/index4", Index4)

	router.Run(":8080")
}
