package processor

import (
	"errors"
	"math/rand"

	"github.com/takuya-okada-01/badminist/app/domain/community"
	"github.com/takuya-okada-01/badminist/app/domain/community/player"
	"github.com/takuya-okada-01/badminist/app/infrastructure/entity"
	"github.com/takuya-okada-01/badminist/app/interface_adaptor_if/dao_if"
	"github.com/takuya-okada-01/badminist/app/readmodel"

	"gorm.io/gorm"
)

type Rule int

const (
	Singles Rule = iota
	Doubles
)

func RuleFromStr(str string) Rule {
	switch str {
	case "singles":
		return Singles
	case "doubles":
		return Doubles
	default:
		return Singles
	}
}

type QueryProcessor interface {
	GenerateMatchCombination(
		communityId community.CommunityId,
		numCourt int,
		rule Rule,
	) (readmodel.MatchCombination, error)
}

type queryProcessor struct {
	db *gorm.DB
	dao_if.CommunityDao
}

func NewQueryProcessor(db *gorm.DB, communityDao dao_if.CommunityDao) QueryProcessor {
	return &queryProcessor{db, communityDao}
}

func (q *queryProcessor) GenerateMatchCombination(
	communityId community.CommunityId,
	numCourt int,
	rule Rule,
) (readmodel.MatchCombination, error) {
	status, _ := player.NewPlayerStatus(player.Attend)
	players, err := q.CommunityDao.FindPlayersWithStatusByCommunityId(q.db, communityId, status)
	if err != nil {
		return readmodel.MatchCombination{}, err
	}

	if len(players) < 2 {
		return readmodel.MatchCombination{}, errors.New("not enough attend players")
	}

	var matches []readmodel.Match
	var restPlayer []entity.Player
	if rule == Singles {
		matches, restPlayer = q.generateSinglesMatchCombination(players, numCourt)
	} else {
		matches, restPlayer = q.generateDoublesMatchCombination(players, numCourt)
	}
	return readmodel.MatchCombination{Matches: matches, RestPlayer: restPlayer}, nil
}

func (q *queryProcessor) generateSinglesMatchCombination(
	players []entity.Player,
	numCourt int,
) ([]readmodel.Match, []entity.Player) {
	var matches []readmodel.Match
	var restPlayer []entity.Player

	// シャッフル
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	if len(players) < numCourt*2 {
		if len(players)%2 == 0 {
			restPlayer = []entity.Player{}
		} else {
			restPlayer = players[len(players)-1:]
			players = players[:len(players)-1]
		}
	} else {
		restPlayer = players[numCourt*2:]
		players = players[:numCourt*2]
	}

	for i := 0; i < len(players); i += 2 {
		left := readmodel.Team{Players: []entity.Player{players[i]}}
		right := readmodel.Team{Players: []entity.Player{players[i+1]}}
		matches = append(matches, readmodel.Match{Left: left, Right: right})
	}
	return matches, restPlayer
}

func (q *queryProcessor) generateDoublesMatchCombination(
	players []entity.Player,
	numCourt int,
) ([]readmodel.Match, []entity.Player) {
	var matches []readmodel.Match
	var restPlayer []entity.Player

	// シャッフル
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	if len(players) < numCourt*4 {
		if len(players)%4 == 0 {
			restPlayer = []entity.Player{}
		} else {
			restPlayer = players[len(players)-len(players)%4:]
			players = players[:len(players)-len(players)%4]
		}
	} else {
		restPlayer = players[numCourt*4:]
		players = players[:numCourt*4]
	}

	for i := 0; i < len(players); i += 4 {
		left := readmodel.Team{Players: []entity.Player{players[i], players[i+1]}}
		right := readmodel.Team{Players: []entity.Player{players[i+2], players[i+3]}}
		matches = append(matches, readmodel.Match{Left: left, Right: right})
	}
	return matches, restPlayer
}
