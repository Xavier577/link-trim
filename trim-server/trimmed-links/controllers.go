package trimmedlinks

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTrimmedLink(context *gin.Context) {
	link_id, err := strconv.Atoi(context.Params.ByName("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid params"})
		return
	}

	affected, err := getLinkById(uint(link_id), new(LinkRepositoryStruct))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		log.Fatal(err)
	}

	context.String(http.StatusOK, fmt.Sprint(affected))
}
