package readmodel

import "github.com/takuya-okada-01/badminist/api/infrastructure/entity"

type PlayerList struct {
	Players []entity.Player `json:"players"`
}
