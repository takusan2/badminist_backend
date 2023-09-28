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

func (u *UserConfirmPass) ReissueConfirmPass() error {
	randomString := random.String(6)
	u.value = randomString
	return nil
}

func UserConfirmPassFromStr(str string) (UserConfirmPass, error) {
	if len(str) != 6 {
		return UserConfirmPass{}, errors.New("確認コードは6文字です")
	}
	return UserConfirmPass{str}, nil
}

func (u UserConfirmPass) Value() string {
	return u.value
}

func (u UserConfirmPass) CompareConfirmPass(confirmPass UserConfirmPass) bool {
	return u.value == confirmPass.value
}
