package main

import (
	"github.com/gin-gonic/gin"
	"go_tutorial/gin_frame/chapter04/demo/routers"
)

func main() {
	router := gin.Default()
	routers.Routers(router)

	router.Run(":8080")
}
