package player

import (
	"errors"
	"fmt"
	"regexp"
)

type PlayerName struct {
	value string
}

func NewPlayerName(name string) (PlayerName, error) {
	if name == "" {
		return PlayerName{}, errors.New("Invalid player name")
	}
	if len(name) > 255 {
		return PlayerName{}, fmt.Errorf("user name is too long")
	}
	regexp_pattern := regexp.MustCompile(`^[ 　]+$`)
	if regexp_pattern.MatchString(name) {
		return PlayerName{}, fmt.Errorf("user name is invalid")
	}
	return PlayerName{value: name}, nil
}

func (p *PlayerName) Value() string {
	return p.value
}
