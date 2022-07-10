package users

import (
	"example/trim-server/database"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitializeUserRepository(db_driver *gorm.DB) {
	db = db_driver
	error := db_driver.AutoMigrate(&database.User{})

	if error != nil {
		panic(error)
	}
}

type UserRepositoryStruct struct{}

type IUserRepository interface {
	FindUser(id string) *gorm.DB
	CreateUser() bool
}

func (userRepo *UserRepositoryStruct) FindUser(id uint) *gorm.DB {
	result := db.First(&database.User{ID: id})
	return result
}

func (userRepo *UserRepositoryStruct) CreateUser(userDto *CreateUserDto) bool {
	result := db.Model(&database.User{}).Create(map[string]interface{}{
		"Username": userDto.Username,
		"Email":    userDto.Email,
		"Password": userDto.Password,
	})

	return result.Error == nil
}

var userRepository = new(UserRepositoryStruct)
