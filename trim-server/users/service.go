package users

import (
	"example/trim-server/database"
	"example/trim-server/hasher"
	"fmt"
)

type UserService struct{}

var userRepository = new(UserRepository)

func (userService *UserService) FindUser(id uint) (bool, database.User, error) {
	return userRepository.FindUserById(id)
}

func (UserService *UserService) FindMany() (bool, []database.User, error) {
	return userRepository.FindMany()
}

func (userService *UserService) CreateUser(userDto *CreateUserDto) (bool, database.User, error) {
	passHash, err := hasher.Hash(userDto.Password)

	if err != nil {
		fmt.Println("hash error")
		panic(err.Error())
	}

	userDto.Password = passHash

	return userRepository.CreateUser(userDto)

}
