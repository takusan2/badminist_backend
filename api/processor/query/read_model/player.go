package read_model

type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
	Level    string `json:"level"`
	NumGames int    `json:"num_games"`
	Status   string `json:"status"`
}

func NewPlayer(
	id string,
	name string,
	gender string,
	age int,
	level string,
	numGames int,
	status string,
) Player {
	return Player{
		ID:       id,
		Name:     name,
		Gender:   gender,
		Age:      age,
		Level:    level,
		NumGames: numGames,
		Status:   status,
	}
}
