package player

type Players struct {
	value []Player
}

func NewPlayers(players []Player) Players {
	return Players{
		value: players,
	}
}

func PlayersFromList(players []Player) Players {
	return NewPlayers(players)
}

func (p *Players) AddPlayer(player Player) {
	p.value = append(p.value, player)
}

func (p *Players) RemovePlayer(remoevPlayerId PlayerId) {
	for i, player := range p.value {
		if player.id == remoevPlayerId {
			p.value = append(p.value[:i], p.value[i+1:]...)
		}
	}
}

func (p *Players) IsPlayer(playerId PlayerId) bool {
	for _, p := range p.value {
		if p.id == playerId {
			return true
		}
	}
	return false
}

func (p *Players) ChangePlayerProperty(changePlayer Player) {
	for i, player := range p.value {
		if player.id == changePlayer.id {
			p.value[i].ChangePlayerProperty(changePlayer)
		}
	}
}

func (p *Players) ResetPlayerNumGames(playerId PlayerId) {
	for i, player := range p.value {
		if player.id == playerId {
			p.value[i].ResetNumGames()
		}
	}
}

func (p *Players) GetPlayer(playerId PlayerId) Player {
	for _, player := range p.value {
		if player.id == playerId {
			return player
		}
	}
	return Player{}
}

func (p *Players) Value() []Player {
	return p.value
}

func (p *Players) Len() int {
	return len(p.value)
}
