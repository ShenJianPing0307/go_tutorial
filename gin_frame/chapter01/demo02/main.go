package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func StructDemo(ctx *gin.Context) {
	user := User{
		Id:   1,
		Name: "lily",
		Age:  25,
	}
	ctx.HTML(http.StatusOK, "index.html", user)
}

func ArrayDemo(ctx *gin.Context) {
	numArr := [5]int{1, 2, 3, 4, 5}
	ctx.HTML(http.StatusOK, "index1.html", numArr)
}

func ArrStruDemo(ctx *gin.Context) {
	arr_struct := [2]User{
		{
			Id:   1,
			Name: "lily",
			Age:  25,
		},
		{
			Id:   2,
			Name: "berry",
			Age:  15,
		},
	}
	ctx.HTML(http.StatusOK, "index2.html", arr_struct)
}

func MapDemo(ctx *gin.Context) {
	data := map[string]string{
		"Name": "zhangsan",
		"Age":  "15",
	}
	ctx.HTML(http.StatusOK, "index3.html", data)
}

func MapStruDemo(ctx *gin.Context) {
	data := map[string]User{
		"u1": {
			Id:   1,
			Name: "lily",
			Age:  25,
		},
		"u2": {
			Id:   2,
			Name: "berry",
			Age:  15,
		},
	}
	ctx.HTML(http.StatusOK, "index4.html", data)
}

func SliceDemo(ctx *gin.Context) {
	numSlice := []int{1, 2} // 无需指定数组的长度，动态数组
	ctx.HTML(http.StatusOK, "index5.html", numSlice)
}

func main() {
	router := gin.Default()
	// 加载模板路径
	router.LoadHTMLGlob("template/*")

	// 结构体渲染
	router.GET("/index", StructDemo)
	// 数组渲染
	router.GET("/index1", ArrayDemo)
	// 数组结构体
	router.GET("/index2", ArrStruDemo)
	// map
	router.GET("/index3", MapDemo)
	// map结构体
	router.GET("/index4", MapStruDemo)
	// 切片
	router.GET("/index5", SliceDemo)

	router.Run(":8080")
}
