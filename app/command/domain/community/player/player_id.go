package player

import (
	"github.com/google/uuid"
)

type PlayerId struct {
	value string
}

func NewPlayerId() PlayerId {
	id := uuid.New().String()
	return PlayerId{id}
}

func PlayerIdFromStr(id string) (PlayerId, error) {
	if _, err := uuid.Parse(id); err != nil {
		return PlayerId{}, err
	}
	return PlayerId{id}, nil
}

func (p PlayerId) Value() string {
	return p.value
}
