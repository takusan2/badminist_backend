package player

import "errors"

type PlayerGender struct {
	value string
}

type Gender int

const (
	Unknown Gender = iota
	Male
	Female
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return "unknown"
	}
}

func NewPlayerGender(gender Gender) (PlayerGender, error) {
	return PlayerGender{gender.String()}, nil
}

func PlayerGenderFromStr(gender string) (PlayerGender, error) {
	switch gender {
	case Male.String():
		gender, _ := NewPlayerGender(Male)
		return gender, nil
	case Female.String():
		gender, _ := NewPlayerGender(Female)
		return gender, nil
	case Unknown.String():
		gender, _ := NewPlayerGender(Unknown)
		return gender, nil
	}
	return PlayerGender{}, errors.New("invalid gender")
}

func (p *PlayerGender) Value() string {
	return p.value
}
