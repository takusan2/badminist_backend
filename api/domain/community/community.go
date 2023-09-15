package community

import (
	"fmt"
	"time"

	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
)

type Community struct {
	id          CommunityId
	name        CommunityName
	description CommunityDescription
	players     player.Players
	members     member.Members
}

func NewCommunity(
	id CommunityId,
	name CommunityName,
	description CommunityDescription,
	players player.Players,
	members member.Members,
) Community {
	return Community{
		id:          id,
		name:        name,
		description: description,
		players:     players,
		members:     members,
	}
}

func (c *Community) SetPlayers(players player.Players) {
	c.players = players
}
func (c *Community) SetMembers(members member.Members) {
	c.members = members
}

func (c *Community) BreachEncapsulationOfId() CommunityId {
	return c.id
}
func (c *Community) BreachEncapsulationOfName() CommunityName {
	return c.name
}
func (c *Community) BreachEncapsulationOfDescription() CommunityDescription {
	return c.description
}
func (c *Community) BreachEncapsulationOfPlayers() player.Players {
	return c.players
}
func (c *Community) BreachEncapsulationOfMembers() member.Members {
	return c.members
}

// コミュニティの名前を変更する
//
// # 引数
//
//   - name: 変更後のコミュニティ名
//   - executor: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
func (c *Community) RenameCommunity(
	name CommunityName,
	executorId user.UserId,
) (CommunityEventRenamedBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventRenamedBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventRenamedBody{}, fmt.Errorf("executor is not staff")
	}
	c.name = name
	return CommunityEventRenamedBody{
		CommunityId: c.id,
		Name:        c.name,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティの説明を演習する
//
// # 引数
//
//   - description: 変更後のコミュニティ説明
//   - executorId: 実行者のユーザId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
func (c *Community) EditDescription(
	description CommunityDescription,
	executorId user.UserId,
) (CommunityEventEditDescriptionBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventEditDescriptionBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventEditDescriptionBody{}, fmt.Errorf("executor is not staff")
	}
	c.description = description
	return CommunityEventEditDescriptionBody{
		CommunityId: c.id,
		Description: c.description,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティを削除する
//
// # 引数
//
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
func (c *Community) DeleteCommunity(
	executorId user.UserId,
) (CommunityEventDeletedBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventDeletedBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsAdmin(executorId) {
		return CommunityEventDeletedBody{}, fmt.Errorf("executor is not admin")
	}
	return CommunityEventDeletedBody{
		CommunityId: c.id,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティにプレイヤーを追加する
//
// # 引数
//
//   - playerId: 追加するプレイヤーのId
//   - playerName: 追加するプレイヤーの名前
//   - playerAge: 追加するプレイヤーの年齢
//   - playerLevel: 追加するプレイヤーのレベル
//   - playerGender: 追加するプレイヤーの性別
//   - playerStatus: 追加するプレイヤーのステータス
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
func (c *Community) AddPlayer(
	player player.Player,
	executorId user.UserId,
) (CommunityEventAddPlayerBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventAddPlayerBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventAddPlayerBody{}, fmt.Errorf("executor is not staff")
	}
	if c.players.IsPlayer(player.BreachEncapsulationOfId()) {
		return CommunityEventAddPlayerBody{}, fmt.Errorf("player is already player")
	}

	return CommunityEventAddPlayerBody{
		CommunityId: c.id,
		Player:      player,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティからプレイヤーを削除する
//
// # 引数
//
//   - playerId: 削除するプレイヤーのId
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - 削除するプレイヤーが存在しない場合はエラー
func (c *Community) RemovePlayer(
	playerId player.PlayerId,
	executorId user.UserId,
) (CommunityEventRemovePlayerBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventRemovePlayerBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventRemovePlayerBody{}, fmt.Errorf("executor is not staff")
	}
	if !c.players.IsPlayer(playerId) {
		return CommunityEventRemovePlayerBody{}, fmt.Errorf("player is not player")
	}
	c.players.RemovePlayer(playerId)
	return CommunityEventRemovePlayerBody{
		CommunityId: c.id,
		PlayerId:    playerId,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティのプレイヤーのプロパティを変更する
//
// # 引数
//
//   - playerId: 変更するプレイヤーのId
//   - player: 変更後のプレイヤー
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - 変更するプレイヤーが存在しない場合はエラー
func (c *Community) ChangePlayerProperty(
	playerId player.PlayerId,
	playerName player.PlayerName,
	playerGender player.PlayerGender,
	playerAge player.PlayerAge,
	playerLevel player.PlayerLevel,
	playerNumGames player.PlayerNumGames,
	playerStatus player.PlayerStatus,
	executorId user.UserId,
) (CommunityEventChangePlayerPropertyBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventChangePlayerPropertyBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventChangePlayerPropertyBody{}, fmt.Errorf("executor is not staff")
	}
	if !c.players.IsPlayer(playerId) {
		return CommunityEventChangePlayerPropertyBody{}, fmt.Errorf("player is not player")
	}
	player := player.NewPlayer(
		playerId,
		playerName,
		playerGender,
		playerAge,
		playerLevel,
		playerNumGames,
		playerStatus,
	)
	c.players.ChangePlayerProperty(player)
	return CommunityEventChangePlayerPropertyBody{
		CommunityId: c.id,
		Player:      player,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティのプレイヤーの試合数をリセットする
//
// # 引数
//
//   - playerId: リセットするプレイヤーのId
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - リセットするプレイヤーが存在しない場合はエラー
func (c *Community) ResetPlayerNumGames(
	playerId player.PlayerId,
	executorId user.UserId,
) (CommunityEventResetNumGamesBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventResetNumGamesBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventResetNumGamesBody{}, fmt.Errorf("executor is not staff")
	}
	if !c.players.IsPlayer(playerId) {
		return CommunityEventResetNumGamesBody{}, fmt.Errorf("player is not player")
	}
	c.players.ResetPlayerNumGames(playerId)
	numGames, _ := player.NewPlayerNumGames(0)
	return CommunityEventResetNumGamesBody{
		CommunityId: c.id,
		PlayerId:    playerId,
		NumGames:    numGames,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティのプレイヤーの試合数を変更する
//
// # 引数
//
//   - playerId: 変更するプレイヤーのId
//   - num: 変更後の試合数
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - 変更するプレイヤーが存在しない場合はエラー
func (c *Community) ChangePlayerNumGames(
	playerId player.PlayerId,
	num player.PlayerNumGames,
	executorId user.UserId,
) (CommunityEventChangeNumGamesBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventChangeNumGamesBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventChangeNumGamesBody{}, fmt.Errorf("executor is not staff")
	}
	if !c.players.IsPlayer(playerId) {
		return CommunityEventChangeNumGamesBody{}, fmt.Errorf("player is not player")
	}

	c.players.ChangePlayerNumGames(playerId, num)
	return CommunityEventChangeNumGamesBody{
		CommunityId: c.id,
		PlayerId:    playerId,
		OccurredAt:  time.Now(),
	}, nil
}

// コミュニティにメンバーを追加する
//
// # 引数
//
//   - memberId: 追加するメンバーのId
//   - memberRole: 追加するメンバーの役割
//   - userId: 追加するメンバーのユーザーId
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - 追加するメンバーが既にメンバーの場合はエラー
func (c *Community) AddMember(
	member member.Member,
	executorId user.UserId,
) (CommunityEventAddMemberBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventAddMemberBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventAddMemberBody{}, fmt.Errorf("executor is not staff")
	}
	if c.members.IsMember(member.BreachEncapsulationOfUserId()) {
		return CommunityEventAddMemberBody{}, fmt.Errorf("user is already member")
	}
	c.members.AddMember(member)
	return CommunityEventAddMemberBody{
		CommunityId: c.id,
		Member:      member,
		OccurredAt:  time.Now(),
	}, nil

}

// コミュニティからメンバーを削除する
//
// # 引数
//
//   - memberId: 削除するメンバーのId
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - 削除するメンバーが存在しない場合はエラー
func (c *Community) RemoveMember(
	userId user.UserId,
	executorId user.UserId,
) (CommunityEventRemoveMemberBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventRemoveMemberBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventRemoveMemberBody{}, fmt.Errorf("executor is not staff")
	}
	if !c.members.IsMember(userId) {
		return CommunityEventRemoveMemberBody{}, fmt.Errorf("member is not member")
	}

	c.members.RemoveMember(userId)
	return CommunityEventRemoveMemberBody{
		CommunityId: c.id,
		UserId:      userId,
		OccurredAt:  time.Now(),
	}, nil

}

// コミュニティのメンバーの役割を変更する
//
// # 引数
//
//   - userId: 変更するメンバーのユーザーId
//   - memberRole: 変更後のメンバーの役割
//   - executorId: 実行者のユーザーId
//
// # 戻り値
//
//   - 実行者がコミュニティのメンバーでない場合はエラー
//   - 実行者がコミュニティのスタッフでない場合はエラー
//   - 変更するメンバーが存在しない場合はエラー
func (c *Community) ChangeMemberRole(
	userId user.UserId,
	role member.MemberRole,
	executorId user.UserId,
) (CommunityEventChangeMemberRoleBody, error) {
	if !c.members.IsMember(executorId) {
		return CommunityEventChangeMemberRoleBody{}, fmt.Errorf("executor is not member")
	}
	if !c.members.IsStaff(executorId) {
		return CommunityEventChangeMemberRoleBody{}, fmt.Errorf("executor is not staff")
	}
	if !c.members.IsMember(userId) {
		return CommunityEventChangeMemberRoleBody{}, fmt.Errorf("member is not member")
	}
	c.members.ChangeRole(userId, role)
	return CommunityEventChangeMemberRoleBody{
		CommunityId: c.id,
		Member:      c.members.GetMemberByUserId(userId),
		OccurredAt:  time.Now(),
	}, nil

}
