package read_model

type MatchCombination struct {
	Matches     []Match    `json:"matches"`
	RestPlayers PlayerList `json:"rest_players"`
}
