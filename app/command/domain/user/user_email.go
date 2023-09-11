package user

import (
	"fmt"
	"regexp"
)

type UserEmail struct {
	value string
}

func NewUserEmail(email string) (UserEmail, error) {
	// emailの形式かどうかのバリデーション
	// 正規表現でバリデーションを行う

	emial_regexp := regexp.MustCompile(`^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`)
	if !emial_regexp.MatchString(email) {
		return UserEmail{}, fmt.Errorf("invalid email")
	}

	return UserEmail{email}, nil
}

func (u *UserEmail) Value() string {
	return u.value
}
