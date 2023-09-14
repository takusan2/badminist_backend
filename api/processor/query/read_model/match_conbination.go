package read_model

type MatchCombination struct {
	Matches    []Match    `json:"matches"`
	RestPlayer PlayerList `json:"rest_player"`
}
