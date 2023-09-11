package user

import (
	"errors"

	"github.com/google/uuid"
)

type UserId struct {
	value string
}

func NewUserId() UserId {
	id := uuid.New().String()
	return UserId{id}
}

func UserIdFromStr(id string) (UserId, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return UserId{}, errors.New("invalid id")
	}
	return UserId{id}, nil
}

func (u UserId) Value() string {
	return u.value
}
