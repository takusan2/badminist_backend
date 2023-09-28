package user

import (
	"errors"
	"regexp"
)

type UserEmail struct {
	value string
}

func NewUserEmail(email string) (UserEmail, error) {
	// emailの形式かどうかのバリデーション
	// 正規表現でバリデーションを行う

	email_regexp := regexp.MustCompile(`^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`)
	if !email_regexp.MatchString(email) {
		return UserEmail{}, errors.New("正しいメールアドレスを入力してください")
	}

	return UserEmail{email}, nil
}

func (u UserEmail) Value() string {
	return u.value
}
