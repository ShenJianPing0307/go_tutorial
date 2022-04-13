package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	// 上下文
	map_data := map[string]interface{}{
		"arr_num": [3]int{1, 2, 3},
		"name":    "lily",
	}

	ctx.HTML(http.StatusOK, "index.html", map_data)
}

func Index1(ctx *gin.Context) {
	// 上下文
	map_data := map[string]interface{}{
		"arr_num": [3]int{1, 2, 3},
		"name":    "lily",
	}

	ctx.HTML(http.StatusOK, "index1.html", map_data)
}

func Index2(ctx *gin.Context) {
	// 上下文
	map_data := map[string]interface{}{
		"arr_num": [3]int{1, 2, 3},
		"name":    "lily",
	}

	ctx.HTML(http.StatusOK, "index2.html", map_data)
}

func main() {
	router := gin.Default()

	// 加载模板
	router.LoadHTMLGlob("template/*")

	// 模板语法上下文
	router.GET("/index", Index)
	// 模板语法if条件
	router.GET("/index1", Index1)
	// range的用法
	router.GET("/index2", Index2)

	router.Run(":8080")
}
