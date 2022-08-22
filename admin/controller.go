package admin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService IAdminService
}

func (adminController *AdminController) GetAllAdmins(ctx *gin.Context) {
	admins, err := adminController.adminService.GetAdmins()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "INTERNAL_SERVER_ERROR"})
		log.Fatal(err.Error())
	}
	ctx.JSON(http.StatusOK, admins)
}
