package community

import (
	"github.com/google/uuid"
)

type CommunityId struct {
	value string
}

func NewCommunityId() CommunityId {
	id := uuid.New().String()
	return CommunityId{value: id}
}

func CommunityIdFromStr(id string) (CommunityId, error) {
	if _, err := uuid.Parse(id); err != nil {
		return CommunityId{}, err
	}
	return CommunityId{value: id}, nil
}

func (c CommunityId) Value() string {
	return c.value
}
