package users

import (
	"example/trim-server/shared/hasher"
	"example/trim-server/shared/middlewares"
	"example/trim-server/shared/orm"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {

	userController := UserController{
		&UserService{
			&UserRepository{orm.Driver},
			&hasher.HashService{}}}

	// router.GET("all", userController.GetAllUsers)
	router.GET("/", middlewares.Authentize(), userController.GetAuthenticatedUser)
	// router.POST("create", userController.Create)
}
