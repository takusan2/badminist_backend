package community

import (
	"errors"
)

type CommunityDescription struct {
	value string
}

func NewCommunityDescription(description string) (CommunityDescription, error) {
	if len(description) > 255 {
		return CommunityDescription{}, errors.New("コミュニティの説明は255文字以内で入力してください")
	}
	return CommunityDescription{description}, nil
}

func (c *CommunityDescription) Value() string {
	return c.value
}
