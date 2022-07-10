package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerStruct struct{}

func (userController *UserControllerStruct) GetUser(context *gin.Context) {
	user := userService.FindUser("13")
	context.String(http.StatusOK, user)
}

func (UserControllerStruct *UserControllerStruct) Create(context *gin.Context) {
	var userDto CreateUserDto

	if err := context.ShouldBindJSON(&userDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isCreated := userService.CreateUser(&userDto)

	if isCreated {
		context.JSON(http.StatusOK, gin.H{"message": "user successfully created"})
	} else {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
}

var userController = new(UserControllerStruct)
