package auth

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {
	router.POST("/auth/login", Login)
}
