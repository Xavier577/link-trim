package auth

import (
	"example/trim-server/shared/hasher"
	"example/trim-server/shared/orm"
	"example/trim-server/users"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {

	authController := AuthController{&AuthService{
		&users.UserRepository{DbClient: orm.Driver},
		&hasher.HashService{}}}

	r.POST("login", authController.Login)
	r.POST("sign-up", authController.SignUP)

}
