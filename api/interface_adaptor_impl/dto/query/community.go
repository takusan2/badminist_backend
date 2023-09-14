package query_dto

import "github.com/takuya-okada-01/badminist/api/processor/query/read_model"

type GetCommunityListResponseBody struct {
	Communities []read_model.CommunityList `json:"communities"`
}

type GetPlayerListRequestParam struct {
	CommunityId string `json:"community_id"`
}
type GetPlayerListResponseBody struct {
	Players []read_model.PlayerList `json:"players"`
}

type GetMemberListRequestParam struct {
	CommunityId string `json:"community_id"`
}
type GetMemberListResponseBody struct {
	Members []read_model.MemberList `json:"members"`
}

type GenerateMatchCombinationRequestParam struct {
	CommunityId string `json:"community_id"`
	NumCourt    int    `json:"num_court"`
	Rule        string `json:"rule"`
}
type GenerateMatchCombinationResponseBody struct {
	MatchCombination read_model.MatchCombination `json:"match_combination"`
}
