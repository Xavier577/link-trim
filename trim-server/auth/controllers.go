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

// @Summary this endpoint enable the user to login
// @Tags Authentication
// @ID login
// @Produce json
// @Param data body LoginDto true "Login Credentials"
// @Success 200 {object} LoginSuccessResponse
// @Failure 400 {object} AuthFailureResponse
// @Router /auth/login [post]
func (authController *AuthController) Login(context *gin.Context) {

	var loginDto LoginDto

	if err := context.ShouldBind(&loginDto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(loginDto)

	isValid, userPayload, err := authController.authService.ValidateUser(UserQueryParam{
		Email:    loginDto.Identifier,
		Username: loginDto.Identifier,
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
		context.JSON(http.StatusInternalServerError, gin.H{"message": "INTERNAL_SERVER_ERROR", "status": "faliure"})
		log.Fatal(err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "login successful!", "token": signedToken, "status": "success"})

}

// @Summary endpoint enable users to register
// @Tags Authentication
// @ID Register
// @Produce json
// @Param data body users.CreateUserDto true "user information"
// @Success 200 {object} SignUpSuccess
// @Failure 400 {object} AuthFailureResponse
// @Router /auth/sign-up [post]
func (authController *AuthController) SignUP(context *gin.Context) {
	var userDto users.CreateUserDto

	if err := context.ShouldBindJSON(&userDto); err != nil {

		if err.Error() == "EOF" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		log.Println(err.Error())

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
			"user_id":       user.UserID,
			"username":      user.Username,
			"email":         user.Email,
			"trimmed_links": user.TrimmedLinks}})
}
