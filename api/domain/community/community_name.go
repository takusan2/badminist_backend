package community

import (
	"errors"
	"regexp"
)

type CommunityName struct {
	value string
}

func NewCommunityName(name string) (CommunityName, error) {
	if name == "" {
		return CommunityName{}, errors.New("コミュニティ名を入力してください")
	}
	if len(name) > 255 {
		return CommunityName{}, errors.New("コミュニティ名は255文字以内で入力してください")
	}
	// 半角空白または全角空白のみの場合はエラー
	regexp_pattern := regexp.MustCompile(`^[ 　]+$`)
	if regexp_pattern.MatchString(name) {
		return CommunityName{}, errors.New("空白のみのコミュニティ名は登録できません")
	}

	return CommunityName{name}, nil
}

func (c *CommunityName) Value() string {
	return c.value
}
