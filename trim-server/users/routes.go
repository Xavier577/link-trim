package users

import (
	"github.com/gin-gonic/gin"
)

var userController = new(UserController)

func Routes(router *gin.RouterGroup) {
	router.GET("user/all", userController.GetAllUsers)
	router.GET("/user/:id", userController.GetUser)
	router.POST("/user/create", userController.Create)
}
