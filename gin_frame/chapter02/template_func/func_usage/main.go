package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func PrintFunc(ctx *gin.Context) {
	name := "alice"
	ctx.HTML(http.StatusOK, "print_func.html", name)
}

func BoolFunc(ctx *gin.Context) {

	data := map[string]interface{}{
		"arr": [3]int{1, 2, 3},
		"a":   "hello",
	}

	ctx.HTML(http.StatusOK, "and_or_not_func.html", data)

}

func IndexFunc(ctx *gin.Context) {

	data := map[string]interface{}{
		"arr": [3]int{1, 2, 3},
		"map_data": map[string]string{
			"name": "alice",
			"addr": "beijing",
		},
	}

	ctx.HTML(http.StatusOK, "index_len_func.html", data)
}

func CompareFunc(ctx *gin.Context) {
	age := 20
	ctx.HTML(http.StatusOK, "compare_func.html", age)
}

func FormatFunc(ctx *gin.Context) {
	now_time := time.Now().Format("2006-01-02 15:04:05")
	ctx.HTML(http.StatusOK, "time_format_func.html", now_time)
}

func DefineFunc(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "define_func.html", nil)
}

func add(a int, b int) int {
	return a + b
}

func main() {

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"Add": add, // 字符串名称前端使用
	})

	router.LoadHTMLGlob("template/*")

	router.GET("/print_func", PrintFunc)
	router.GET("/and_or_not_func", BoolFunc)
	router.GET("/index_len_func", IndexFunc)
	router.GET("/compare_func", CompareFunc)
	router.GET("/time_format_func", FormatFunc)
	router.GET("/define_func", DefineFunc)

	router.Run(":8080")

}
