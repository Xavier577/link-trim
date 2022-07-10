package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {

	var loginCred LoginCredentials

	if err := context.ShouldBindJSON(&loginCred); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := LoginService(&loginCred)

	if valid == false {
		context.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"ok": valid})

}
