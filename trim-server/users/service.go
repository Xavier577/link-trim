package users

import (
	"example/trim-server/shared/hasher"
	"fmt"
)

type UserService struct {
	userRepository IUserRepository
	hashService    hasher.IHashService
}

type IUserService interface {
	FindUser(id uint) (bool, User, error)
	FindMany() (bool, []User, error)
	CreateUser(userDto *CreateUserDto) (bool, User, error)
}

func (US *UserService) FindUser(id uint) (bool, User, error) {
	return US.userRepository.FindUserById(id)
}

func (US *UserService) FindMany() (bool, []User, error) {
	return US.userRepository.FindMany()
}

func (US *UserService) CreateUser(userDto *CreateUserDto) (bool, User, error) {
	passHash, err := US.hashService.Hash(userDto.Password)

	if err != nil {
		fmt.Println("hash error")
		panic(err.Error())
	}

	userDto.Password = passHash

	return US.userRepository.CreateUser(userDto)

}
