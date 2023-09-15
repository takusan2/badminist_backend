package player

import (
	"errors"
)

type PlayerNumGames struct {
	value int
}

func NewPlayerNumGames(numGames int) (PlayerNumGames, error) {
	if numGames < 0 {
		return PlayerNumGames{}, errors.New("numGames is invalid")
	}
	return PlayerNumGames{value: numGames}, nil
}

func (p *PlayerNumGames) Change(num PlayerNumGames) {
	p.value = num.value
}

func (p *PlayerNumGames) Reset() {
	p.value = 0
}

func (p *PlayerNumGames) Value() int {
	return p.value
}
