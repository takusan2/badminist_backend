package query_processor

import (
	"errors"
	"math/rand"

	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	query_dao_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/dao_if/query"

	"github.com/takuya-okada-01/badminist/api/processor/query/read_model"

	"gorm.io/gorm"
)

type Rule int

const (
	Singles Rule = iota
	Doubles
)

func RuleFromStr(str string) (Rule, error) {
	switch str {
	case "singles":
		return Singles, nil
	case "doubles":
		return Doubles, nil
	default:
		return Singles, errors.New("invalid rule")
	}
}

type QueryProcessor interface {
	GenerateMatchCombination(
		communityId community.CommunityId,
		numCourt int,
		rule Rule,
	) (read_model.MatchCombination, error)
	GetCommunityList(
		userId user.UserId,
	) (read_model.CommunityList, error)
	GetPlayerList(
		communityId community.CommunityId,
	) (read_model.PlayerList, error)
	GetMemberList(
		communityId community.CommunityId,
	) (read_model.MemberList, error)
	GetUserById(
		id user.UserId,
	) (read_model.User, error)
}

type queryProcessor struct {
	db           *gorm.DB
	communityDao query_dao_if.CommunityDao
	userDao      query_dao_if.UserDao
}

func NewQueryProcessor(
	db *gorm.DB,
	communityDao query_dao_if.CommunityDao,
	userDao query_dao_if.UserDao,
) QueryProcessor {
	return &queryProcessor{db, communityDao, userDao}
}

func (q *queryProcessor) GenerateMatchCombination(
	communityId community.CommunityId,
	numCourt int,
	rule Rule,
) (read_model.MatchCombination, error) {
	status, _ := player.NewPlayerStatus(player.Attend)
	players, err := q.communityDao.FindPlayersWithStatusByCommunityId(q.db, communityId, status)
	if err != nil {
		return read_model.MatchCombination{}, err
	}
	if len(players) < 2 {
		return read_model.MatchCombination{}, errors.New("not enough attend players")
	}

	var playerList []read_model.Player
	for _, p := range players {
		playerList = append(playerList, p.ToReadModel())
	}

	var matches []read_model.Match
	var restPlayer read_model.PlayerList
	if rule == Singles {
		matches, restPlayer = q.generateSinglesMatchCombination(playerList, numCourt)
	} else {
		matches, restPlayer = q.generateDoublesMatchCombination(playerList, numCourt)
	}
	return read_model.MatchCombination{Matches: matches, RestPlayer: restPlayer}, nil
}

func (q *queryProcessor) generateSinglesMatchCombination(
	players read_model.PlayerList,
	numCourt int,
) ([]read_model.Match, read_model.PlayerList) {
	var matches []read_model.Match
	var restPlayer read_model.PlayerList

	// シャッフル
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	if len(players) < numCourt*2 {
		if len(players)%2 == 0 {
			restPlayer = read_model.PlayerList{}
		} else {
			restPlayer = players[len(players)-1:]
			players = players[:len(players)-1]
		}
	} else {
		restPlayer = players[numCourt*2:]
		players = players[:numCourt*2]
	}

	for i := 0; i < len(players); i += 2 {
		left := read_model.Team{Players: read_model.PlayerList{players[i]}}
		right := read_model.Team{Players: read_model.PlayerList{players[i+1]}}
		matches = append(matches, read_model.Match{Left: left, Right: right})
	}
	return matches, restPlayer
}

func (q *queryProcessor) generateDoublesMatchCombination(
	players read_model.PlayerList,
	numCourt int,
) ([]read_model.Match, read_model.PlayerList) {
	var matches []read_model.Match
	var restPlayer read_model.PlayerList

	// シャッフル
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	if len(players) < numCourt*4 {
		if len(players)%4 == 0 {
			restPlayer = read_model.PlayerList{}
		} else {
			restPlayer = players[len(players)-len(players)%4:]
			players = players[:len(players)-len(players)%4]
		}
	} else {
		restPlayer = players[numCourt*4:]
		players = players[:numCourt*4]
	}

	for i := 0; i < len(players); i += 4 {
		left := read_model.Team{Players: read_model.PlayerList{players[i], players[i+1]}}
		right := read_model.Team{Players: read_model.PlayerList{players[i+2], players[i+3]}}
		matches = append(matches, read_model.Match{Left: left, Right: right})
	}
	return matches, restPlayer
}

func (q *queryProcessor) GetCommunityList(
	userId user.UserId,
) (read_model.CommunityList, error) {
	communityEntity, err := q.communityDao.FindCommunitiesByUserId(q.db, userId)
	if err != nil {
		return read_model.CommunityList{}, err
	}

	var communityList read_model.CommunityList
	for _, c := range communityEntity {
		communityList = append(communityList, c.ToReadModel())
	}
	return communityList, nil
}

func (q *queryProcessor) GetPlayerList(
	communityId community.CommunityId,
) (read_model.PlayerList, error) {
	players, err := q.communityDao.FindPlayersByCommunityId(q.db, communityId)
	if err != nil {
		return read_model.PlayerList{}, err
	}

	var playerList read_model.PlayerList
	for _, p := range players {
		playerList = append(playerList, p.ToReadModel())
	}
	return playerList, nil
}

func (q *queryProcessor) GetMemberList(
	communityId community.CommunityId,
) (read_model.MemberList, error) {
	members, err := q.communityDao.FindMembersByCommunityId(q.db, communityId)
	if err != nil {
		return read_model.MemberList{}, err
	}

	var memberList read_model.MemberList
	for _, m := range members {
		memberList = append(memberList, m.ToReadModel())
	}
	return memberList, nil
}

func (q *queryProcessor) GetUserById(
	id user.UserId,
) (read_model.User, error) {
	user, err := q.userDao.FindUserById(q.db, id)
	if err != nil {
		return read_model.User{}, err
	}
	return user.ToReadModel(), nil
}
