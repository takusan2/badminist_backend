package read_model

import "time"

type Community struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCommunity(
	id string,
	name string,
	description string,
) Community {
	return Community{
		ID:          id,
		Name:        name,
		Description: description,
	}
}
