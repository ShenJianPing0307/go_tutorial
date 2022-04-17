package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func Index(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	fmt.Println(err)
	fmt.Println(user.Username, user.Password) // "llkk" "123"

	ctx.String(http.StatusOK, "success")

}

func main() {
	router := gin.Default()

	router.GET("/index", Index)

	router.Run(":8080")
}
