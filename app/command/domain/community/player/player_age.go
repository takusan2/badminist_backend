package player

import "errors"

type PlayerAge struct {
	value int
}

func NewPlayerAge(age int) (PlayerAge, error) {
	if age < 0 {
		return PlayerAge{}, errors.New("age must be greater than 0")
	}
	if age > 150 {
		return PlayerAge{}, errors.New("invalid age")
	}
	return PlayerAge{age}, nil
}

func (p *PlayerAge) Value() int {
	return p.value
}
