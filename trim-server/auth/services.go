package auth

import (
	"fmt"
)

var fakeUser = map[string]string{"email": "josephtsegen10@gmail.com", "password": "1234"}

type MyError struct {
	reason string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v", e.reason)
}

func LoginService(loginCred *LoginCredentials) bool {

	fmt.Println(loginCred)

	valid := false

	if loginCred.Email == fakeUser["email"] && loginCred.Password == fakeUser["password"] {
		valid = true
	}

	return valid

}
