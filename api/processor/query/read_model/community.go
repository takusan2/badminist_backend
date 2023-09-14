package read_model

type Community struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
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
