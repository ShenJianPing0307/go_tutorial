package routers

import (
	"github.com/gin-gonic/gin"
	"go_tutorial/gin_frame/chapter04/demo/user"
)

func Routers(router *gin.Engine) {
	user_router := router.Group("/v1")
	user.Routers(user_router)
}
