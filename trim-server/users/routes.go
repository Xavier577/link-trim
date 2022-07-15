package users

import (
	"github.com/gin-gonic/gin"
)

var userController = new(UserController)

func Routes(router *gin.RouterGroup) {
	router.GET("all", userController.GetAllUsers)
	router.GET(":id", userController.GetUser)
	router.POST("create", userController.Create)
}
