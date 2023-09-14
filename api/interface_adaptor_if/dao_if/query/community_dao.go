package query_dao_if

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
	FindCommunitiesByUserId(
		db *gorm.DB,
		userId user.UserId,
	) ([]entity.Community, error)
	// Member
	FindMemberById(
		db *gorm.DB,
		memberId member.MemberId,
	) (entity.Member, error)
	FindMembersByCommunityId(
		db *gorm.DB,
		communityId community.CommunityId,
	) ([]entity.Member, error)
	// Player
	FindPlayersWithStatusByCommunityId(
		db *gorm.DB,
		communityId community.CommunityId,
		status player.PlayerStatus,
	) ([]entity.Player, error)
	FindPlayersByCommunityId(
		db *gorm.DB,
		communityId community.CommunityId,
	) ([]entity.Player, error)
}
