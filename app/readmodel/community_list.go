package readmodel

import "github.com/takuya-okada-01/badminist/app/infrastructure/entity"

type CommunityList struct {
	Communities []entity.Community `json:"communities"`
	NumPlayers  int                `json:"num_players"`
	NumMembers  int                `json:"num_members"`
}
