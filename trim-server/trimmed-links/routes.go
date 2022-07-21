package trimmedlinks

import (
	"example/trim-server/shared/middlewares"
	"example/trim-server/shared/orm"
	"example/trim-server/shared/uuid"

	"github.com/gin-gonic/gin"
)

func Routes(apiRouter *gin.RouterGroup, indexRouter *gin.Engine) {

	trimmedLinkController := TrimmedLinkController{
		&TrimmedLinkService{
			&TrimmedLinkRepository{orm.Driver},
			&uuid.UUIDService{},
		}}

	indexRouter.GET(":link_uuid", trimmedLinkController.RedirectToOrignalLink)
	apiRouter.POST("create", middlewares.Authentize(), trimmedLinkController.CreateTrimmedLink)

}
