package trimmedlinks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var trimmedLinkService = new(TrimmedLinkService)

func (userController *UserController) GetTrimmedLink(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "comming soon"})
}

func (UserController *UserController) CreateTrimmedLink(context *gin.Context) {

	var createTrimmedLinkDto *CreateTrimmedLinkDto

	if err := context.ShouldBind(&createTrimmedLinkDto); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	trimmedLink, err := trimmedLinkService.CreateTrimmedLink(createTrimmedLinkDto)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error!"})
		return
	}

	context.JSON(http.StatusOK, trimmedLink)
}
