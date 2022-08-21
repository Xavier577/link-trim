package users

import (
	"example/trim-server/shared/hasher"
)

type UserService struct {
	userRepository IUserRepository
	hashService    hasher.IHashService
}

type IUserService interface {
	FindUser(userId string) (bool, User, error)

	FindMany() (bool, []User, error)

	CreateUser(userDto *CreateUserDto) (bool, User, error)
}

func (US *UserService) FindUser(userId string) (bool, User, error) {
	return US.userRepository.FindUserById(userId)
}

func (US *UserService) FindMany() (bool, []User, error) {
	return US.userRepository.FindMany()
}

func (US *UserService) CreateUser(userDto *CreateUserDto) (bool, User, error) {
	passHash, err := US.hashService.Hash(userDto.Password)

	if err != nil {
		var user User
		return false, user, err

	}

	userDto.Password = passHash

	return US.userRepository.CreateUser(userDto)

}
