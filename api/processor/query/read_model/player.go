package read_model

import "time"

type Player struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Age       int       `json:"age"`
	Level     string    `json:"level"`
	NumGames  int       `json:"num_games"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPlayer(
	id string,
	name string,
	gender string,
	age int,
	level string,
	numGames int,
	status string,
	createdAt time.Time,
	updatedAt time.Time,
) Player {
	return Player{
		ID:        id,
		Name:      name,
		Gender:    gender,
		Age:       age,
		Level:     level,
		NumGames:  numGames,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
