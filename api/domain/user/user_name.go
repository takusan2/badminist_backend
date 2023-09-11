package user

import (
	"fmt"
	"regexp"
)

type UserName struct {
	value string
}

func NewUserName(name string) (UserName, error) {
	if name == "" {
		return UserName{}, fmt.Errorf("user name is empty")
	}
	if len(name) > 255 {
		return UserName{}, fmt.Errorf("user name is too long")
	}
	// 半角空白または全角空白のみの場合はエラー
	regexp_pattern := regexp.MustCompile(`^[ 　]+$`)
	if regexp_pattern.MatchString(name) {
		return UserName{}, fmt.Errorf("user name is invalid")
	}

	return UserName{name}, nil
}

func (u *UserName) Value() string {
	return u.value
}
