package entity

import (
	"time"

	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/processor/query/read_model"
)

type Player struct {
	Id          string    `gorm:"type:varchar(36);primaryKey;"`
	CommunityId string    `gorm:"type:varchar(36);not null;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	Gender      string    `gorm:"type:varchar(255);not null;"`
	Age         int       `gorm:"type:int(11);not null;"`
	Level       string    `gorm:"type:varchar(255);not null;"`
	NumGames    int       `gorm:"type:int(11);not null;"`
	Status      string    `gorm:"type:varchar(255);not null;"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func NewPlayer(
	id string,
	communityId string,
	name string,
	gender string,
	age int,
	level string,
	numGames int,
	status string,
) Player {
	return Player{
		Id:          id,
		CommunityId: communityId,
		Name:        name,
		Gender:      gender,
		Age:         age,
		Level:       level,
		NumGames:    numGames,
		Status:      status,
	}
}

func (p *Player) ToDomainObject() player.Player {
	playerId, err := player.PlayerIdFromStr(p.Id)
	if err != nil {
		panic(err)
	}

	name, err := player.NewPlayerName(p.Name)
	if err != nil {
		panic(err)
	}

	gender, err := player.PlayerGenderFromStr(p.Gender)
	if err != nil {
		panic(err)
	}

	age, err := player.NewPlayerAge(p.Age)
	if err != nil {
		panic(err)
	}

	level, err := player.PlayerLevelFromStr(p.Level)
	if err != nil {
		panic(err)
	}

	numGames, err := player.NewPlayerNumGames(p.NumGames)
	if err != nil {
		panic(err)
	}

	status, err := player.PlayerStatusFromStr(p.Status)
	if err != nil {
		panic(err)
	}

	return player.NewPlayer(
		playerId,
		name,
		gender,
		age,
		level,
		numGames,
		status,
	)
}

func (p *Player) ToReadModel() read_model.Player {
	return read_model.NewPlayer(
		p.Id,
		p.Name,
		p.Gender,
		p.Age,
		p.Level,
		p.NumGames,
		p.Status,
	)
}
