package trimmedlinks

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrimmedLinkController struct {
	trimmedLinkService ITrimmedLinkService
}

// @Summary redirect to the original link from the uuid
// @ID redirect-to-original-link
// @Produce json
// @Param link_uuid path string true "short uuid"
// @Success 300
// @Failure 404
// @Router /:link_uuid [get]
func (TLC *TrimmedLinkController) RedirectToOrignalLink(context *gin.Context) {

	linkUUID := context.Param("link_uuid")

	if linkUUID == "" {

		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid url"})

		return
	}

	isNotFoundErr, trimmedLink, err := TLC.trimmedLinkService.FetchOriginalLink(linkUUID)

	if isNotFoundErr {
		context.JSON(http.StatusNotFound, "not Found")

		return
	}

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		log.Fatal(err.Error())

		return
	}

	context.Redirect(http.StatusMovedPermanently, trimmedLink.Link)

}

func (TLC *TrimmedLinkController) CreateTrimmedLink(context *gin.Context) {

	var user_id, _ = context.Get("user_id")

	userId := uint(user_id.(float64))

	var createTrimmedLinkDto *CreateTrimmedLinkDto

	if err := context.ShouldBind(&createTrimmedLinkDto); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createTrimmedLinkDto.UserId = userId

	trimmedLink, err := TLC.trimmedLinkService.CreateTrimmedLink(createTrimmedLinkDto)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	context.JSON(http.StatusOK, trimmedLink)
}
