package users

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DbClient *gorm.DB
}

type IUserRepository interface {
	FindUserById(id string) (bool, User, error)

	FindUserByEmail(email string) (bool, User, error)

	FindUserByUsername(username string) (bool, User, error)

	FindUserByEmailOrUsername(email string, username string) (bool, User, error)

	FindMany() (bool, []User, error)

	CreateUser(userDto *CreateUserDto) (bool, User, error)
}

func (userRepo *UserRepository) FindUserById(userId string) (bool, User, error) {
	var user User
	err := userRepo.DbClient.Model(&User{}).Preload("TrimmedLinks").Take(&user, User{UserID: userId}).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, user, err
}

func (userRepo *UserRepository) FindUserByEmailOrUsername(email string, username string) (bool, User, error) {
	var user User
	err := userRepo.DbClient.Model(&User{}).Preload("TrimmedLinks").Where(User{Email: email}).
		Or(User{Username: username}).Take(&user).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, user, err
}

func (userRepo *UserRepository) FindUserByEmail(email string) (bool, User, error) {
	var user User
	err := userRepo.DbClient.Model(&User{}).Preload("TrimmedLinks").Take(&user, User{Email: email}).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, user, err
}

func (userRepo *UserRepository) FindUserByUsername(username string) (bool, User, error) {
	var user User
	err := userRepo.DbClient.Model(&User{}).Preload("TrimmedLinks").Take(&user, User{Username: username}).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, user, err
}

func (userRepo *UserRepository) FindMany() (bool, []User, error) {
	var users []User
	err := userRepo.DbClient.Model(&User{}).Preload("TrimmedLinks").Find(&users).Error
	isNotFoundErr := errors.Is(err, gorm.ErrRecordNotFound)
	return isNotFoundErr, users, err
}

func (userRepo *UserRepository) CreateUser(userDto *CreateUserDto) (bool, User, error) {

	user := User{Username: userDto.Username,
		Email:    userDto.Email,
		Password: userDto.Password}

	userAlreadyExists := userRepo.UserExistsWithUsernameOrEmail(userDto.Username, user.Email)

	if !userAlreadyExists {
		result := userRepo.DbClient.Create(&user)

		return userAlreadyExists, user, result.Error
	}

	return userAlreadyExists, user, nil

}

func (userRepo *UserRepository) UserExistsWithUsernameOrEmail(username, email string) bool {
	var exists bool
	userRepo.DbClient.Model(&User{}).Select("count(*) > 0").Where("username = ? or email = ?", username, email).Find(&exists)
	return exists
}
