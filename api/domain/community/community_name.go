package community

import (
	"fmt"
	"regexp"
)

type CommunityName struct {
	value string
}

func NewCommunityName(name string) (CommunityName, error) {
	if name == "" {
		return CommunityName{}, fmt.Errorf("user name is empty")
	}
	if len(name) > 255 {
		return CommunityName{}, fmt.Errorf("user name is too long")
	}
	// 半角空白または全角空白のみの場合はエラー
	regexp_pattern := regexp.MustCompile(`^[ 　]+$`)
	if regexp_pattern.MatchString(name) {
		return CommunityName{}, fmt.Errorf("user name is invalid")
	}

	return CommunityName{name}, nil
}

func (c *CommunityName) Value() string {
	return c.value
}
