package trimmedlinks

import (
	"example/trim-server/shared/middlewares"
	"example/trim-server/shared/orm"
	"example/trim-server/shared/uuid"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {

	trimmedLinkController := TrimmedLinkController{
		&TrimmedLinkService{
			&TrimmedLinkRepository{orm.Driver},
			&uuid.UUIDService{},
		}}

	r.GET(":link_uuid", trimmedLinkController.RedirectToOrignalLink)
	r.POST("create", middlewares.Authenticate(), trimmedLinkController.CreateTrimmedLink)

}
