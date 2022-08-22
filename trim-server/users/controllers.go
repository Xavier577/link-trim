package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService IUserService
}

// @Tags User
// @ID GetAuthenticated User
// @Produce json
// @Security Authorization
// @Success 200
// @Failure 401
// @Router /user [get]
func (userController *UserController) GetAuthenticatedUser(context *gin.Context) {

	var userId, _ = context.Get("user_id")

	_, user, err := userController.userService.FindUser(userId.(string))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"user_id":       user.UserID,
		"username":      user.Username,
		"email":         user.Email,
		"trimmed_links": user.TrimmedLinks})

}

/*

// @Summary get all the users in the database
// @Produce json
// @Success 200 {array} models.User
// @Router /api/v1/user/all [get]
// func (userController *UserController) GetAllUsers(context *gin.Context) {
// 	isNotFoundErr, users, err := userController.userService.FindMany()

// 	if isNotFoundErr {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "no users in the database"})
// 		return
// 	}

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
// 		log.Println(err.Error())
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"data": users})

// }

*/

func (userController *UserController) Create(context *gin.Context) {
	var userDto CreateUserDto

	if err := context.ShouldBindJSON(&userDto); err != nil {

		if err.Error() == "EOF" {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		}

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAlreadyExists, user, err := userController.userService.CreateUser(&userDto)

	if userAlreadyExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user already exists"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "INTERNAL_SERVER_ERROR"})
		log.Fatal(err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "user successfully created",
		"user": gin.H{
			"user_id":       user.UserID,
			"username":      user.Username,
			"email":         user.Email,
			"trimmed_links": user.TrimmedLinks}})

}
