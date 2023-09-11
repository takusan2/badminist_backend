package user

import (
	"errors"

	"github.com/labstack/gommon/random"
)

type UserConfirmPass struct {
	value string
}

func NewUserConfirmPass() UserConfirmPass {
	randomString := random.String(6)
	return UserConfirmPass{randomString}
}

func UserConfirmPassFromStr(str string) (UserConfirmPass, error) {
	if len(str) != 6 {
		return UserConfirmPass{}, errors.New("length is not 6")
	}
	return UserConfirmPass{str}, nil
}

func (u UserConfirmPass) Value() string {
	return u.value
}

func (u UserConfirmPass) CompareConfirmPass(confirmPass UserConfirmPass) bool {
	return u.value == confirmPass.value
}
