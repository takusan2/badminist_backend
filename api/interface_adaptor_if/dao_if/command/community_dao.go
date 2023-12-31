package command_dao_if

import (
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
	"gorm.io/gorm"
)

type CommunityDao interface {
	// Community
	FindCommunityById(
		db *gorm.DB,
		communityId community.CommunityId,
	) (entity.Community, error)
	InsertCommunity(
		db *gorm.DB,
		community entity.Community,
	) error
	UpdateCommunityName(
		db *gorm.DB,
		communityId community.CommunityId,
		communityName community.CommunityName,
	) error
	UpdateCommunityDescription(
		db *gorm.DB,
		communityId community.CommunityId,
		communityDescription community.CommunityDescription,
	) error
	DeleteCommunity(
		db *gorm.DB,
		communityId community.CommunityId,
	) error

	// Player
	FindPlayerById(
		db *gorm.DB,
		communityId community.CommunityId,
		playerId player.PlayerId,
	) (entity.Player, error)
	FindPlayersByCommunityId(
		db *gorm.DB,
		communityId community.CommunityId,
	) ([]entity.Player, error)
	InsertPlayer(
		db *gorm.DB,
		player entity.Player,
	) error
	DeletePlayer(
		db *gorm.DB,
		communityId community.CommunityId,
		playerId player.PlayerId,
	) error
	UpdatePlayer(
		db *gorm.DB,
		communityId community.CommunityId,
		playerId player.PlayerId,
		playerName player.PlayerName,
		playerGender player.PlayerGender,
		playerAge player.PlayerAge,
		playerLevel player.PlayerLevel,
		playerNumGames player.PlayerNumGames,
		playerStatus player.PlayerStatus,
	) error
	UpdatePlayerNumGames(
		db *gorm.DB,
		communityId community.CommunityId,
		playerId player.PlayerId,
		num player.PlayerNumGames,
	) error

	// Member
	FindMemberById(
		db *gorm.DB,
		memberId member.MemberId,
	) (entity.Member, error)
	FindMembersByCommunityId(
		db *gorm.DB,
		communityId community.CommunityId,
	) ([]entity.Member, error)
	InsertMember(
		db *gorm.DB,
		member entity.Member,
	) error
	InsertMembers(
		db *gorm.DB,
		members []entity.Member,
	) error
	UpdateMember(
		db *gorm.DB,
		communityId community.CommunityId,
		userId user.UserId,
		memberRole member.MemberRole,
	) error
	DeleteMember(
		db *gorm.DB,
		communityId community.CommunityId,
		userId user.UserId,
	) error
}
