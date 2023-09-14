package command_dto

type CreateCommunityRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type RenameCommunityRequestBody struct {
	CommunityId string `json:"community_id"`
	Name        string `json:"name"`
}

type EditCommunityDescriptionRequestBody struct {
	CommunityId string `json:"community_id"`
	Description string `json:"description"`
}

type DeleteCommunityRequestBody struct {
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

type RemovePlayerRequestBody struct {
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

type ResetPlayerNumGamesRequestBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type AddMemberRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
	MemberRole  string `json:"member_role"`
}

type RemoveMemberRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
}

type ChangeMemberRoleRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
	MemberRole  string `json:"member_role"`
}

type GenerateMatchCombinationRequestBody struct {
	CommunityId string `json:"community_id"`
	NumCourt    int    `json:"num_court"`
	Rule        string `json:"rule"`
}
