package query_dao

import (
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	query_dao_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/dao_if/query"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
	"gorm.io/gorm"
)

type CommunityDaoImpl struct{}

func NewCommunityDaoImpl() query_dao_if.CommunityDao {
	return &CommunityDaoImpl{}
}

func (*CommunityDaoImpl) FindCommunitiesByUserId(
	db *gorm.DB,
	userId user.UserId,
) ([]entity.Community, error) {
	var members []entity.Member
	if err := db.Where("user_id = ?", userId.Value()).Find(&members).Error; err != nil {
		return nil, err
	}
	var communityIds []string
	for _, member := range members {
		communityIds = append(communityIds, member.CommunityId)
	}
	var communities []entity.Community
	if err := db.Where("id IN ?", communityIds).Find(&communities).Error; err != nil {
		return nil, err
	}
	return communities, nil
}

func (*CommunityDaoImpl) FindCommunityById(
	db *gorm.DB,
	communityId community.CommunityId,
) (entity.Community, error) {
	var community entity.Community
	if err := db.Where("id = ?", communityId.Value()).First(&community).Error; err != nil {
		return entity.Community{}, err
	}
	return community, nil
}

func (*CommunityDaoImpl) FindMemberById(
	db *gorm.DB,
	memberId member.MemberId,
) (entity.Member, error) {
	var member entity.Member
	if err := db.First(&member, memberId.Value()).Error; err != nil {
		return entity.Member{}, err
	}
	return member, nil
}

func (*CommunityDaoImpl) FindMembersByCommunityId(
	db *gorm.DB,
	communityId community.CommunityId,
) ([]entity.Member, error) {
	var members []entity.Member
	if err := db.Where("community_id = ?", communityId.Value()).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (*CommunityDaoImpl) FindPlayersByCommunityId(
	db *gorm.DB,
	communityId community.CommunityId,
) ([]entity.Player, error) {
	var players []entity.Player
	if err := db.Where("community_id = ?", communityId.Value()).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}

func (*CommunityDaoImpl) FindPlayersWithStatusByCommunityId(
	db *gorm.DB,
	communityId community.CommunityId,
	status player.PlayerStatus,
) ([]entity.Player, error) {
	var players []entity.Player
	if err := db.
		Where(
			"community_id = ? AND status = ?",
			communityId.Value(),
			status.Value(),
		).
		Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
