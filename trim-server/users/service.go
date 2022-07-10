package users

import "example/trim-server/hasher"

type UserServiceStruct struct{}

func (userService *UserServiceStruct) FindUser(id string) string {
	return id
}

func (userService *UserServiceStruct) CreateUser(userDto *CreateUserDto) bool {

	passHash, err := hasher.Hash(userDto.Password)

	if err != nil {
		panic(err)
	}

	userDto.Password = passHash

	isCreated := userRepository.CreateUser(userDto)

	return isCreated
}

var userService = new(UserServiceStruct)
