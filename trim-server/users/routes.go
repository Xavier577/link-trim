package users

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/user", userController.GetUser)
	router.POST("/user/create", userController.Create)
}
