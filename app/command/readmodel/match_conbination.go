package readmodel

import "github.com/takuya-okada-01/badminist/app/infrastructure/entity"

type MatchCombination struct {
	Matches    []Match         `json:"matches"`
	RestPlayer []entity.Player `json:"rest_player"`
}
