package trimmedlinks

import "github.com/gin-gonic/gin"

func Routes(router *gin.RouterGroup) {

	router.GET("/trimmed-link", GetTrimmedLink)
}
