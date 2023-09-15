package player

type Player struct {
	id       PlayerId
	name     PlayerName
	gender   PlayerGender
	age      PlayerAge
	level    PlayerLevel
	numGames PlayerNumGames
	status   PlayerStatus
}

func NewPlayer(
	id PlayerId,
	name PlayerName,
	gender PlayerGender,
	age PlayerAge,
	level PlayerLevel,
	numGames PlayerNumGames,
	status PlayerStatus,
) Player {
	return Player{
		id:       id,
		name:     name,
		gender:   gender,
		age:      age,
		level:    level,
		numGames: numGames,
		status:   status,
	}
}

func (p *Player) BreachEncapsulationOfId() PlayerId {
	return p.id
}
func (p *Player) BreachEncapsulationOfName() PlayerName {
	return p.name
}
func (p *Player) BreachEncapsulationOfGender() PlayerGender {
	return p.gender
}
func (p *Player) BreachEncapsulationOfAge() PlayerAge {
	return p.age
}
func (p *Player) BreachEncapsulationOfLevel() PlayerLevel {
	return p.level
}
func (p *Player) BreachEncapsulationOfNumGames() PlayerNumGames {
	return p.numGames
}
func (p *Player) BreachEncapsulationOfStatus() PlayerStatus {
	return p.status
}

func (p *Player) Rename(name PlayerName) {
	p.name = name
}

func (p *Player) ChangeGender(gender PlayerGender) {
	p.gender = gender
}

func (p *Player) ChangeAge(age PlayerAge) {
	p.age = age
}

func (p *Player) ChangeLevel(level PlayerLevel) {
	p.level = level
}

func (p *Player) ChangeStatus(status PlayerStatus) {
	p.status = status
}

func (p *Player) ChangePlayerProperty(player Player) {
	p.name = player.name
	p.gender = player.gender
	p.age = player.age
	p.level = player.level
	p.status = player.status
}

func (p *Player) ResetNumGames() {
	p.numGames.Reset()
}

func (p *Player) ChangeNumGames(numGames PlayerNumGames) {
	p.numGames = numGames
}
