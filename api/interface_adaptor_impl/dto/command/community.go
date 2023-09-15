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
	PlayerName     string `json:"name"`
	PlayerGender   string `json:"gender"`
	PlayerAge      int    `json:"age"`
	PlayerLevel    string `json:"level"`
	PlayerNumGames int    `json:"num_games"`
	PlayerStatus   string `json:"status"`
}

type RemovePlayerRequestBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type ChangePlayerPropertyRequestBody struct {
	CommunityId    string `json:"community_id"`
	PlayerId       string `json:"player_id"`
	PlayerName     string `json:"name"`
	PlayerGender   string `json:"gender"`
	PlayerAge      int    `json:"age"`
	PlayerLevel    string `json:"level"`
	PlayerNumGames int    `json:"num_games"`
	PlayerStatus   string `json:"status"`
}

type ResetPlayerNumGamesRequestBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
}

type ChangePlayerNumGamesRequestBody struct {
	CommunityId string `json:"community_id"`
	PlayerId    string `json:"player_id"`
	NumGames    int    `json:"num_games"`
}

type AddMemberRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
	MemberRole  string `json:"role"`
}

type RemoveMemberRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
}

type ChangeMemberRoleRequestBody struct {
	CommunityId string `json:"community_id"`
	UserId      string `json:"user_id"`
	MemberRole  string `json:"role"`
}
