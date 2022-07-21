package auth

import (
	"example/trim-server/shared/tokenizer"
	"example/trim-server/users"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService IAuthService
}

func (authController *AuthController) Login(context *gin.Context) {

	var loginDto LoginDto

	if err := context.ShouldBind(&loginDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, userPayload, err := authController.authService.ValidateUser(UserQueryParam{
		Email:    loginDto.UsernameOrEmail,
		Username: loginDto.UsernameOrEmail,
	}, loginDto.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		log.Fatal(err.Error())
		return
	}

	if !isValid {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid credentials"})
		return

	}

	signedToken, err := authController.authService.AssignToken(tokenizer.JWTPayload{"user_id": userPayload.UserId})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		log.Fatal(err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful!", "token": signedToken})

}

func (authController *AuthController) SignUP(context *gin.Context) {
	var userDto users.CreateUserDto

	if err := context.ShouldBindJSON(&userDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAlreadyExists, user, err := authController.authService.Register(userDto)

	if userAlreadyExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user account already exists"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Sign up successful",
		"user": gin.H{
			"user_id":       user.ID,
			"username":      user.Username,
			"email":         user.Email,
			"trimmed_links": user.TrimmedLinks}})
}
