package dto

type CreateCommunityRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type CreateCommunityResponseBody struct {
	CommunityId string `json:"community_id"`
}

type RenameCommunityRequestBody struct {
	CommunityId string `json:"community_id"`
	Name        string `json:"name"`
}
type RenameCommunityResponseBody struct {
	CommunityId string `json:"community_id"`
}

type EditCommunityDescriptionRequestBody struct {
	CommunityId string `json:"community_id"`
	Description string `json:"description"`
}
type EditCommunityDescriptionResponseBody struct {
	CommunityId string `json:"community_id"`
}

type DeleteCommunityRequestBody struct {
	CommunityId string `json:"community_id"`
}
type DeleteCommunityResponseBody struct {
	CommunityId string `json:"community_id"`
}

type AddPlayerRequestBody struct {
	CommunityId    string `json:"community_id"`
	PlayerName     string `json:"player_name"`
	PlayerGender   string `json:"player_gender"`
	PlayerAge      int    `json:"player_age"`
	PlayerLevel    string `json:"player_level"`
	PlayerNumGames int    `json:"player_num_games"`
	PlayerStatus   string `json:"player_status"`
}
type AddPlayerResponseBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type RemovePlayerRequestBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}
type RemovePlayerResponseBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type ChangePlayerPropertyRequestBody struct {
	CommunityId    string `json:"community_id"`
	PlayerId       string `json:"player_id"`
	PlayerName     string `json:"player_name"`
	PlayerGender   string `json:"player_gender"`
	PlayerAge      int    `json:"player_age"`
	PlayerLevel    string `json:"player_level"`
	PlayerNumGames int    `json:"player_num_games"`
	PlayerStatus   string `json:"player_status"`
	ExecutorId     string `json:"executor_id"`
}
type ChangePlayerPropertyResponseBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type ResetPlayerNumGamesRequestBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}
type ResetPlayerNumGamesResponseBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type AddMemberRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
	MemberRole  string `json:"membe_role"`
}
type AddMemberResponseBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
}

type RemoveMemberRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
}
type RemoveMemberResponseBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
}

type ChangeMemberRoleRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
	MemberRole  string `json:"member_role"`
}
type ChangeMemberRoleResponseBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
}

type GenerateMatchCombinationRequestBody struct {
	CommunityId string `json:"community_id"`
	NumCourt    int    `json:"num_court"`
	Rule        string `json:"rule"`
}
