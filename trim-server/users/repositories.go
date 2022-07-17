package users

import (
	"errors"
	"example/trim-server/database"

	"gorm.io/gorm"
)

type UserRepository struct{}

var db *gorm.DB

func InitializeUserRepository(db_driver *gorm.DB) {
	db = db_driver
	err := db_driver.AutoMigrate(&database.User{})

	if err != nil {
		panic(err)
	}
}

func (userRepo *UserRepository) FindUserById(id uint) (bool, database.User, error) {
	var user database.User
	err := db.Model(&database.User{}).Preload("TrimmedLinks").Take(&user, id).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, user, err
}

func (userRepo *UserRepository) FindMany() (bool, []database.User, error) {
	var users []database.User
	err := db.Model(&database.User{}).Preload("TrimmedLinks").Find(&users).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, users, err
}

func (userRepo *UserRepository) CreateUser(userDto *CreateUserDto) (bool, database.User, error) {

	user := database.User{Username: userDto.Username,
		Email:    userDto.Email,
		Password: userDto.Password}

	userAlreadyExists := userRepo.UserExistsWithUsernameOrEmail(userDto.Username, user.Email)

	if !userAlreadyExists {
		result := db.Create(&user)

		return userAlreadyExists, user, result.Error
	}

	return userAlreadyExists, user, nil

}

func (userRepo *UserRepository) UserExistsWithUsernameOrEmail(username, email string) bool {
	var exists bool
	db.Model(&database.User{}).Select("count(*) > 0").Where("username = ? or email = ?", username, email).Find(&exists)
	return exists
}
