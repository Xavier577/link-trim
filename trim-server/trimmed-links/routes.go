package trimmedlinks

import "github.com/gin-gonic/gin"

var userController = new(UserController)

func Routes(router *gin.RouterGroup) {

	router.GET("/trimmed-link/:id", userController.GetTrimmedLink)
	router.POST("/trimmed-link/create", userController.CreateTrimmedLink)

}
