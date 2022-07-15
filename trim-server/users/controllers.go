package users

import (
	"example/trim-server/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userService = new(UserService)

type UserController struct{}

func (userController *UserController) GetUser(context *gin.Context) {

	userId, paramTypeErr := global.StringToUInt(context.Param("id"))

	isNotFoundErr, user, err := userService.FindUser(userId)

	if isNotFoundErr || paramTypeErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid request param"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user_id": user.ID, "username": user.Username, "email": user.Email})

}

func (userController *UserController) GetAllUsers(context *gin.Context) {
	isNotFoundErr, users, err := userService.FindMany()

	if isNotFoundErr {
		context.JSON(http.StatusBadRequest, gin.H{"message": "no users in the database"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": users})

}

func (userController *UserController) Create(context *gin.Context) {
	var userDto CreateUserDto

	if err := context.ShouldBindJSON(&userDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAlreadyExists, user, err := userService.CreateUser(&userDto)

	if userAlreadyExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "user successfully created",
		"user": gin.H{
			"user_id":  user.ID,
			"username": user.Username,
			"email":    user.Email}})

}
