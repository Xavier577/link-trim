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
// @Tags Index
// @ID redirect-to-original-link
// @Produce json
// @Param link_uuid path string true "short uuid"
// @Success 300
// @Failure 404
// @Router /:link_uuid [get]
func (TLC *TrimmedLinkController) RedirectToOrignalLink(context *gin.Context) {

	linkUUID := context.Param("link_uuid")

	if linkUUID == "" {

		context.JSON(http.StatusBadRequest, gin.H{"message": "INVALID_URL"})

		return
	}

	isNotFoundErr, trimmedLink, err := TLC.trimmedLinkService.FetchOriginalLink(linkUUID)

	if isNotFoundErr {
		context.JSON(http.StatusNotFound, "NOT_FOUND")

		return
	}

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "INTERNAL_SERVER_ERROR"})
		log.Fatal(err.Error())

		return
	}

	context.Redirect(http.StatusMovedPermanently, trimmedLink.Link)

}

// @Summary endpoints create identifier for a link on server
// @Tags Links
// @ID Create_Link
// @Produce json
// @Param data body CreateTrimmedLinkDto true "url info"
// @Success 200 {string} link_id
// @Failure 400
// @Router /link [post]
func (TLC *TrimmedLinkController) CreateTrimmedLink(context *gin.Context) {

	var userId, _ = context.Get("user_id")

	var createTrimmedLinkDto *CreateTrimmedLinkDto

	if err := context.ShouldBind(&createTrimmedLinkDto); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createTrimmedLinkDto.UserId = userId.(string)

	trimmedLink, err := TLC.trimmedLinkService.CreateTrimmedLink(createTrimmedLinkDto)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "INTERNAL_SERVER_ERROR"})
		log.Fatal(err.Error())
		return
	}

	context.JSON(http.StatusOK, trimmedLink)
}
