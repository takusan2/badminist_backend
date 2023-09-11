package repository_if

import (
	"github.com/takuya-okada-01/badminist/app/command/domain/community"
	"github.com/takuya-okada-01/badminist/app/command/domain/community/member"
	"github.com/takuya-okada-01/badminist/app/command/domain/community/player"
	"github.com/takuya-okada-01/badminist/app/command/domain/user"
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
		plyaerGender player.PlayerGender,
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
		plyaerGender player.PlayerGender,
		playerAge player.PlayerAge,
		playerLevel player.PlayerLevel,
		playerNumGames player.PlayerNumGames,
		playerStatus player.PlayerStatus,
	) error
	ResetPlayerNumGames(
		communityId community.CommunityId,
		playerId player.PlayerId,
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
