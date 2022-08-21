package auth

import (
	"example/trim-server/global"
	"example/trim-server/shared/hasher"
	"example/trim-server/shared/tokenizer"
	"example/trim-server/users"
	"time"
)

type AuthService struct {
	userRepository users.IUserRepository
	hashService    hasher.IHashService
}

type IAuthService interface {
	ValidateUser(userQueryParam UserQueryParam, pass string) (bool, UserPayload, error)

	Register(createUserDto users.CreateUserDto) (bool, users.User, error)

	AssignToken(payload tokenizer.JWTPayload) (string, error)
}

type UserQueryParam struct {
	Email    string
	Username string
}

type UserPayload struct {
	UserId string
}

func (authService *AuthService) Register(createUserDto users.CreateUserDto) (bool, users.User, error) {
	return authService.userRepository.CreateUser(&createUserDto)
}

func (authService *AuthService) ValidateUser(userQueryParam UserQueryParam, pass string) (bool, UserPayload, error) {
	var isValid bool
	var internalError error
	var userPayload UserPayload

	isNotFoundErr, user, err := authService.userRepository.FindUserByEmailOrUsername(userQueryParam.Email, userQueryParam.Username)

	if isNotFoundErr {
		isValid = false
	}

	if err != nil && !isNotFoundErr {
		internalError = err
	}

	if passMatch := authService.hashService.Compare(pass, user.Password); passMatch {
		isValid = true
		userPayload.UserId = user.UserId
	}

	return isValid, userPayload, internalError
}

func (authSerice *AuthService) AssignToken(payload tokenizer.JWTPayload) (string, error) {
	return tokenizer.GenerateHS256Token(tokenizer.TokenGenOptions{
		Payload:    payload,
		Secret:     global.Env("JWT_SECRET"),
		ExpiryDate: time.Now().Add(time.Duration(time.Second) * global.TOKEN_HOUR_LIFESPAN),
	})

}
