package readmodel

import "github.com/takuya-okada-01/badminist/api/infrastructure/entity"

type MemberList struct {
	Members []entity.Member `json:"members"`
}