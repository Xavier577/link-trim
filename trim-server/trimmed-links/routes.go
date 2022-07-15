package trimmedlinks

import "github.com/gin-gonic/gin"

var trimmedLinkController = new(TrimmedLinkController)

func Routes(router *gin.RouterGroup) {
	router.GET(":id", trimmedLinkController.GetTrimmedLink)
	router.POST("create", trimmedLinkController.CreateTrimmedLink)

}
