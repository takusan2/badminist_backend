package command_repository_if

import (
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
)

type CommunityRepository interface {
	FindCommunityById(
		communityId community.CommunityId,
	) (community.Community, error)
	// Community
	CreateCommunity(
		communityId community.CommunityId,
		communityName community.CommunityName,
		communityDescription community.CommunityDescription,
		members member.Members,
	) error
	RenameCommunity(
		communityId community.CommunityId,
		communityName community.CommunityName,
	) error
	EditCommunityDescription(
		communityId community.CommunityId,
		communityName community.CommunityName,
		communityDescription community.CommunityDescription,
	) error
	DeleteCommunity(
		communityId community.CommunityId,
	) error

	// Player
	AddPlayer(
		communityId community.CommunityId,
		playerId player.PlayerId,
		playerName player.PlayerName,
		playerGender player.PlayerGender,
		playerAge player.PlayerAge,
		playerLevel player.PlayerLevel,
		playerNumGames player.PlayerNumGames,
		playerStatus player.PlayerStatus,
	) error
	RemovePlayer(
		communityId community.CommunityId,
		playerId player.PlayerId,
	) error
	ChangePlayerProperty(
		communityId community.CommunityId,
		playerId player.PlayerId,
		playerName player.PlayerName,
		playerGender player.PlayerGender,
		playerAge player.PlayerAge,
		playerLevel player.PlayerLevel,
		playerNumGames player.PlayerNumGames,
		playerStatus player.PlayerStatus,
	) error
	ChangePlayerNumGames(
		communityId community.CommunityId,
		playerId player.PlayerId,
		playerNumGames player.PlayerNumGames,
	) error
	// Member
	AddMember(
		memberId member.MemberId,
		communityId community.CommunityId,
		userId user.UserId,
		memberRole member.MemberRole,
	) error
	ChangeMemberRole(
		communityId community.CommunityId,
		userId user.UserId,
		memberRole member.MemberRole,
	) error
	RemoveMember(
		communityId community.CommunityId,
		userId user.UserId,
	) error
}
