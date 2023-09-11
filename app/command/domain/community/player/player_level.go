package player

import "errors"

type PlayerLevel struct {
	value string
}

type Level int

const (
	Beginner Level = iota
	Intermediate
	Advanced
)

func (l Level) String() string {
	switch l {
	case Advanced:
		return "advanced"
	case Intermediate:
		return "intermediate"
	default:
		return "beginner"
	}
}

func PlayerLevelFromStr(level string) (PlayerLevel, error) {
	switch level {
	case Advanced.String():
		level, _ := NewPlayerLevel(Advanced)
		return level, nil
	case Intermediate.String():
		level, _ := NewPlayerLevel(Intermediate)
		return level, nil
	case Beginner.String():
		level, _ := NewPlayerLevel(Beginner)
		return level, nil
	default:
		return PlayerLevel{}, errors.New("invalid player level")
	}
}

func NewPlayerLevel(level Level) (PlayerLevel, error) {
	return PlayerLevel{level.String()}, nil
}

func (p *PlayerLevel) Value() string {
	return p.value
}
