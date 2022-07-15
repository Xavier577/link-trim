package users

import (
	"errors"
	"example/trim-server/database"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct{}

var db *gorm.DB

func InitializeUserRepository(db_driver *gorm.DB) {
	db = db_driver
	error := db_driver.AutoMigrate(&database.User{})

	if error != nil {
		panic(error)
	}
}

func (userRepo *UserRepository) FindUser(id uint) (bool, database.User, error) {
	var user database.User
	result := db.Take(&user, id)
	isNotFoundErr := errors.Is(result.Error, gorm.ErrRecordNotFound)
	err := result.Error
	return isNotFoundErr, user, err
}

func (userRepo *UserRepository) FindMany() (bool, []database.User, error) {
	var users []database.User
	result := db.Find(&users)
	isNotFoundErr := errors.Is(result.Error, gorm.ErrRecordNotFound)
	err := result.Error
	return isNotFoundErr, users, err
}

func (userRepo *UserRepository) CreateUser(userDto *CreateUserDto) (bool, database.User, error) {
	user := database.User{Username: userDto.Username,
		Email:    userDto.Email,
		Password: userDto.Password}
	result := db.Or(database.User{Username: userDto.Username,
		Email: userDto.Email}).FirstOrCreate(&user)
	userAlreadyExists := user.ID == 0

	fmt.Println(userAlreadyExists)
	return userAlreadyExists, user, result.Error
}
