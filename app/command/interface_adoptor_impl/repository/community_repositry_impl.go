package repository

import (
	"github.com/takuya-okada-01/badminist/app/command/domain/community"
	"github.com/takuya-okada-01/badminist/app/command/domain/community/member"
	"github.com/takuya-okada-01/badminist/app/command/domain/community/player"
	"github.com/takuya-okada-01/badminist/app/command/domain/user"
	"github.com/takuya-okada-01/badminist/app/command/interface_adaptor_if/dao_if"
	"github.com/takuya-okada-01/badminist/app/command/interface_adaptor_if/repository_if.go"
	"github.com/takuya-okada-01/badminist/app/infrastructure/entity"
	"gorm.io/gorm"
)

type communityRepositoryImpl struct {
	db  *gorm.DB
	dao dao_if.CommunityDao
}

func NewCommunityRepositoryImpl(
	db *gorm.DB,
	dao dao_if.CommunityDao,
) repository_if.CommunityRepository {
	return &communityRepositoryImpl{
		db:  db,
		dao: dao,
	}
}

// ResetPlayerNumGames implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) ResetPlayerNumGames(communityId community.CommunityId, playerId player.PlayerId) error {
	if err := c.dao.UpdatePlayerNumGames(
		c.db,
		communityId,
		playerId,
		0,
	); err != nil {
		return err
	}
	return nil
}

// AddMember implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) AddMember(memberId member.MemberId, communityId community.CommunityId, userId user.UserId, memberRole member.MemberRole) error {
	if err := c.dao.InsertMember(
		c.db,
		entity.NewMember(
			memberId.Value(),
			userId.Value(),
			communityId.Value(),
			memberRole.Value(),
		),
	); err != nil {
		return err
	}
	return nil
}

// AddPlayer implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) AddPlayer(communityId community.CommunityId, playerId player.PlayerId, playerName player.PlayerName, plyaerGender player.PlayerGender, playerAge player.PlayerAge, playerLevel player.PlayerLevel, playerNumGames player.PlayerNumGames, playerStatus player.PlayerStatus) error {
	if err := c.dao.InsertPlayer(
		c.db,
		entity.NewPlayer(
			playerId.Value(),
			communityId.Value(),
			playerName.Value(),
			plyaerGender.Value(),
			playerAge.Value(),
			playerLevel.Value(),
			playerNumGames.Value(),
			playerStatus.Value(),
		),
	); err != nil {
		return err
	}
	return nil
}

// ChangeMemberRole implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) ChangeMemberRole(communityId community.CommunityId, userId user.UserId, memberRole member.MemberRole) error {
	if err := c.dao.UpdateMember(
		c.db,
		communityId,
		userId,
		memberRole,
	); err != nil {
		return err
	}
	return nil
}

// ChangePlayerProperty implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) ChangePlayerProperty(communityId community.CommunityId, playerId player.PlayerId, playerName player.PlayerName, plyaerGender player.PlayerGender, playerAge player.PlayerAge, playerLevel player.PlayerLevel, playerNumGames player.PlayerNumGames, playerStatus player.PlayerStatus) error {
	if err := c.dao.UpdatePlayer(
		c.db,
		communityId,
		playerId,
		playerName,
		plyaerGender,
		playerAge,
		playerLevel,
		playerNumGames,
		playerStatus,
	); err != nil {
		return err
	}
	return nil
}

// DeleteCommunity implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) DeleteCommunity(communityId community.CommunityId) error {
	if err := c.dao.DeleteCommunity(
		c.db,
		communityId,
	); err != nil {
		return err
	}
	return nil
}

func (c *communityRepositoryImpl) RenameCommunity(communityId community.CommunityId, communityName community.CommunityName) error {
	if err := c.dao.UpdateCommunityName(
		c.db,
		communityId,
		communityName,
	); err != nil {
		return err
	}
	return nil
}

// EditCommunityDescription implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) EditCommunityDescription(communityId community.CommunityId, communityName community.CommunityName, communityDescription community.CommunityDescription) error {
	if err := c.dao.UpdateCommunityDescription(
		c.db,
		communityId,
		communityDescription,
	); err != nil {
		return err
	}
	return nil
}

// RemoveMember implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) RemoveMember(communityId community.CommunityId, userId user.UserId) error {
	if err := c.dao.DeleteMember(
		c.db,
		communityId,
		userId,
	); err != nil {
		return err
	}
	return nil
}

// RemovePlayer implements repository_if.CommunityRepository.
func (c *communityRepositoryImpl) RemovePlayer(communityId community.CommunityId, playerId player.PlayerId) error {
	if err := c.dao.DeletePlayer(
		c.db,
		communityId,
		playerId,
	); err != nil {
		return err
	}
	return nil
}

func (r *communityRepositoryImpl) FindCommunityById(
	communityId community.CommunityId,
) (community.Community, error) {
	communityEntity, err := r.dao.FindCommunityById(r.db, communityId)
	if err != nil {
		return community.Community{}, err
	}
	playerEntityList, err := r.dao.FindPlayersByCommunityId(r.db, communityId)
	if err != nil {
		return community.Community{}, err
	}
	memberEntityList, err := r.dao.FindMembersByCommunityId(r.db, communityId)
	if err != nil {
		return community.Community{}, err
	}
	communityId, err = community.CommunityIdFromStr(communityEntity.Id)
	if err != nil {
		return community.Community{}, err
	}
	communityName, err := community.NewCommunityName(communityEntity.Name)
	if err != nil {
		return community.Community{}, err
	}
	communityDescription, err := community.NewCommunityDescription(communityEntity.Description)
	if err != nil {
		return community.Community{}, err
	}

	playerList := []player.Player{}
	for _, playerEntities := range playerEntityList {
		player := playerEntities.ToDomainObject()
		if err != nil {
			return community.Community{}, err
		}
		playerList = append(playerList, player)
	}

	memberList := []member.Member{}
	for _, memberEntities := range memberEntityList {
		member := memberEntities.ToDomainObject()
		if err != nil {
			return community.Community{}, err
		}
		memberList = append(memberList, member)
	}
	players := player.PlayersFromList(playerList)
	members := member.MembersFromList(memberList)

	community := community.NewCommunity(
		communityId,
		communityName,
		communityDescription,
		players,
		members,
	)
	return community, nil

}

func (r *communityRepositoryImpl) CreateCommunity(
	communityId community.CommunityId,
	communityName community.CommunityName,
	communityDescription community.CommunityDescription,
	members member.Members,
) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err := tx.Error; err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}()

	communityEntity := entity.NewCommunity(
		communityId.Value(),
		communityName.Value(),
		communityDescription.Value(),
	)
	if err := r.dao.InsertCommunity(
		tx,
		communityEntity,
	); err != nil {
		tx.Rollback()
		return err
	}

	for _, member := range members.Value() {
		memberEntity := entity.NewMember(
			member.BreachEncapsulationOfId().Value(),
			member.BreachEncapsulationOfUserId().Value(),
			communityId.Value(),
			member.BreachEncapsulationOfRole().Value(),
		)
		if err := r.dao.InsertMember(
			tx,
			memberEntity,
		); err != nil {
			tx.Rollback()
			return err
		}
	}
	return nil
}
