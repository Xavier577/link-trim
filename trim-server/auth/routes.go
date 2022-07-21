package auth

import (
	"example/trim-server/shared/hasher"
	"example/trim-server/shared/orm"
	"example/trim-server/users"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {

	authController := AuthController{&AuthService{
		&users.UserRepository{DbClient: orm.Driver},
		&hasher.HashService{}}}

	router.POST("login", authController.Login)
	router.POST("sign-up", authController.SignUP)

}
