package user

import (
	"errors"
	"regexp"
)

type UserName struct {
	value string
}

func NewUserName(name string) (UserName, error) {
	if name == "" {
		return UserName{}, errors.New("空白のユーザー名は登録できません")
	}
	if len(name) > 255 {
		return UserName{}, errors.New("ユーザー名は255文字以内で入力してください")
	}
	// 半角空白または全角空白のみの場合はエラー
	regexp_pattern := regexp.MustCompile(`^[ 　]+$`)
	if regexp_pattern.MatchString(name) {
		return UserName{}, errors.New("空白のみのユーザー名は登録できません")
	}

	return UserName{name}, nil
}

func (u *UserName) Value() string {
	return u.value
}
