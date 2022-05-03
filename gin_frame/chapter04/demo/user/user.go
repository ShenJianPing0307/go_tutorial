package user

import "github.com/gin-gonic/gin"

func Routers(user_router *gin.RouterGroup) {

	{
		user_router.GET("/user", User)
		user_router.POST("/do_user", DoUser)
	}

}

func DoUser(context *gin.Context) {

}

func User(context *gin.Context) {

}
