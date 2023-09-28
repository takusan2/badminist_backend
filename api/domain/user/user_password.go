package user

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UserPassword struct {
	value string
}

func NewUserPassword(password string) (UserPassword, error) {
	if len(password) < 6 {
		return UserPassword{}, errors.New("パスワードは6文字以上で入力してください")
	}
	if len(password) > 255 {
		return UserPassword{}, errors.New("パスワードは255文字以内で入力してください")
	}
	// 正規表現でバリデーションを行う
	password_regexp := regexp.MustCompile(`^(?i:[^ @"<>]+|".*")`)
	if !password_regexp.MatchString(password) {
		return UserPassword{}, errors.New("パスワードに「スペース, @, \", <>」は使用できません")
	}
	return UserPassword{password}, nil
}

func NewUserPasswordFromHash(password string) (UserPassword, error) {
	return UserPassword{password}, nil
}

func (u *UserPassword) Value() string {
	return u.value
}

// 暗号化
func (u *UserPassword) Encrypt() error {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(u.value), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.value = string(bcryptPassword)
	return nil
}

// 暗号化されたパスワードとの比較
func (u *UserPassword) Authenticate(password UserPassword) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.value), []byte(password.value)); err != nil {
		return false
	}
	return true
}
