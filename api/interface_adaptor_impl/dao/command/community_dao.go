package command_dao

import (
	"errors"

	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	command_dao_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/dao_if/command"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"

	"gorm.io/gorm"
)

type CommunityDaoImpl struct{}

func NewCommunityDaoImpl() command_dao_if.CommunityDao {
	return &CommunityDaoImpl{}
}

// FindCommunitiesByUserId implements dao_if.CommunityDao.
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

// FindCommunityById implements dao_if.CommunityDao.
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

// FindMemberById implements dao_if.CommunityDao.
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

// FindMembersByCommunityId implements dao_if.CommunityDao.
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

// FindPlayerById implements dao_if.CommunityDao.
func (*CommunityDaoImpl) FindPlayerById(
	db *gorm.DB,
	communityId community.CommunityId,
	playerId player.PlayerId,
) (entity.Player, error) {
	var player entity.Player
	if err := db.
		Where(
			"community_id = ? AND id = ?",
			communityId.Value(),
			playerId.Value(),
		).
		First(&player).Error; err != nil {
		return entity.Player{}, err
	}
	return player, nil
}

// FindPlayersByCommunityId implements dao_if.CommunityDao.
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

// InsertCommunity implements dao_if.CommunityDao.
func (*CommunityDaoImpl) InsertCommunity(
	db *gorm.DB,
	community entity.Community,
) error {
	if err := db.Create(&community).Error; err != nil {
		switch err.Error() {
		case "Error 1062: Duplicate entry 'community_id' for key 'PRIMARY'":
			return errors.New("既に存在するコミュニティIDです")
		default:
			return err
		}
	}
	return nil
}

// InsertMember implements dao_if.CommunityDao.
func (*CommunityDaoImpl) InsertMember(
	db *gorm.DB,
	member entity.Member,
) error {
	if err := db.Create(&member).Error; err != nil {
		return err
	}
	return nil
}

// InsertMembers implements dao_if.CommunityDao.
func (*CommunityDaoImpl) InsertMembers(
	db *gorm.DB,
	members []entity.Member,
) error {
	if err := db.Create(&members).Error; err != nil {
		return err
	}
	return nil
}

// InsertPlayer implements dao_if.CommunityDao.
func (*CommunityDaoImpl) InsertPlayer(
	db *gorm.DB,
	player entity.Player,
) error {
	if err := db.Create(&player).Error; err != nil {
		return err
	}
	return nil
}

func (*CommunityDaoImpl) DeleteCommunity(
	db *gorm.DB,
	communityId community.CommunityId,
) error {
	if err := db.
		Delete(
			&entity.Community{},
			"id = ?",
			communityId.Value(),
		).Error; err != nil {
		return err
	}
	return nil
}

// RemoveMember implements dao_if.CommunityDao.
func (*CommunityDaoImpl) DeleteMember(
	db *gorm.DB,
	communityId community.CommunityId,
	userId user.UserId,
) error {
	if err := db.Delete(
		&entity.Member{},
		"community_id = ? AND user_id = ?",
		communityId.Value(),
		userId.Value(),
	).Error; err != nil {
		return err
	}
	return nil
}

// RemovePlayer implements dao_if.CommunityDao.
func (*CommunityDaoImpl) DeletePlayer(
	db *gorm.DB,
	communityId community.CommunityId,
	playerId player.PlayerId,
) error {
	if err := db.Delete(
		&entity.Player{},
		"community_id = ? AND id = ?",
		communityId.Value(),
		playerId.Value(),
	).Error; err != nil {
		return err
	}
	return nil
}

// UpdateCommunity implements dao_if.CommunityDao.
func (*CommunityDaoImpl) UpdateCommunityName(
	db *gorm.DB,
	communityId community.CommunityId,
	communityName community.CommunityName,
) error {
	if err := db.Model(&entity.Community{}).
		Where(
			"id = ?",
			communityId.Value(),
		).
		Update(
			"name",
			communityName.Value(),
		).Error; err != nil {
		return err
	}
	return nil
}

func (*CommunityDaoImpl) UpdateCommunityDescription(
	db *gorm.DB,
	communityId community.CommunityId,
	communityDescription community.CommunityDescription,
) error {
	if err := db.Model(&entity.Community{}).
		Where(
			"id = ?",
			communityId.Value(),
		).
		Update(
			"description",
			communityDescription.Value(),
		).Error; err != nil {
		return err
	}
	return nil
}

// UpdateMember implements dao_if.CommunityDao.
func (*CommunityDaoImpl) UpdateMember(
	db *gorm.DB,
	communityId community.CommunityId,
	userId user.UserId,
	memberRole member.MemberRole,
) error {
	if err := db.Model(&entity.Member{}).
		Where(
			"community_id = ? AND user_id = ?",
			communityId.Value(),
			userId.Value(),
		).
		Update(
			"role",
			memberRole.Value(),
		).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePlayer implements dao_if.CommunityDao.
func (*CommunityDaoImpl) UpdatePlayer(
	db *gorm.DB,
	communityId community.CommunityId,
	playerId player.PlayerId,
	playerName player.PlayerName,
	playerGender player.PlayerGender,
	playerAge player.PlayerAge,
	playerLevel player.PlayerLevel,
	playerNumGames player.PlayerNumGames,
	playerStatus player.PlayerStatus,
) error {
	if err := db.Model(&entity.Player{}).
		Where(
			"community_id = ? AND id = ?",
			communityId.Value(),
			playerId.Value(),
		).Updates(
		map[string]any{
			"name":      playerName.Value(),
			"gender":    playerGender.Value(),
			"age":       playerAge.Value(),
			"level":     playerLevel.Value(),
			"num_games": playerNumGames.Value(),
			"status":    playerStatus.Value(),
		},
	).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePlayerNumGames implements dao_if.CommunityDao.
func (*CommunityDaoImpl) UpdatePlayerNumGames(
	db *gorm.DB,
	communityId community.CommunityId,
	playerId player.PlayerId,
	num player.PlayerNumGames,
) error {
	if err := db.Model(&entity.Player{}).
		Where(
			"community_id = ? AND id = ?",
			communityId.Value(),
			playerId.Value(),
		).
		Update(
			"num_games",
			num.Value(),
		).Error; err != nil {
		return err
	}
	return nil
}

func (*CommunityDaoImpl) FindPlayersWithStatusByCommunityId(
	db *gorm.DB,
	communityId community.CommunityId,
	status player.PlayerStatus,
) ([]entity.Player, error) {
	var players []entity.Player
	if err := db.Where(
		"community_id = ? AND status = ?",
		communityId.Value(),
		status.Value(),
	).Find(&players).Error; err != nil {
		return nil, err
	}
	return players, nil
}
