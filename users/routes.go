package users

import (
	"example/trim-server/shared/hasher"
	"example/trim-server/shared/middlewares"
	"example/trim-server/shared/orm"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {

	userController := UserController{
		&UserService{
			&UserRepository{orm.Driver},
			&hasher.HashService{}}}

	// router.GET("all", userController.GetAllUsers)
	r.GET("/", middlewares.Authenticate(), userController.GetAuthenticatedUser)
	// router.POST("create", userController.Create)
}
