package read_model

type Member struct {
	ID   string `json:"id"`
	User User   `json:"user"`
	Role string `json:"role"`
}

func NewMember(
	id string,
	user User,
	role string,
) Member {
	return Member{
		ID:   id,
		User: user,
		Role: role,
	}
}
