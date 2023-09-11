package member

import (
	"errors"

	"github.com/google/uuid"
)

type MemberId struct {
	value string
}

func NewMemberId() MemberId {
	// uuidの形式かどうかのバリデーション
	id := uuid.New().String()
	return MemberId{id}
}

func MemberIdFromStr(id string) (MemberId, error) {
	// uuidの形式かどうかのバリデーション
	if _, err := uuid.Parse(id); err != nil {
		return MemberId{}, errors.New("invalid id")
	}
	return MemberId{id}, nil
}

func (m MemberId) Value() string {
	return m.value
}
