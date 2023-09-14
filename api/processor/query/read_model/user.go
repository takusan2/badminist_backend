package read_model

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(
	id string,
	name string,
	email string,
) User {
	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
